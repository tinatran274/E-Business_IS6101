package models

import (
	"context"

	db "10.0.0.50/tuan.quang.tran/aioz-ads/db/generated"
	"github.com/google/uuid"
)

type CartRepository interface {
	AddCartItem(
		ctx context.Context,
		cartItem *Cart,
	) (*Cart, error)
	GetCartItemByUserIdAndProductVariantId(
		ctx context.Context,
		userID, productVariantID uuid.UUID,
	) (*Cart, error)
	GetCartItemsByUserID(
		ctx context.Context,
		userID uuid.UUID,
		filter FilterParams,
	) ([]*Cart, error)
	CountCartItemsByUserID(
		ctx context.Context,
		userID uuid.UUID,
	) (int, error)
	DeleteCartItem(
		ctx context.Context,
		userID, productVariantID uuid.UUID,
	) error
	UpdateCartItem(
		ctx context.Context,
		cartItem *Cart,
	) (*Cart, error)
}

type Cart struct {
	UserID           uuid.UUID `json:"user_id"`
	ProductVariantID uuid.UUID `json:"product_variant_id"`
	Quantity         int32     `json:"quantity"`
}

func NewCart(
	userID uuid.UUID,
	productVariantID uuid.UUID,
	quantity int32,
) *Cart {
	return &Cart{
		UserID:           userID,
		ProductVariantID: productVariantID,
		Quantity:         quantity,
	}
}

func ToCart(cart *db.Cart) *Cart {
	return &Cart{
		UserID:           cart.UserID,
		ProductVariantID: cart.ProductVariantID,
		Quantity:         cart.Quantity,
	}
}
