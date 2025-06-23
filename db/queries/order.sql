-- name: CreateOrder :exec
INSERT INTO orders (
    id,
    user_id,
    order_date,
    receiver_name,
    receiver_phone,
    receiver_address,
    shipping_cost,
    payment_method_id,
    payment_status,
    shipping_status,
    order_status,
    created_at,
    created_by,
    updated_at,
    updated_by
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15);

-- name: GetOrderById :one
SELECT *
FROM orders
WHERE id = $1 AND order_status != 'deleted';

-- name: GetOrdersByUserId :many
SELECT l.*
FROM orders l
WHERE created_by = $3 AND
    (
        (@status::text = '' AND l.status != 'deleted') OR 
        (l.status = @status::text)
    )
ORDER BY
    CASE 
        WHEN @sort_by::text = 'created_at' THEN 
            CASE 
                WHEN @order_by::text = 'asc' THEN l.created_at 
            END 
    END ASC,
    CASE 
        WHEN @sort_by::text = 'created_at' THEN 
            CASE 
                WHEN @order_by::text = 'desc' THEN l.created_at 
            END 
    END DESC
LIMIT $1
OFFSET $2;

-- name: CountOrdersByUserId :one
SELECT COUNT(*) AS count
FROM orders l
WHERE created_by = $1 AND
    (
        (@status::text = '' AND l.status != 'deleted') OR 
        (l.status = @status::text)
    );

-- name: UpdateOrder :exec
UPDATE orders
SET
    user_id = $2,
    order_date = $3,
    receiver_name = $4,
    receiver_phone = $5,
    receiver_address = $6,
    shipping_cost = $7,
    payment_method_id = $8,
    payment_status = $9,
    shipping_status = $10,
    order_status = $11,
    updated_at = $12,
    updated_by = $13
WHERE id = $1;

-- name: DeleteOrder :exec
UPDATE orders
SET
    order_status = 'deleted',
    deleted_at = $2,
    deleted_by = $3
WHERE id = $1;