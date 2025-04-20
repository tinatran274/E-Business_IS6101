package models

import (
	"time"

	db "10.0.0.50/tuan.quang.tran/aioz-ads/db/generated"
	"github.com/google/uuid"
)

type CategoryRepository interface {
}

type Category struct {
	ID          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	Description *string    `json:"description"`
	Status      string     `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	CreatedBy   *uuid.UUID `json:"created_by"`
	UpdatedAt   time.Time  `json:"updated_at"`
	UpdatedBy   *uuid.UUID `json:"updated_by"`
	DeletedAt   *time.Time `json:"deleted_at"`
	DeletedBy   *uuid.UUID `json:"deleted_by"`
}

func NewCategory(
	name string,
	description *string,
) *Category {
	return &Category{
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

func ToCategory(c db.Category) *Category {
	return &Category{
		ID:          c.ID,
		Name:        c.Name,
		Description: c.Description,
		Status:      c.Status,
		CreatedAt:   c.CreatedAt,
		CreatedBy:   c.CreatedBy,
		UpdatedAt:   c.UpdatedAt,
		UpdatedBy:   c.UpdatedBy,
		DeletedAt:   c.DeletedAt,
	}
}
