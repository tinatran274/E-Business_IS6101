package models

import (
	"context"
	"time"

	db "10.0.0.50/tuan.quang.tran/aioz-ads/db/generated"
	"github.com/google/uuid"
)

type ProductVariantRepository interface {
	CreateProductVariant(
		ctx context.Context,
		productVariant *ProductVariant,
	) (*ProductVariant, error)
	GetProductVariantByID(
		ctx context.Context,
		id uuid.UUID,
	) (*ProductVariant, error)
	GetProductVariantsByProductID(
		ctx context.Context,
		productID uuid.UUID,
	) ([]*ProductVariant, error)
	UpdateProductVariant(
		ctx context.Context,
		productVariant *ProductVariant,
	) (*ProductVariant, error)
	DeleteProductVariant(
		ctx context.Context,
		productVariant *ProductVariant,
	) error
	IsProductVariantExist(
		ctx context.Context,
		productID uuid.UUID,
		color string,
	) (bool, error)
}

type ProductVariant struct {
	ID          uuid.UUID  `json:"id"`
	ProductID   uuid.UUID  `json:"product_id"`
	Description *string    `json:"description"`
	Color       string     `json:"color"`
	RetailPrice float64    `json:"retail_price"`
	Stock       int32      `json:"stock"`
	Status      string     `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	CreatedBy   *uuid.UUID `json:"created_by"`
	UpdatedAt   time.Time  `json:"updated_at"`
	UpdatedBy   *uuid.UUID `json:"updated_by"`
	DeletedAt   *time.Time `json:"deleted_at"`
	DeletedBy   *uuid.UUID `json:"deleted_by"`
}

type CreateProductVariantInProductRequest struct {
	Description *string `json:"description"`
	Color       string  `json:"color"`
	RetailPrice float64 `json:"retail_price"`
	Stock       int32   `json:"stock"`
}

type CreateProductVariantRequest struct {
	ProductID   string  `json:"product_id"`
	Description *string `json:"description"`
	Color       string  `json:"color"`
	RetailPrice float64 `json:"retail_price"`
	Stock       int32   `json:"stock"`
}

func NewProductVariant(
	productID uuid.UUID,
	description *string,
	color string,
	retailPrice float64,
	stock int32,
	createdBy *uuid.UUID,
) *ProductVariant {
	return &ProductVariant{
		ID:          uuid.New(),
		ProductID:   productID,
		Description: description,
		Color:       color,
		RetailPrice: retailPrice,
		Stock:       stock,
		Status:      PendingStatus,
		CreatedAt:   time.Now().UTC(),
		CreatedBy:   createdBy,
		UpdatedAt:   time.Now().UTC(),
		UpdatedBy:   createdBy,
	}
}

func ToProductVariant(i db.ProductVariant) *ProductVariant {
	return &ProductVariant{
		ID:          i.ID,
		ProductID:   i.ProductID,
		Description: i.Description,
		Color:       i.Color,
		RetailPrice: i.RetailPrice,
		Stock:       i.Stock,
		Status:      i.Status,
		CreatedAt:   i.CreatedAt,
		CreatedBy:   i.CreatedBy,
		UpdatedAt:   i.UpdatedAt,
		UpdatedBy:   i.UpdatedBy,
	}
}
