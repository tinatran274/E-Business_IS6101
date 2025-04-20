package models

import "github.com/google/uuid"

type RecipeRepository interface {
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
