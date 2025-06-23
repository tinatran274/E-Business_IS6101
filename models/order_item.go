package models

import (
	"context"

	db "10.0.0.50/tuan.quang.tran/aioz-ads/db/generated"
	"github.com/google/uuid"
)

type OrderItemRepository interface {
	CreateOrderItem(
		ctx context.Context,
		orderItem *OrderItem,
	) (*OrderItem, error)
	GetOrderItemsByOrderID(
		ctx context.Context,
		orderID uuid.UUID,
	) ([]*OrderItem, error)
}

type OrderItem struct {
	OrderID          uuid.UUID `json:"order_id"`
	ProductVariantID uuid.UUID `json:"product_variant_id"`
	Quantity         int32     `json:"quantity"`
	RetailPrice      float64   `json:"retail_price"`
}

type OrderItemRequest struct {
	ProductVariantID uuid.UUID `json:"product_variant_id"`
	Quantity         int32     `json:"quantity"`
}

func NewOrderItem(
	orderID, productVariantID uuid.UUID,
	quantity int32,
	retailPrice float64,
) *OrderItem {
	return &OrderItem{
		OrderID:          orderID,
		ProductVariantID: productVariantID,
		Quantity:         quantity,
		RetailPrice:      retailPrice,
	}
}

func ToOrderItem(db db.OrderItem) *OrderItem {
	return &OrderItem{
		OrderID:          db.OrderID,
		ProductVariantID: db.ProductVariantID,
		Quantity:         db.Quantity,
		RetailPrice:      db.RetailPrice,
	}
}
