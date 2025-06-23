package repositories

import (
	"context"
	"time"

	db "10.0.0.50/tuan.quang.tran/aioz-ads/db/generated"
	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/utils/metrics"
	"10.0.0.50/tuan.quang.tran/aioz-ads/models"
	"github.com/google/uuid"
)

type OrderItemRepository struct {
	db *db.Queries
}

func NewOrderItemRepository(db *db.Queries) *OrderItemRepository {
	return &OrderItemRepository{
		db: db,
	}
}

func (r *OrderItemRepository) CreateOrderItem(
	ctx context.Context,
	orderItem *models.OrderItem,
) (*models.OrderItem, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("CreateOrderItem").
			Observe(time.Since(t).Seconds())
	}()

	params := db.CreateOrderItemParams{
		OrderID:          orderItem.OrderID,
		ProductVariantID: orderItem.ProductVariantID,
		Quantity:         orderItem.Quantity,
		RetailPrice:      orderItem.RetailPrice,
	}
	err := r.db.CreateOrderItem(ctx, params)
	if err != nil {
		return nil, err
	}

	return orderItem, nil
}

func (r *OrderItemRepository) GetOrderItemsByOrderID(
	ctx context.Context,
	orderID uuid.UUID,
) ([]*models.OrderItem, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("GetOrderItemsByOrderID").
			Observe(time.Since(t).Seconds())
	}()

	orderItemsDB, err := r.db.GetOrderItemsByOrderId(ctx, orderID)
	if err != nil {
		return nil, err
	}

	orderItems := make([]*models.OrderItem, len(orderItemsDB))
	for i, item := range orderItemsDB {
		orderItems[i] = models.ToOrderItem(item)
	}

	return orderItems, nil
}
