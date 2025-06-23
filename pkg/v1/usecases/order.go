package usecases

import (
	"context"
	"regexp"
	"time"

	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/utils/response"
	"10.0.0.50/tuan.quang.tran/aioz-ads/models"
	"github.com/jackc/pgx/v5"
)

type OrderUseCase struct {
	OrderRepo          models.OrderRepository
	OrderItemRepo      models.OrderItemRepository
	CartRepo           models.CartRepository
	PaymentMethodRepo  models.PaymentMethodRepository
	ProductVariantRepo models.ProductVariantRepository
}

func NewOrderUseCase(
	orderRepo models.OrderRepository,
	orderItemRepo models.OrderItemRepository,
	cartRepo models.CartRepository,
	paymentMethodRepo models.PaymentMethodRepository,
	productVariantRepo models.ProductVariantRepository,
) *OrderUseCase {
	return &OrderUseCase{
		OrderRepo:          orderRepo,
		OrderItemRepo:      orderItemRepo,
		CartRepo:           cartRepo,
		PaymentMethodRepo:  paymentMethodRepo,
		ProductVariantRepo: productVariantRepo,
	}
}

func (o *OrderUseCase) CreateOrder(
	ctx context.Context,
	authInfo models.AuthenticationInfo,
	orderRequest *models.OrderRequest,
) (*models.Order, error) {
	paymentMethod, err := o.PaymentMethodRepo.GetPaymentMethodByID(
		ctx,
		orderRequest.PaymentMethodID,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, response.NewNotFoundError("Payment method not found.")
		}

		return nil, response.NewInternalServerError(err)
	}

	isValidPhoneNum := isValidPhoneNumber(orderRequest.ReceiverPhone)
	if !isValidPhoneNum {
		return nil, response.NewBadRequestError("Invalid phone number.")
	}

	//Validate receiver address then calc distane from seller and buyer

	order := models.NewOrder(
		authInfo.User.ID,
		time.Now().UTC(),
		orderRequest.ReceiverName,
		orderRequest.ReceiverPhone,
		orderRequest.ReceiverAddress,
		models.DefaultShippingCost,
		paymentMethod.ID,
		&authInfo.User.ID,
	)
	orderItems := make([]*models.OrderItem, 0, len(orderRequest.OrderItems))
	for _, item := range orderRequest.OrderItems {
		productVariant, err := o.ProductVariantRepo.GetProductVariantByID(
			ctx,
			item.ProductVariantID,
		)
		if err != nil {
			if err == pgx.ErrNoRows {
				return nil, response.NewNotFoundError("Product variant not found.")
			}

			return nil, response.NewInternalServerError(err)
		}

		if item.Quantity > productVariant.Stock {
			return nil, response.NewBadRequestError("Not enough product in stock.")
		}

		orderItem := models.NewOrderItem(
			order.ID,
			productVariant.ID,
			item.Quantity,
			productVariant.RetailPrice,
		)
		orderItems = append(orderItems, orderItem)
	}

	order, err = o.OrderRepo.CreateOrder(ctx, order)
	if err != nil {
		return nil, response.NewInternalServerError(err)
	}

	for _, item := range orderItems {
		_, err := o.OrderItemRepo.CreateOrderItem(ctx, item)
		if err != nil {
			return nil, response.NewInternalServerError(err)
		}
	}

	return order, nil
}

func isValidPhoneNumber(phone string) bool {
	re := regexp.MustCompile(`^\+?[0-9]{8,15}$`)
	return re.MatchString(phone)
}
