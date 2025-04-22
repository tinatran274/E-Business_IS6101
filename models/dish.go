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
}

type Dish struct {
	ID          uuid.UUID    `json:"id"`
	Name        string       `json:"name"`
	Description *string      `json:"description"`
	Kcal        float64      `json:"kcal"`
	Protein     float64      `json:"protein"`
	Lipits      float64      `json:"lipits"`
	Glucids     float64      `json:"glucids"`
	Canxi       float64      `json:"canxi"`
	Phosphor    float64      `json:"phosphor"`
	Fe          float64      `json:"fe"`
	VitaminA    float64      `json:"vitamin_a"`
	VitaminB1   float64      `json:"vitamin_b1"`
	VitaminB2   float64      `json:"vitamin_b2"`
	VitaminC    float64      `json:"vitamin_c"`
	VitaminPp   float64      `json:"vitamin_pp"`
	BetaCaroten float64      `json:"beta_caroten"`
	Category    Category     `json:"category"`
	Ingredients []Ingredient `json:"ingredients"`
	Status      string       `json:"status"`
	CreatedAt   time.Time    `json:"created_at"`
	CreatedBy   *uuid.UUID   `json:"created_by"`
	UpdatedAt   time.Time    `json:"updated_at"`
	UpdatedBy   *uuid.UUID   `json:"updated_by"`
	DeletedAt   *time.Time   `json:"deleted_at"`
	DeletedBy   *uuid.UUID   `json:"deleted_by"`
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

func (d *Dish) CalculateNutritionalValues() *Dish {
	for _, recipe := range d.Ingredients {
		d.Kcal += recipe.Kcal
		d.Protein += recipe.Protein
		d.Glucids += recipe.Glucids
		d.Lipits += recipe.Lipits
		d.Canxi += recipe.Canxi
		d.Phosphor += recipe.Phosphor
		d.Fe += recipe.Fe
		d.VitaminA += recipe.VitaminA
		d.BetaCaroten += recipe.BetaCaroten
		d.VitaminB1 += recipe.VitaminB1
		d.VitaminB2 += recipe.VitaminB2
		d.VitaminPp += recipe.VitaminPp
		d.VitaminC += recipe.VitaminC
	}

	return d
}
