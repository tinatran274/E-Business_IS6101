package usecases

import (
	"context"

	"10.0.0.50/tuan.quang.tran/aioz-ads/models"
	"github.com/google/uuid"
)

type IngredientUseCase interface {
	GetIngredientByID(ctx context.Context, id uuid.UUID) (*models.Ingredient, error)
	GetIngredients(ctx context.Context, filter models.FilterParams) ([]*models.Ingredient, int, error)
}
