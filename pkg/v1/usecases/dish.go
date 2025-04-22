package usecases

import (
	"context"

	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/utils/response"
	"10.0.0.50/tuan.quang.tran/aioz-ads/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type DishUseCase struct {
	dishRepo     models.DishRepository
	favoriteRepo models.FavoriteRepository
}

func NewDishUseCase(
	dishRepo models.DishRepository,
	favoriteRepo models.FavoriteRepository,
) *DishUseCase {
	return &DishUseCase{
		dishRepo:     dishRepo,
		favoriteRepo: favoriteRepo,
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

func (s *DishUseCase) LikeDish(
	ctx context.Context,
	dishId uuid.UUID,
	authInfo models.AuthenticationInfo,
) error {
	dish, err := s.dishRepo.GetDishByID(ctx, dishId)
	if err != nil {
		if err == pgx.ErrNoRows {
			return response.NewNotFoundError("Dish not found.")
		}

		return response.NewInternalServerError(err)
	}

	existingFavorite, err := s.favoriteRepo.GetFavoriteByUserIdAndDishId(
		ctx,
		authInfo.User.ID,
		dish.ID,
	)
	if err != nil && err != pgx.ErrNoRows {
		return response.NewInternalServerError(err)
	}

	if existingFavorite == nil {
		favorite := &models.Favorite{
			UserID: authInfo.User.ID,
			DishID: dishId,
			Value:  1,
		}
		_, err = s.favoriteRepo.CreateFavorite(ctx, favorite)
		if err != nil {
			return response.NewInternalServerError(err)
		}

		return nil
	}

	if existingFavorite.Value == 1 {
		return response.NewBadRequestError("Dish already liked.")
	}

	favorite := &models.Favorite{
		UserID: authInfo.User.ID,
		DishID: dish.ID,
		Value:  1,
	}
	_, err = s.favoriteRepo.UpdateFavorite(ctx, favorite)
	if err != nil {
		return response.NewInternalServerError(err)
	}

	return nil
}

func (s *DishUseCase) UnlikeDish(
	ctx context.Context,
	dishId uuid.UUID,
	authInfo models.AuthenticationInfo,
) error {
	dish, err := s.dishRepo.GetDishByID(ctx, dishId)
	if err != nil {
		if err == pgx.ErrNoRows {
			return response.NewNotFoundError("Dish not found.")
		}

		return response.NewInternalServerError(err)
	}

	existingFavorite, err := s.favoriteRepo.GetFavoriteByUserIdAndDishId(
		ctx,
		authInfo.User.ID,
		dish.ID,
	)
	if err != nil && err != pgx.ErrNoRows {
		return response.NewInternalServerError(err)
	}

	if existingFavorite == nil {
		return response.NewBadRequestError("Dish not liked yet.")
	}

	if existingFavorite.Value == 0 {
		return response.NewBadRequestError("Dish already unliked.")
	}

	favorite := &models.Favorite{
		UserID: authInfo.User.ID,
		DishID: dish.ID,
		Value:  0,
	}
	_, err = s.favoriteRepo.UpdateFavorite(ctx, favorite)
	if err != nil {
		return response.NewInternalServerError(err)
	}

	return nil
}
