-- name: GetIngredientById :one
SELECT *
FROM ingredients
WHERE id = $1 AND status != 'deleted';

-- name: GetAllIngredient :one
SELECT *
FROM ingredients 
WHERE status != 'deleted';

-- name: CreateIngredient :exec
INSERT INTO ingredients (
  id, 
  name,
  description,
  removal,
  kcal,
  protein,
  lipits,
  glucids,
  canxi,
  phosphor,
  fe,
  vitamin_a,
  vitamin_b1,
  vitamin_b2,
  vitamin_c,
  vitamin_pp,
  beta_caroten,
  category_id,
  status,
  created_at,
  created_by,
  updated_at,
  updated_by
  )
VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23
);

-- name: UpdateIngredient :exec
UPDATE ingredients
SET
  name = $2,
  description = $3,
  removal = $4,
  kcal = $5,
  protein = $6,
  lipits = $7,
  glucids = $8,
  canxi = $9, 
  phosphor = $10,
  fe = $11,
  vitamin_a = $12,
  vitamin_b1 = $13,
  vitamin_b2 = $14,
  vitamin_c = $15,
  vitamin_pp = $16,
  beta_caroten = $17,
  category_id = $18,
  status = $19,
  updated_at = $20,
  updated_by = $21
WHERE id = $1 AND status != 'deleted';

-- name: DeleteIngredient :exec
UPDATE ingredients
SET
  status = 'deleted',
  deleted_at = $2,
  deleted_by = $3
WHERE id = $1 AND status != 'deleted';

-- name: GetIngredientByDishId :many
SELECT i.* 
FROM ingredients i
JOIN recipes r ON i.id = r.ingredient_id
WHERE r.dish_id = $1 AND i.status != 'deleted';


