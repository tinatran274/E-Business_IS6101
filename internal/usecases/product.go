package usecases

import (
	"context"

	"10.0.0.50/tuan.quang.tran/aioz-ads/models"
	"github.com/google/uuid"
)

type ProductUseCase interface {
	CreateProduct(
		ctx context.Context,
		payload models.CreateProductRequest,
		authInfo models.AuthenticationInfo,
	) (*models.Product, error)
	GetProductByID(
		ctx context.Context,
		id uuid.UUID,
	) (*models.Product, error)
	GetProducts(
		ctx context.Context,
		categoryId *uuid.UUID,
		filter models.FilterParams,
	) ([]*models.Product, int, error)
	GetMyProducts(
		ctx context.Context,
		categoryId *uuid.UUID,
		createdBy uuid.UUID,
		filter models.FilterParams,
	) ([]*models.Product, int, error)
	UpdateProduct(
		ctx context.Context,
		id uuid.UUID,
		payload models.CreateProductRequest,
		authInfo models.AuthenticationInfo,
	) (*models.Product, error)
	DeleteProduct(
		ctx context.Context,
		id uuid.UUID,
		authInfo models.AuthenticationInfo,
	) error
	ApproveProduct(
		ctx context.Context,
		id uuid.UUID,
		authInfo models.AuthenticationInfo,
	) error
	RejectProduct(
		ctx context.Context,
		id uuid.UUID,
		authInfo models.AuthenticationInfo,
	) error
}
