package models

import (
	"time"

	db "10.0.0.50/tuan.quang.tran/aioz-ads/db/generated"
	"github.com/google/uuid"
)

type DishRepository interface {
}

type Dish struct {
	ID          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	Description *string    `json:"description"`
	CategoryID  uuid.UUID  `json:"category_id"`
	Status      string     `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	CreatedBy   *uuid.UUID `json:"created_by"`
	UpdatedAt   time.Time  `json:"updated_at"`
	UpdatedBy   *uuid.UUID `json:"updated_by"`
	DeletedAt   *time.Time `json:"deleted_at"`
	DeletedBy   *uuid.UUID `json:"deleted_by"`
}

func NewDish(
	name string,
	description *string,
	categoryID uuid.UUID,
) *Dish {
	return &Dish{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
		CategoryID:  categoryID,
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
		CategoryID:  d.CategoryID,
		Status:      d.Status,
		CreatedAt:   d.CreatedAt,
		CreatedBy:   d.CreatedBy,
		UpdatedAt:   d.UpdatedAt,
		UpdatedBy:   d.UpdatedBy,
		DeletedAt:   d.DeletedAt,
	}
}
