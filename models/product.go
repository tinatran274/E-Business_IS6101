package models

import (
	"context"
	"time"

	db "10.0.0.50/tuan.quang.tran/aioz-ads/db/generated"
	"github.com/google/uuid"
)

type ProductRepository interface {
	CreateProduct(
		ctx context.Context,
		product *Product,
	) (*Product, error)
	GetProductByID(
		ctx context.Context,
		id uuid.UUID,
	) (*Product, error)
	GetProducts(
		ctx context.Context,
		categoryId *uuid.UUID,
		filter FilterParams,
	) ([]*Product, error)
	CountProducts(
		ctx context.Context,
		categoryId *uuid.UUID,
		filter FilterParams,
	) (int, error)
	GetMyProducts(
		ctx context.Context,
		categoryId *uuid.UUID,
		createdBy uuid.UUID,
		filter FilterParams,
	) ([]*Product, error)
	CountMyProducts(
		ctx context.Context,
		categoryId *uuid.UUID,
		createdBy uuid.UUID,
		filter FilterParams,
	) (int, error)
	UpdateProduct(
		ctx context.Context,
		product *Product,
	) (*Product, error)
	DeleteProduct(
		ctx context.Context,
		product *Product,
	) error
}

type Product struct {
	ID              uuid.UUID         `json:"id"`
	Name            string            `json:"name"`
	Description     *string           `json:"description"`
	Brand           *string           `json:"brand"`
	Origin          *string           `json:"origin"`
	UserGuide       *string           `json:"user_guide"`
	CategoryID      uuid.UUID         `json:"category_id"`
	Category        *ProductCategory  `json:"product_category"`
	ProductVariants []*ProductVariant `json:"product_variants"`
	Status          string            `json:"status"`
	CreatedAt       time.Time         `json:"created_at"`
	CreatedBy       *uuid.UUID        `json:"created_by"`
	UpdatedAt       time.Time         `json:"updated_at"`
	UpdatedBy       *uuid.UUID        `json:"updated_by"`
	DeletedAt       *time.Time        `json:"deleted_at"`
	DeletedBy       *uuid.UUID        `json:"deleted_by"`
}

type CreateProductRequest struct {
	Name            string                                  `json:"name"`
	Description     *string                                 `json:"description"`
	Brand           *string                                 `json:"brand"`
	Origin          *string                                 `json:"origin"`
	UserGuide       *string                                 `json:"user_guide"`
	CategoryID      string                                  `json:"category_id"`
	ProductVariants []*CreateProductVariantInProductRequest `json:"product_variants"`
}

func NewProduct(
	name string,
	description *string,
	brand *string,
	origin *string,
	userGuide *string,
	categoryID uuid.UUID,
	createdBy *uuid.UUID,
) *Product {
	return &Product{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
		Brand:       brand,
		Origin:      origin,
		UserGuide:   userGuide,
		CategoryID:  categoryID,
		Status:      PendingStatus,
		CreatedAt:   time.Now().UTC(),
		CreatedBy:   createdBy,
		UpdatedAt:   time.Now().UTC(),
		UpdatedBy:   createdBy,
	}
}

func ToProduct(i db.Product) *Product {
	return &Product{
		ID:          i.ID,
		Name:        i.Name,
		Description: i.Description,
		Brand:       i.Brand,
		Origin:      i.Origin,
		UserGuide:   i.UserGuide,
		CategoryID:  i.CategoryID,
		Status:      i.Status,
		CreatedAt:   i.CreatedAt,
		CreatedBy:   i.CreatedBy,
		UpdatedAt:   i.UpdatedAt,
		UpdatedBy:   i.UpdatedBy,
		DeletedAt:   i.DeletedAt,
		DeletedBy:   i.DeletedBy,
	}
}
