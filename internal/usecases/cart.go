package usecases

import (
	"context"

	"10.0.0.50/tuan.quang.tran/aioz-ads/models"
	"github.com/google/uuid"
)

type CartUseCase interface {
	AddCartItem(
		ctx context.Context,
		productVariantID uuid.UUID,
		quantity int32,
		authInfo models.AuthenticationInfo,
	) (*models.Cart, error)
	GetCartItemsByUserID(
		ctx context.Context,
		authInfo models.AuthenticationInfo,
		filter models.FilterParams,
	) ([]*models.Cart, int, error)
	UpdateCartItem(
		ctx context.Context,
		cartItemID uuid.UUID,
		quantity int32,
		authInfo models.AuthenticationInfo,
	) (*models.Cart, error)
	DeleteCartItem(
		ctx context.Context,
		cartItemID uuid.UUID,
		authInfo models.AuthenticationInfo,
	) error
}
