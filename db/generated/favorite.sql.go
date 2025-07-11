// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: favorite.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createFavorite = `-- name: CreateFavorite :exec
INSERT INTO favorites (
  user_id,
  dish_id,
  value
) VALUES (
  $1, $2, $3
)
`

type CreateFavoriteParams struct {
	UserID uuid.UUID `json:"user_id"`
	DishID uuid.UUID `json:"dish_id"`
	Value  int32     `json:"value"`
}

func (q *Queries) CreateFavorite(ctx context.Context, arg CreateFavoriteParams) error {
	_, err := q.db.Exec(ctx, createFavorite, arg.UserID, arg.DishID, arg.Value)
	return err
}

const deleteFavorite = `-- name: DeleteFavorite :exec
DELETE FROM favorites
WHERE user_id = $1 AND dish_id = $2
`

type DeleteFavoriteParams struct {
	UserID uuid.UUID `json:"user_id"`
	DishID uuid.UUID `json:"dish_id"`
}

func (q *Queries) DeleteFavorite(ctx context.Context, arg DeleteFavoriteParams) error {
	_, err := q.db.Exec(ctx, deleteFavorite, arg.UserID, arg.DishID)
	return err
}

const getFavoriteByUserIdAndDishId = `-- name: GetFavoriteByUserIdAndDishId :one
SELECT dish_id, user_id, value
FROM favorites
WHERE user_id = $1 AND dish_id = $2
`

type GetFavoriteByUserIdAndDishIdParams struct {
	UserID uuid.UUID `json:"user_id"`
	DishID uuid.UUID `json:"dish_id"`
}

func (q *Queries) GetFavoriteByUserIdAndDishId(ctx context.Context, arg GetFavoriteByUserIdAndDishIdParams) (Favorite, error) {
	row := q.db.QueryRow(ctx, getFavoriteByUserIdAndDishId, arg.UserID, arg.DishID)
	var i Favorite
	err := row.Scan(&i.DishID, &i.UserID, &i.Value)
	return i, err
}

const updateFavorite = `-- name: UpdateFavorite :exec
UPDATE favorites
SET
  value = $3
WHERE user_id = $1 AND dish_id = $2
`

type UpdateFavoriteParams struct {
	UserID uuid.UUID `json:"user_id"`
	DishID uuid.UUID `json:"dish_id"`
	Value  int32     `json:"value"`
}

func (q *Queries) UpdateFavorite(ctx context.Context, arg UpdateFavoriteParams) error {
	_, err := q.db.Exec(ctx, updateFavorite, arg.UserID, arg.DishID, arg.Value)
	return err
}
