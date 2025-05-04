-- name: CreateProductVariant :exec
INSERT INTO product_variants (
  id, 
  product_id,
  description,
  color,
  retail_price,
  stock,
  status,
  created_at,
  created_by,
  updated_at,
  updated_by
  )
VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
);

-- name: GetProductVariantById :one
SELECT *
FROM product_variants
WHERE id = $1 AND status != 'deleted';

-- name: GetProductVariantsByProductId :many
SELECT l.*
FROM product_variants l
WHERE l.product_id = $1 AND l.status != 'deleted';

-- name: UpdateProductVariant :exec
UPDATE product_variants
SET
  product_id = $2,
  description = $3,
  color = $4,
  retail_price = $5,
  stock = $6,
  status = $7,
  updated_at = $8,
  updated_by = $9
WHERE id = $1 AND status != 'deleted';

-- name: DeleteProductVariant :exec
UPDATE product_variants
SET
  status = 'deleted',
  deleted_at = $2,
  deleted_by = $3
WHERE id = $1 AND status != 'deleted';

-- name: IsColorExist :one
SELECT EXISTS (
  SELECT 1
  FROM product_variants
  WHERE color = $1
    AND product_id = $2
    AND status != 'deleted'
) AS exists;


