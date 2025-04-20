package models

import (
	"time"

	db "10.0.0.50/tuan.quang.tran/aioz-ads/db/generated"
	"github.com/google/uuid"
)

type IngredientRepository interface {
}

type Ingredient struct {
	ID          uuid.UUID  `json:"id"`
	Email       string     `json:"email"`
	Name        string     `json:"name"`
	Description *string    `json:"description"`
	Removal     float64    `json:"removal"`
	Kcal        float64    `json:"kcal"`
	Protein     float64    `json:"protein"`
	Lipits      float64    `json:"lipits"`
	Glucids     float64    `json:"glucids"`
	Canxi       float64    `json:"canxi"`
	Phosphor    float64    `json:"phosphor"`
	Fe          float64    `json:"fe"`
	VitaminA    float64    `json:"vitamin_a"`
	VitaminB1   float64    `json:"vitamin_b1"`
	VitaminB2   float64    `json:"vitamin_b2"`
	VitaminC    float64    `json:"vitamin_c"`
	VitaminPp   float64    `json:"vitamin_pp"`
	BetaCaroten float64    `json:"beta_caroten"`
	CategoryID  uuid.UUID  `json:"category_id"`
	Status      string     `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	CreatedBy   *uuid.UUID `json:"created_by"`
	UpdatedAt   time.Time  `json:"updated_at"`
	UpdatedBy   *uuid.UUID `json:"updated_by"`
	DeletedAt   *time.Time `json:"deleted_at"`
	DeletedBy   *uuid.UUID `json:"deleted_by"`
}

func NewIngredient(
	name string,
	description *string,
	removal, kcal, protein, lipits, glucids, canxi, phosphor, fe,
	vitaminA, vitaminB1, vitaminB2, vitaminC, vitaminPp, betaCaroten float64,
	categoryID uuid.UUID,
) *Ingredient {
	return &Ingredient{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
		Removal:     removal,
		Kcal:        kcal,
		Protein:     protein,
		Lipits:      lipits,
		Glucids:     glucids,
		Canxi:       canxi,
		Phosphor:    phosphor,
		Fe:          fe,
		VitaminA:    vitaminA,
		VitaminB1:   vitaminB1,
		VitaminB2:   vitaminB2,
		VitaminC:    vitaminC,
		VitaminPp:   vitaminPp,
		BetaCaroten: betaCaroten,
		Status:      ActiveStatus,
		CreatedAt:   time.Now().UTC(),
		CreatedBy:   nil,
		UpdatedAt:   time.Now().UTC(),
		UpdatedBy:   nil,
		DeletedAt:   nil,
	}
}

func ToIngredient(i db.Ingredient) *Ingredient {
	return &Ingredient{
		ID:          i.ID,
		Name:        i.Name,
		Description: i.Description,
		Removal:     i.Removal,
		Kcal:        i.Kcal,
		Protein:     i.Protein,
		Lipits:      i.Lipits,
		Glucids:     i.Glucids,
		Canxi:       i.Canxi,
		Phosphor:    i.Phosphor,
		Fe:          i.Fe,
		VitaminA:    i.VitaminA,
		VitaminB1:   i.VitaminB1,
		VitaminB2:   i.VitaminB2,
		VitaminC:    i.VitaminC,
		VitaminPp:   i.VitaminPp,
		BetaCaroten: i.BetaCaroten,
		Status:      i.Status,
	}
}
