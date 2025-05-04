-- name: CreateProduct :exec
INSERT INTO products (
  id, 
  name,
  description,
  brand,
  origin,
  user_guide,
  category_id,
  status,
  created_at,
  created_by,
  updated_at,
  updated_by
  )
VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
);

-- name: GetProductById :one
SELECT *
FROM products
WHERE id = $1 AND status != 'deleted';

-- name: GetProducts :many
SELECT l.*
FROM products l
WHERE   
    (
        (@status::text = '' AND l.status != 'deleted') OR 
        (l.status = @status::text)
    )
    AND 
    (
        (@category_id::text = '' OR l.category_id = @category_id::uuid)
    )
    AND (@keyword::text = '' OR l.name ilike concat('%',@keyword::text,'%'))
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

-- name: GetMyProducts :many
SELECT l.*
FROM products l
WHERE created_by = $3 AND
    (
        (@status::text = '' AND l.status != 'deleted') OR 
        (l.status = @status::text)
    )
    AND 
    (
        (@category_id::text = '' OR l.category_id = @category_id::uuid)
    )
    AND (@keyword::text = '' OR l.name ilike concat('%',@keyword::text,'%'))
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

-- name: CountMyProducts :one
SELECT COUNT(*)::int
FROM products l
WHERE created_by = $1 AND
    (
        (@status::text = '' AND l.status != 'deleted') OR 
        (l.status = @status::text)
    )
    AND 
    (
        (@category_id::text = '' OR l.category_id = @category_id::uuid)
    )
    AND (@keyword::text = '' OR l.name ilike concat('%',@keyword::text,'%'));

-- name: CountProducts :one
SELECT COUNT(*)::int
FROM products l
WHERE   
    (
        (@status::text = '' AND l.status != 'deleted') OR 
        (l.status = @status::text)
    )
    AND 
    (
        (@category_id::text = '' OR l.category_id = @category_id::uuid)
    )
    AND (@keyword::text = '' OR l.name ilike concat('%',@keyword::text,'%'));

-- name: GetProductsByCategoryId :many
SELECT l.*
FROM products l
WHERE l.category_id = $3

 AND l.status != 'deleted' AND
    (@keyword::text = '' or l.name ilike concat('%',@keyword::text,'%'))
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

-- name: CountProductsByCategoryId :one
SELECT COUNT(*)::int
FROM products l
WHERE l.category_id = $1 AND l.status != 'deleted' AND 
    (@keyword::text = '' or l.name ilike concat('%',@keyword::text,'%'));

-- name: UpdateProduct :exec
UPDATE products
SET
  name = $2,
  description = $3,
  brand = $4,
  origin = $5,
  user_guide = $6,
  category_id = $7,
  status = $8,
  updated_at = $9,
  updated_by = $10
WHERE id = $1 AND status != 'deleted';

-- name: DeleteProduct :exec
UPDATE products
SET
  status = 'deleted',
  deleted_at = $2,
  deleted_by = $3
WHERE id = $1 AND status != 'deleted';
