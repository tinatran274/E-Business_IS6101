package repositories

import (
	"context"
	"time"

	db "10.0.0.50/tuan.quang.tran/aioz-ads/db/generated"
	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/utils/metrics"
	"10.0.0.50/tuan.quang.tran/aioz-ads/models"
	"github.com/google/uuid"
)

type FavoriteRepository struct {
	db *db.Queries
}

func NewFavoriteRepository(db *db.Queries) *FavoriteRepository {
	return &FavoriteRepository{
		db: db,
	}
}

func (r *FavoriteRepository) CreateFavorite(
	ctx context.Context,
	favorite *models.Favorite,
) (*models.Favorite, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("CreateFavorite").
			Observe(time.Since(t).Seconds())
	}()

	err := r.db.CreateFavorite(ctx, db.CreateFavoriteParams{
		UserID: favorite.UserID,
		DishID: favorite.DishID,
		Value:  favorite.Value,
	})
	if err != nil {
		return nil, err
	}

	return favorite, nil
}

func (r *FavoriteRepository) DeleteFavorite(
	ctx context.Context,
	userId, dishId uuid.UUID,
) error {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("DeleteFavorite").
			Observe(time.Since(t).Seconds())
	}()

	err := r.db.DeleteFavorite(ctx, db.DeleteFavoriteParams{
		UserID: userId,
		DishID: dishId,
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *FavoriteRepository) UpdateFavorite(
	ctx context.Context,
	favorite *models.Favorite,
) (*models.Favorite, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("UpdateFavorite").
			Observe(time.Since(t).Seconds())
	}()

	err := r.db.UpdateFavorite(ctx, db.UpdateFavoriteParams{
		UserID: favorite.UserID,
		DishID: favorite.DishID,
		Value:  favorite.Value,
	})
	if err != nil {
		return nil, err
	}

	return favorite, nil
}

func (r *FavoriteRepository) GetFavoriteByUserIdAndDishId(
	ctx context.Context,
	userId, dishId uuid.UUID,
) (*models.Favorite, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("GetFavoriteByUserIdAndDishId").
			Observe(time.Since(t).Seconds())
	}()

	favorite, err := r.db.GetFavoriteByUserIdAndDishId(ctx, db.GetFavoriteByUserIdAndDishIdParams{
		UserID: userId,
		DishID: dishId,
	})
	if err != nil {
		return nil, err
	}

	return models.ToFavorite(favorite), nil
}
