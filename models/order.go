package models

import (
	"context"
	"time"

	db "10.0.0.50/tuan.quang.tran/aioz-ads/db/generated"
	"github.com/google/uuid"
)

type OrderRepository interface {
	CreateOrder(
		ctx context.Context,
		order *Order,
	) (*Order, error)
	GetOrderByID(
		ctx context.Context,
		id uuid.UUID,
	) (*Order, error)
	GetOrdersByUserID(
		ctx context.Context,
		userID uuid.UUID,
		filter *FilterParams,
	) ([]*Order, error)
	CountOrdersByUserID(
		ctx context.Context,
		userID uuid.UUID,
		filter *FilterParams,
	) (int, error)
	UpdateOrder(
		ctx context.Context,
		order *Order,
	) (*Order, error)
	DeleteOrder(
		ctx context.Context,
		order *Order,
	) error
}

type Order struct {
	ID              uuid.UUID  `json:"id"`
	UserID          uuid.UUID  `json:"user_id"`
	OrderDate       time.Time  `json:"order_date"`
	ReceiverName    string     `json:"receiver_name"`
	ReceiverPhone   string     `json:"receiver_phone"`
	ReceiverAddress string     `json:"receiver_address"`
	ShippingCost    float64    `json:"shipping_cost"`
	PaymentMethodID uuid.UUID  `json:"payment_method_id"`
	PaymentStatus   string     `json:"payment_status"`
	ShippingStatus  string     `json:"shipping_status"`
	OrderStatus     string     `json:"order_status"`
	CreatedAt       time.Time  `json:"created_at"`
	CreatedBy       *uuid.UUID `json:"created_by"`
	UpdatedAt       time.Time  `json:"updated_at"`
	UpdatedBy       *uuid.UUID `json:"updated_by"`
	DeletedAt       *time.Time `json:"deleted_at"`
	DeletedBy       *uuid.UUID `json:"deleted_by"`
}

type OrderRequest struct {
	ReceiverName    string              `json:"receiver_name"`
	ReceiverPhone   string              `json:"receiver_phone"`
	ReceiverAddress string              `json:"receiver_address"`
	PaymentMethodID uuid.UUID           `json:"payment_method_id"`
	OrderItems      []*OrderItemRequest `json:"order_items"`
}

func NewOrder(
	userID uuid.UUID,
	orderDate time.Time,
	receiverName, receiverPhone, receiverAddress string,
	shippingCost float64,
	paymentMethodID uuid.UUID,
	createdBy *uuid.UUID,
) *Order {
	return &Order{
		ID:              uuid.New(),
		UserID:          userID,
		OrderDate:       orderDate,
		ReceiverName:    receiverName,
		ReceiverPhone:   receiverPhone,
		ReceiverAddress: receiverAddress,
		ShippingCost:    shippingCost,
		PaymentMethodID: paymentMethodID,
		PaymentStatus:   PendingStatus,
		ShippingStatus:  PendingStatus,
		OrderStatus:     ActiveStatus,
		CreatedAt:       time.Now().UTC(),
		CreatedBy:       createdBy,
		UpdatedAt:       time.Now().UTC(),
		UpdatedBy:       createdBy,
	}
}

func ToOrder(d db.Order) *Order {
	return &Order{
		ID:              d.ID,
		UserID:          d.UserID,
		OrderDate:       d.OrderDate,
		ReceiverName:    d.ReceiverName,
		ReceiverPhone:   d.ReceiverPhone,
		ReceiverAddress: d.ReceiverAddress,
		ShippingCost:    d.ShippingCost,
		PaymentMethodID: d.PaymentMethodID,
		PaymentStatus:   d.PaymentStatus,
		ShippingStatus:  d.ShippingStatus,
		OrderStatus:     d.OrderStatus,
		CreatedAt:       d.CreatedAt,
		CreatedBy:       d.CreatedBy,
		UpdatedAt:       d.UpdatedAt,
		UpdatedBy:       d.UpdatedBy,
	}
}
