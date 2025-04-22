package models

import (
	"context"

	"github.com/google/uuid"
)

type RecipeRepository interface {
	CreateRecipe(
		ctx context.Context,
		recipe *Recipe,
	) (*Recipe, error)
	DeleteRecipe(
		ctx context.Context,
		dishId, ingredientId uuid.UUID,
	) error
}

type Recipe struct {
	DishID       uuid.UUID `json:"dish_id"`
	IngredientID uuid.UUID `json:"ingredient_id"`
	Unit         float64   `json:"unit"`
}

func NewRecipe(
	dishID, ingredientID uuid.UUID,
	unit float64,
) *Recipe {
	return &Recipe{
		DishID:       dishID,
		IngredientID: ingredientID,
		Unit:         unit,
	}
}

func ToRecipe(r Recipe) *Recipe {
	return &Recipe{
		DishID:       r.DishID,
		IngredientID: r.IngredientID,
		Unit:         r.Unit,
	}
}
