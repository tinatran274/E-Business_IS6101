package repositories

import (
	"context"
	"time"

	db "10.0.0.50/tuan.quang.tran/aioz-ads/db/generated"
	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/utils/metrics"
	"10.0.0.50/tuan.quang.tran/aioz-ads/models"
	"github.com/google/uuid"
)

type ProductCategoryRepository struct {
	db *db.Queries
}

func NewProductCategoryRepository(db *db.Queries) *ProductCategoryRepository {
	return &ProductCategoryRepository{
		db: db,
	}
}

func (r *ProductCategoryRepository) GetProductCategories(
	ctx context.Context,
) ([]*models.ProductCategory, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("CreateProductVariant").
			Observe(time.Since(t).Seconds())
	}()

	productCategories, err := r.db.GetProductCategories(ctx)
	if err != nil {
		return nil, err
	}

	productCategoryModels := make([]*models.ProductCategory, len(productCategories))
	for i, category := range productCategories {
		productCategoryModels[i] = models.ToProductCategory(category)
	}

	return productCategoryModels, nil
}

func (r *ProductCategoryRepository) GetProductCategoryByID(
	ctx context.Context,
	id uuid.UUID,
) (*models.ProductCategory, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("GetProductCategoryByID").
			Observe(time.Since(t).Seconds())
	}()

	productCategory, err := r.db.GetProductCategoryById(ctx, id)
	if err != nil {
		return nil, err
	}

	return models.ToProductCategory(productCategory), nil
}
