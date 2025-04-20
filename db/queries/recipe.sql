-- name: CreateRecipe :exec
INSERT INTO recipes (
  dish_id,
  ingredient_id,
  unit
  )
VALUES (
  $1, $2, $3
);

-- name: DeleteRecipe :exec
DELETE FROM recipes
WHERE dish_id = $1 AND ingredient_id = $2;