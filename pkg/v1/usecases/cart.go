package usecases

import (
	"context"

	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/utils/response"
	"10.0.0.50/tuan.quang.tran/aioz-ads/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type CartUseCase struct {
	CartRepo models.CartRepository
}

func NewCartUseCase(
	cartRepo models.CartRepository,
) *CartUseCase {
	return &CartUseCase{
		CartRepo: cartRepo,
	}
}

func (s *CartUseCase) AddCartItem(
	ctx context.Context,
	productVariantID uuid.UUID,
	quantity int32,
	authInfo models.AuthenticationInfo,
) (*models.Cart, error) {
	cartItem, err := s.CartRepo.GetCartItemByUserIdAndProductVariantId(
		ctx,
		authInfo.User.ID,
		productVariantID,
	)
	if err != nil && err != pgx.ErrNoRows {
		return nil, response.NewInternalServerError(err)
	}

	if cartItem == nil {
		cartItem = models.NewCart(
			authInfo.User.ID,
			productVariantID,
			quantity,
		)
		cartItem, err = s.CartRepo.AddCartItem(ctx, cartItem)
		if err != nil {
			return nil, response.NewInternalServerError(err)
		}

		return cartItem, nil
	}

	cartItem.Quantity += quantity
	cartItem, err = s.CartRepo.UpdateCartItem(ctx, cartItem)
	if err != nil {
		return nil, response.NewInternalServerError(err)
	}

	return cartItem, nil
}

func (s *CartUseCase) GetCartItemsByUserID(
	ctx context.Context,
	authInfo models.AuthenticationInfo,
	filter models.FilterParams,
) ([]*models.Cart, int, error) {
	cartItems, err := s.CartRepo.GetCartItemsByUserID(
		ctx,
		authInfo.User.ID,
		filter,
	)
	if err != nil {
		return nil, 0, response.NewInternalServerError(err)
	}

	count, err := s.CartRepo.CountCartItemsByUserID(
		ctx,
		authInfo.User.ID,
	)
	if err != nil {
		return nil, 0, response.NewInternalServerError(err)
	}

	return cartItems, count, nil
}

func (s *CartUseCase) UpdateCartItem(
	ctx context.Context,
	cartItemID uuid.UUID,
	quantity int32,
	authInfo models.AuthenticationInfo,
) (*models.Cart, error) {
	cartItem, err := s.CartRepo.GetCartItemByUserIdAndProductVariantId(
		ctx,
		authInfo.User.ID,
		cartItemID,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, response.NewNotFoundError("Cart item not found.")
		}

		return nil, response.NewInternalServerError(err)
	}

	cartItem.Quantity = quantity
	cartItem, err = s.CartRepo.UpdateCartItem(ctx, cartItem)
	if err != nil {
		return nil, response.NewInternalServerError(err)
	}

	return cartItem, nil
}

func (s *CartUseCase) DeleteCartItem(
	ctx context.Context,
	cartItemID uuid.UUID,
	authInfo models.AuthenticationInfo,
) error {
	_, err := s.CartRepo.GetCartItemByUserIdAndProductVariantId(
		ctx,
		authInfo.User.ID,
		cartItemID,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return response.NewNotFoundError("Cart item not found.")
		}

		return response.NewInternalServerError(err)
	}

	err = s.CartRepo.DeleteCartItem(
		ctx,
		authInfo.User.ID,
		cartItemID,
	)
	if err != nil {
		return response.NewInternalServerError(err)
	}

	return nil
}
