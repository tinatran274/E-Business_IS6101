package usecases

import (
	"context"

	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/utils/response"
	"10.0.0.50/tuan.quang.tran/aioz-ads/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type DishUseCase struct {
	dishRepo models.DishRepository
}

func NewDishUseCase(dishRepo models.DishRepository) *DishUseCase {
	return &DishUseCase{
		dishRepo: dishRepo,
	}
}

func (s *DishUseCase) GetDishByID(
	ctx context.Context,
	id uuid.UUID,
) (*models.Dish, error) {
	dish, err := s.dishRepo.GetDishByID(ctx, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, response.NewNotFoundError("Dish not found.")
		}

		return nil, response.NewInternalServerError(err)
	}
	return dish, nil
}

func (s *DishUseCase) GetDishes(
	ctx context.Context,
	filter models.FilterParams,
) ([]*models.Dish, int, error) {
	dishes, err := s.dishRepo.GetDishes(ctx, filter)
	if err != nil {
		return nil, 0, response.NewInternalServerError(err)
	}

	total, err := s.dishRepo.CountDishes(ctx, filter)
	if err != nil {
		return nil, 0, response.NewInternalServerError(err)
	}

	return dishes, total, nil
}
