-- name: AddCartItem :exec
INSERT INTO cart (
    user_id, 
    product_variant_id,
    quantity
) VALUES ($1, $2, $3);

-- name: GetCartItemByUserIdAndProductVariantId :one
SELECT *
FROM cart
WHERE user_id = $1 AND product_variant_id = $2;

-- name: GetCartItemsByUserId :many
SELECT *
FROM cart
WHERE user_id = $3
LIMIT $1
OFFSET $2;

-- name: CountCartItemsByUserId :one
SELECT COUNT(*) AS count
FROM cart
WHERE user_id = $1;

-- name: UpdateCartItem :exec
UPDATE cart
SET 
    quantity = $3
WHERE user_id = $1 AND product_variant_id = $2;

-- name: DeleteCartItem :exec
DELETE FROM cart
WHERE user_id = $1 AND product_variant_id = $2;