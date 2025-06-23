-- name: GetPaymentMethodById :one
SELECT *
FROM payment_methods
WHERE id = $1 AND status != 'deleted';

-- name: GetPaymentMethods :many
SELECT *
FROM payment_methods
WHERE status != 'deleted'
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

-- name: CountPaymentMethods :one
SELECT COUNT(*) AS count
FROM payment_methods
WHERE status != 'deleted';

