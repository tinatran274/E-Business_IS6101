package models

import (
	"context"

	db "10.0.0.50/tuan.quang.tran/aioz-ads/db/generated"
	"github.com/google/uuid"
)

type FavoriteRepository interface {
	CreateFavorite(
		ctx context.Context,
		favorite *Favorite,
	) (*Favorite, error)
	DeleteFavorite(
		ctx context.Context,
		userId, dishId uuid.UUID,
	) error
	UpdateFavorite(
		ctx context.Context,
		favorite *Favorite,
	) (*Favorite, error)
	GetFavoriteByUserIdAndDishId(
		ctx context.Context,
		userId, dishId uuid.UUID,
	) (*Favorite, error)
}

type Favorite struct {
	UserID uuid.UUID `json:"user_id"`
	DishID uuid.UUID `json:"dish_id"`
	Value  int32     `json:"value"`
}

func NewFavorite(
	userID, dishID uuid.UUID,
	value int32,
) *Favorite {
	return &Favorite{
		UserID: userID,
		DishID: dishID,
		Value:  value,
	}
}

func ToFavorite(f db.Favorite) *Favorite {
	return &Favorite{
		UserID: f.UserID,
		DishID: f.DishID,
		Value:  f.Value,
	}
}
