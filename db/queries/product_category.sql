-- name: CreateProductCategory :exec
INSERT INTO product_categories (
  id, 
  name,
  description,
  status,
  created_at,
  created_by,
  updated_at,
  updated_by
  )
VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
);

-- name: GetProductCategoryById :one
SELECT *   
FROM product_categories
WHERE id = $1 AND status != 'deleted';

-- name: GetProductCategories :many
SELECT *
FROM product_categories 
WHERE status != 'deleted';

-- name: UpdateProductCategory :exec
UPDATE product_categories
SET
  name = $2,
  description = $3,
  status = $4,
  updated_at = $5,
  updated_by = $6
WHERE id = $1 AND status != 'deleted';

-- name: DeleteProductCategory :exec
UPDATE product_categories
SET
  status = 'deleted',
  deleted_at = $2,
  deleted_by = $3
WHERE id = $1 AND status != 'deleted';