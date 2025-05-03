package usecases

import (
	"context"

	"10.0.0.50/tuan.quang.tran/aioz-ads/models"
	"github.com/google/uuid"
)

type DishUseCase interface {
	GetDishByID(ctx context.Context, id uuid.UUID) (*models.Dish, error)
	GetDishesByIngredientID(
		ctx context.Context,
		id uuid.UUID,
		filter models.FilterParams,
	) ([]*models.Dish, int, error)
	GetDishes(ctx context.Context, filter models.FilterParams) ([]*models.Dish, int, error)
	LikeDish(
		ctx context.Context,
		dishId uuid.UUID,
		authInfo models.AuthenticationInfo,
	) error
	UnlikeDish(
		ctx context.Context,
		dishId uuid.UUID,
		authInfo models.AuthenticationInfo,
	) error
}
