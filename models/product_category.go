package models

import (
	"context"
	"time"

	db "10.0.0.50/tuan.quang.tran/aioz-ads/db/generated"
	"github.com/google/uuid"
)

type ProductCategoryRepository interface {
	GetProductCategories(
		ctx context.Context,
	) ([]*ProductCategory, error)
	GetProductCategoryByID(
		ctx context.Context,
		id uuid.UUID,
	) (*ProductCategory, error)
}

type ProductCategory struct {
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

func NewProductCategory(
	id uuid.UUID,
	name string,
	description *string,
	status string,
	createdBy *uuid.UUID,
) *ProductCategory {
	return &ProductCategory{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
		Status:      ActiveStatus,
		CreatedAt:   time.Now().UTC(),
		CreatedBy:   createdBy,
		UpdatedAt:   time.Now().UTC(),
		UpdatedBy:   createdBy,
	}
}

func ToProductCategory(i db.ProductCategory) *ProductCategory {
	return &ProductCategory{
		ID:          i.ID,
		Name:        i.Name,
		Description: i.Description,
		Status:      i.Status,
		CreatedAt:   i.CreatedAt,
		CreatedBy:   i.CreatedBy,
		UpdatedAt:   i.UpdatedAt,
		UpdatedBy:   i.UpdatedBy,
		DeletedAt:   i.DeletedAt,
		DeletedBy:   i.DeletedBy,
	}
}
