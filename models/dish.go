package models

import (
	"context"
	"time"

	db "10.0.0.50/tuan.quang.tran/aioz-ads/db/generated"
	"github.com/google/uuid"
)

type DishRepository interface {
	GetDishByID(
		ctx context.Context,
		id uuid.UUID,
	) (*Dish, error)
	GetDishes(
		ctx context.Context,
		filter FilterParams,
	) ([]*Dish, error)
	CountDishes(
		ctx context.Context,
		filter FilterParams,
	) (int, error)
	GetDishesByIngredientID(
		ctx context.Context,
		id uuid.UUID,
		filter FilterParams,
	) ([]*Dish, error)
	CountDishesByIngredientID(
		ctx context.Context,
		id uuid.UUID,
		filter FilterParams,
	) (int, error)
}

type Dish struct {
	ID          uuid.UUID             `json:"id"`
	Name        string                `json:"name"`
	Description *string               `json:"description"`
	Kcal        float64               `json:"kcal"`
	Protein     float64               `json:"protein"`
	Lipits      float64               `json:"lipits"`
	Glucids     float64               `json:"glucids"`
	Canxi       float64               `json:"canxi"`
	Phosphor    float64               `json:"phosphor"`
	Fe          float64               `json:"fe"`
	VitaminA    float64               `json:"vitamin_a"`
	VitaminB1   float64               `json:"vitamin_b1"`
	VitaminB2   float64               `json:"vitamin_b2"`
	VitaminC    float64               `json:"vitamin_c"`
	VitaminPp   float64               `json:"vitamin_pp"`
	BetaCaroten float64               `json:"beta_caroten"`
	Category    Category              `json:"category"`
	Ingredients []*IngredientWithUnit `json:"ingredients"`
	Status      string                `json:"status"`
	CreatedAt   time.Time             `json:"created_at"`
	CreatedBy   *uuid.UUID            `json:"created_by"`
	UpdatedAt   time.Time             `json:"updated_at"`
	UpdatedBy   *uuid.UUID            `json:"updated_by"`
	DeletedAt   *time.Time            `json:"deleted_at"`
	DeletedBy   *uuid.UUID            `json:"deleted_by"`
}

func NewDish(
	name string,
	description *string,
) *Dish {
	return &Dish{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
		Status:      ActiveStatus,
		CreatedAt:   time.Now().UTC(),
		CreatedBy:   nil,
		UpdatedAt:   time.Now().UTC(),
		UpdatedBy:   nil,
		DeletedAt:   nil,
		DeletedBy:   nil,
	}
}

func ToDish(d db.Dish) *Dish {
	return &Dish{
		ID:          d.ID,
		Name:        d.Name,
		Description: d.Description,
		Status:      d.Status,
		CreatedAt:   d.CreatedAt,
		CreatedBy:   d.CreatedBy,
		UpdatedAt:   d.UpdatedAt,
		UpdatedBy:   d.UpdatedBy,
		DeletedAt:   d.DeletedAt,
	}
}
func CalculateGrams(ingredient IngredientWithUnit) float64 {
	switch ingredient.Category.Name {
	case "Grains":
		if ingredient.NutritionPer100g.Glucids != 0 {
			return 100 * ingredient.Unit * 20 / ingredient.NutritionPer100g.Glucids
		}
	case "Vegetables", "Fruits":
		return 80 * ingredient.Unit
	case "Protein":
		if ingredient.NutritionPer100g.Protein != 0 {
			return 100 * ingredient.Unit * 7 / ingredient.NutritionPer100g.Protein
		}
	case "Dairy":
		if ingredient.NutritionPer100g.Canxi != 0 {
			return 100 * ingredient.Unit * 100 / ingredient.NutritionPer100g.Canxi
		}
	case "Fats and oils":
		if ingredient.NutritionPer100g.Lipits != 0 {
			return 100 * ingredient.Unit * 5 / ingredient.NutritionPer100g.Lipits
		}
	case "Sugar":
		return 5 * ingredient.Unit
	case "Salt and sauces":
		return 1 * ingredient.Unit
	}
	return 0
}

func GetIngredientDetails(ingredient IngredientWithUnit) *IngredientWithUnit {
	grams := CalculateGrams(ingredient)
	return &IngredientWithUnit{
		Grams:       grams,
		Kcal:        grams * ingredient.NutritionPer100g.Kcal / 100,
		Protein:     grams * ingredient.NutritionPer100g.Protein / 100,
		Lipits:      grams * ingredient.NutritionPer100g.Lipits / 100,
		Glucids:     grams * ingredient.NutritionPer100g.Glucids / 100,
		Canxi:       grams * ingredient.NutritionPer100g.Canxi / 100,
		Phosphor:    grams * ingredient.NutritionPer100g.Phosphor / 100,
		Fe:          grams * ingredient.NutritionPer100g.Fe / 100,
		VitaminA:    grams * ingredient.NutritionPer100g.VitaminA / 100,
		BetaCaroten: grams * ingredient.NutritionPer100g.BetaCaroten / 100,
		VitaminB1:   grams * ingredient.NutritionPer100g.VitaminB1 / 100,
		VitaminB2:   grams * ingredient.NutritionPer100g.VitaminB2 / 100,
		VitaminPp:   grams * ingredient.NutritionPer100g.VitaminPp / 100,
		VitaminC:    grams * ingredient.NutritionPer100g.VitaminC / 100,
	}

}

func (d *Dish) CalculateNutritionalValues() *Dish {
	for i, recipe := range d.Ingredients {
		ingredientDetail := GetIngredientDetails(*recipe)
		d.Ingredients[i] = ingredientDetail
		d.Kcal += ingredientDetail.Kcal
		d.Protein += ingredientDetail.Protein
		d.Glucids += ingredientDetail.Glucids
		d.Lipits += ingredientDetail.Lipits
		d.Canxi += ingredientDetail.Canxi
		d.Phosphor += ingredientDetail.Phosphor
		d.Fe += ingredientDetail.Fe
		d.VitaminA += ingredientDetail.VitaminA
		d.BetaCaroten += ingredientDetail.BetaCaroten
		d.VitaminB1 += ingredientDetail.VitaminB1
		d.VitaminB2 += ingredientDetail.VitaminB2
		d.VitaminPp += ingredientDetail.VitaminPp
		d.VitaminC += ingredientDetail.VitaminC
	}

	return d
}
