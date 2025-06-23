-- name: CreateOrderItem :exec
INSERT INTO order_items (
    order_id,
    product_variant_id,
    quantity,
    retail_price
) VALUES ($1, $2, $3, $4);


-- name: GetOrderItemsByOrderId :many
SELECT oi.*
FROM order_items oi
JOIN orders o ON oi.order_id = o.id
WHERE oi.order_id = $1 AND o.order_status != 'deleted';

