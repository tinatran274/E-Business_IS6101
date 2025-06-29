// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: order_items.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createOrderItem = `-- name: CreateOrderItem :exec
INSERT INTO order_items (
    order_id,
    product_variant_id,
    quantity,
    retail_price
) VALUES ($1, $2, $3, $4)
`

type CreateOrderItemParams struct {
	OrderID          uuid.UUID `json:"order_id"`
	ProductVariantID uuid.UUID `json:"product_variant_id"`
	Quantity         int32     `json:"quantity"`
	RetailPrice      float64   `json:"retail_price"`
}

func (q *Queries) CreateOrderItem(ctx context.Context, arg CreateOrderItemParams) error {
	_, err := q.db.Exec(ctx, createOrderItem,
		arg.OrderID,
		arg.ProductVariantID,
		arg.Quantity,
		arg.RetailPrice,
	)
	return err
}

const getOrderItemsByOrderId = `-- name: GetOrderItemsByOrderId :many
SELECT oi.order_id, oi.product_variant_id, oi.quantity, oi.retail_price
FROM order_items oi
JOIN orders o ON oi.order_id = o.id
WHERE oi.order_id = $1 AND o.order_status != 'deleted'
`

func (q *Queries) GetOrderItemsByOrderId(ctx context.Context, orderID uuid.UUID) ([]OrderItem, error) {
	rows, err := q.db.Query(ctx, getOrderItemsByOrderId, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []OrderItem
	for rows.Next() {
		var i OrderItem
		if err := rows.Scan(
			&i.OrderID,
			&i.ProductVariantID,
			&i.Quantity,
			&i.RetailPrice,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
