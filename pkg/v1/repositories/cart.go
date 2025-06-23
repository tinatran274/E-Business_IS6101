package repositories

import (
	"context"
	"time"

	db "10.0.0.50/tuan.quang.tran/aioz-ads/db/generated"
	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/utils/metrics"
	"10.0.0.50/tuan.quang.tran/aioz-ads/models"
	"github.com/google/uuid"
)

type CartRepository struct {
	db *db.Queries
}

func NewCartRepository(db *db.Queries) *CartRepository {
	return &CartRepository{
		db: db,
	}
}

func (r *CartRepository) AddCartItem(
	ctx context.Context,
	cartItem *models.Cart,
) (*models.Cart, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("AddCartItem").
			Observe(time.Since(t).Seconds())
	}()

	err := r.db.AddCartItem(ctx, db.AddCartItemParams{
		UserID:           cartItem.UserID,
		ProductVariantID: cartItem.ProductVariantID,
		Quantity:         cartItem.Quantity,
	})
	if err != nil {
		return nil, err
	}

	return cartItem, nil
}

func (r *CartRepository) GetCartItemByUserIdAndProductVariantId(
	ctx context.Context,
	userID, productVariantID uuid.UUID,
) (*models.Cart, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("GetCartItemByUserIdAndProductVariantId").
			Observe(time.Since(t).Seconds())
	}()

	cartItem, err := r.db.GetCartItemByUserIdAndProductVariantId(
		ctx,
		db.GetCartItemByUserIdAndProductVariantIdParams{
			UserID:           userID,
			ProductVariantID: productVariantID,
		})
	if err != nil {
		return nil, err
	}

	return models.ToCart(&cartItem), nil
}

func (r *CartRepository) GetCartItemsByUserID(
	ctx context.Context,
	userID uuid.UUID,
	filter models.FilterParams,
) ([]*models.Cart, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("GetCartItemsByUserID").
			Observe(time.Since(t).Seconds())
	}()

	cartItems, err := r.db.GetCartItemsByUserId(
		ctx,
		db.GetCartItemsByUserIdParams{
			UserID: userID,
			Limit:  filter.Limit,
			Offset: filter.Offset,
		},
	)
	if err != nil {
		return nil, err
	}

	var result []*models.Cart
	for _, item := range cartItems {
		result = append(result, models.ToCart(&item))
	}

	return result, nil
}

func (r *CartRepository) CountCartItemsByUserID(
	ctx context.Context,
	userID uuid.UUID,
) (int, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("CountCartItemsByUserID").
			Observe(time.Since(t).Seconds())
	}()

	count, err := r.db.CountCartItemsByUserId(ctx, userID)
	if err != nil {
		return 0, err
	}

	return int(count), nil
}

func (r *CartRepository) DeleteCartItem(
	ctx context.Context,
	userID, productVariantID uuid.UUID,
) error {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("DeleteCartItem").
			Observe(time.Since(t).Seconds())
	}()

	err := r.db.DeleteCartItem(ctx, db.DeleteCartItemParams{
		UserID:           userID,
		ProductVariantID: productVariantID,
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *CartRepository) UpdateCartItem(
	ctx context.Context,
	cartItem *models.Cart,
) (*models.Cart, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("UpdateCartItem").
			Observe(time.Since(t).Seconds())
	}()

	err := r.db.UpdateCartItem(ctx, db.UpdateCartItemParams{
		UserID:           cartItem.UserID,
		ProductVariantID: cartItem.ProductVariantID,
		Quantity:         cartItem.Quantity,
	})
	if err != nil {
		return nil, err
	}

	return cartItem, nil
}
