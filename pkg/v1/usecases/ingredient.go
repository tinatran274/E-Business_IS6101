package usecases

import (
	"context"

	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/utils/response"
	"10.0.0.50/tuan.quang.tran/aioz-ads/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type IngredientUseCase struct {
	ingredientRepo models.IngredientRepository
}

func NewIngredientUseCase(ingredientRepo models.IngredientRepository) *IngredientUseCase {
	return &IngredientUseCase{
		ingredientRepo: ingredientRepo,
	}
}

func (s *IngredientUseCase) GetIngredientByID(
	ctx context.Context,
	id uuid.UUID,
) (*models.Ingredient, error) {
	ingredient, err := s.ingredientRepo.GetIngredientByID(ctx, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, response.NewNotFoundError("Ingredient not found.")
		}

		return nil, response.NewInternalServerError(err)
	}
	return ingredient, nil
}

func (s *IngredientUseCase) GetIngredients(
	ctx context.Context,
	filter models.FilterParams,
) ([]*models.Ingredient, int, error) {
	ingredients, err := s.ingredientRepo.GetIngredients(ctx, filter)
	if err != nil {
		return nil, 0, response.NewInternalServerError(err)
	}

	total, err := s.ingredientRepo.CountIngredients(ctx, filter)
	if err != nil {
		return nil, 0, response.NewInternalServerError(err)
	}

	return ingredients, total, nil
}
