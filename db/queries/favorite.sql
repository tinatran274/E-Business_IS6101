-- name: CreateFavorite :exec
INSERT INTO favorites (
  user_id,
  dish_id,
  value
) VALUES (
  $1, $2, $3
);

-- name: DeleteFavorite :exec
DELETE FROM favorites
WHERE user_id = $1 AND dish_id = $2;

-- name: UpdateFavorite :exec
UPDATE favorites
SET
  value = $3
WHERE user_id = $1 AND dish_id = $2;

-- name: GetFavoriteByUserIdAndDishId :one
SELECT *
FROM favorites
WHERE user_id = $1 AND dish_id = $2;