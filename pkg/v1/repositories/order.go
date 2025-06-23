package repositories

import (
	"context"
	"time"

	db "10.0.0.50/tuan.quang.tran/aioz-ads/db/generated"
	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/utils/metrics"
	"10.0.0.50/tuan.quang.tran/aioz-ads/models"
	"github.com/google/uuid"
)

type OrderRepository struct {
	db *db.Queries
}

func NewOrderRepository(
	db *db.Queries,
) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (r *OrderRepository) CreateOrder(
	ctx context.Context,
	order *models.Order,
) (*models.Order, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("CreateOrder").
			Observe(time.Since(t).Seconds())
	}()

	orderDB := db.CreateOrderParams{
		ID:              order.ID,
		UserID:          order.UserID,
		OrderDate:       order.OrderDate,
		ReceiverName:    order.ReceiverName,
		ReceiverPhone:   order.ReceiverPhone,
		ReceiverAddress: order.ReceiverAddress,
		ShippingCost:    order.ShippingCost,
		PaymentMethodID: order.PaymentMethodID,
		PaymentStatus:   order.PaymentStatus,
		ShippingStatus:  order.ShippingStatus,
		OrderStatus:     order.OrderStatus,
		CreatedAt:       order.CreatedAt,
		CreatedBy:       order.CreatedBy,
		UpdatedAt:       order.UpdatedAt,
		UpdatedBy:       order.UpdatedBy,
	}

	err := r.db.CreateOrder(ctx, orderDB)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (r *OrderRepository) GetOrderByID(
	ctx context.Context,
	id uuid.UUID,
) (*models.Order, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("GetOrderByID").
			Observe(time.Since(t).Seconds())
	}()

	orderDB, err := r.db.GetOrderById(ctx, id)
	if err != nil {
		return nil, err
	}

	return models.ToOrder(orderDB), nil
}

func (r *OrderRepository) GetOrdersByUserID(
	ctx context.Context,
	userID uuid.UUID,
	filter *models.FilterParams,
) ([]*models.Order, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("GetOrdersByUserID").
			Observe(time.Since(t).Seconds())
	}()

	ordersDB, err := r.db.GetOrdersByUserId(
		ctx,
		db.GetOrdersByUserIdParams{
			CreatedBy: &userID,
			Limit:     filter.Limit,
			Offset:    filter.Offset,
			SortBy:    filter.SortBy,
			OrderBy:   filter.OrderBy,
			Status:    filter.Status,
		},
	)
	if err != nil {
		return nil, err
	}

	orders := make([]*models.Order, 0, len(ordersDB))
	for _, orderDB := range ordersDB {
		orders = append(orders, models.ToOrder(orderDB))
	}

	return orders, nil
}

func (r *OrderRepository) CountOrdersByUserID(
	ctx context.Context,
	userID uuid.UUID,
	filter *models.FilterParams,
) (int, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("CountOrdersByUserID").
			Observe(time.Since(t).Seconds())
	}()

	count, err := r.db.CountOrdersByUserId(
		ctx,
		db.CountOrdersByUserIdParams{
			CreatedBy: &userID,
			Status:    filter.Status,
		},
	)
	if err != nil {
		return 0, err
	}

	return int(count), nil
}

func (r *OrderRepository) UpdateOrder(
	ctx context.Context,
	order *models.Order,
) (*models.Order, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("UpdateOrder").
			Observe(time.Since(t).Seconds())
	}()

	orderDB := db.UpdateOrderParams{
		ID:              order.ID,
		UserID:          order.UserID,
		OrderDate:       order.OrderDate,
		ReceiverName:    order.ReceiverName,
		ReceiverPhone:   order.ReceiverPhone,
		ReceiverAddress: order.ReceiverAddress,
		ShippingCost:    order.ShippingCost,
		PaymentMethodID: order.PaymentMethodID,
		PaymentStatus:   order.PaymentStatus,
		ShippingStatus:  order.ShippingStatus,
		OrderStatus:     order.OrderStatus,
		UpdatedAt:       order.UpdatedAt,
		UpdatedBy:       order.UpdatedBy,
	}

	err := r.db.UpdateOrder(ctx, orderDB)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (r *OrderRepository) DeleteOrder(
	ctx context.Context,
	order *models.Order,
) error {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("DeleteOrder").
			Observe(time.Since(t).Seconds())
	}()

	err := r.db.DeleteOrder(ctx, db.DeleteOrderParams{
		ID:        order.ID,
		DeletedAt: order.DeletedAt,
		DeletedBy: order.DeletedBy,
	})
	if err != nil {
		return err
	}

	return nil
}
