package usecases

import (
	"context"

	"10.0.0.50/tuan.quang.tran/aioz-ads/models"
	"github.com/google/uuid"
)

type ProducetVariantUseCase interface {
	CreateProductVariant(
		ctx context.Context,
		payload models.CreateProductVariantRequest,
		authInfo models.AuthenticationInfo,
	) (*models.ProductVariant, error)
	GetProductVariantsByProductID(
		ctx context.Context,
		productID uuid.UUID,
	) ([]*models.ProductVariant, error)
	UpdateProductVariant(
		ctx context.Context,
		productVariantID uuid.UUID,
		payload models.CreateProductVariantRequest,
		authInfo models.AuthenticationInfo,
	) (*models.ProductVariant, error)
	DeleteProductVariant(
		ctx context.Context,
		productVariantID uuid.UUID,
		authInfo models.AuthenticationInfo,
	) error
}
