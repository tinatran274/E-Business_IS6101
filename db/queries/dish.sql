-- name: GetDishById :one
SELECT *
FROM dishes
WHERE id = $1 AND status != 'deleted';

-- name: GetDishes :many
SELECT l.*
FROM dishes l
WHERE (@keyword::text = '' or l.name ilike concat('%',@keyword::text,'%'))
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

-- name: CreateDish :exec
INSERT INTO dishes (
  id, 
  name,
  description,
  category_id,
  status,
  created_at,
  created_by,
  updated_at,
  updated_by
  )
VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9
);

-- name: UpdateDish :exec
UPDATE dishes
SET
  name = $2,
  description = $3,
  category_id = $4,
  status = $5,
  updated_at = $6,
  updated_by = $7
WHERE id = $1 AND status != 'deleted';

-- name: DeleteDish :exec
UPDATE dishes
SET
  status = 'deleted',
  deleted_at = $2,
  deleted_by = $3
WHERE id = $1 AND status != 'deleted';


-- name: GetDishByIngredientId :many
SELECT d.*
FROM dishes d
JOIN recipes r ON d.id = r.dish_id
WHERE r.ingredient_id = $1 AND d.status != 'deleted';

-- name: CountDishes :one
SELECT COUNT(*) FROM dishes
WHERE status != 'deleted';