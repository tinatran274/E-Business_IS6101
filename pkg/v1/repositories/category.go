package repositories

import (
	"context"
	"time"

	db "10.0.0.50/tuan.quang.tran/aioz-ads/db/generated"
	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/utils/metrics"
	"10.0.0.50/tuan.quang.tran/aioz-ads/models"
	"github.com/google/uuid"
)

type CategoryRepository struct {
	db *db.Queries
}

func NewCategoryRepository(db *db.Queries) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

func (r *CategoryRepository) GetCategories(
	ctx context.Context,
) ([]*models.Category, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("GetCategories").
			Observe(time.Since(t).Seconds())
	}()

	categories, err := r.db.GetCategories(ctx)
	if err != nil {
		return nil, err
	}

	categoryModels := make([]*models.Category, len(categories))
	for i, category := range categories {
		categoryModels[i] = models.ToCategory(category)
	}

	return categoryModels, nil
}

func (r *CategoryRepository) GetCategoryByID(
	ctx context.Context,
	id uuid.UUID,
) (*models.Category, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("GetCategoryByID").
			Observe(time.Since(t).Seconds())
	}()

	category, err := r.db.GetCategoryById(ctx, id)
	if err != nil {
		return nil, err
	}

	return models.ToCategory(category), nil
}
