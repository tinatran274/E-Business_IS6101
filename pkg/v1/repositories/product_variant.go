package repositories

import (
	"context"
	"time"

	db "10.0.0.50/tuan.quang.tran/aioz-ads/db/generated"
	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/utils/metrics"
	"10.0.0.50/tuan.quang.tran/aioz-ads/models"
	"github.com/google/uuid"
)

type ProductVariantRepository struct {
	db *db.Queries
}

func NewProductVariantRepository(db *db.Queries) *ProductVariantRepository {
	return &ProductVariantRepository{
		db: db,
	}
}

func (r *ProductVariantRepository) CreateProductVariant(
	ctx context.Context,
	productVariant *models.ProductVariant,
) (*models.ProductVariant, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("CreateProductVariant").
			Observe(time.Since(t).Seconds())
	}()

	err := r.db.CreateProductVariant(ctx, db.CreateProductVariantParams{
		ID:          productVariant.ID,
		ProductID:   productVariant.ProductID,
		Description: productVariant.Description,
		Color:       productVariant.Color,
		RetailPrice: productVariant.RetailPrice,
		Stock:       productVariant.Stock,
		Status:      productVariant.Status,
		CreatedAt:   productVariant.CreatedAt,
		CreatedBy:   productVariant.CreatedBy,
		UpdatedAt:   productVariant.UpdatedAt,
		UpdatedBy:   productVariant.UpdatedBy,
	})
	if err != nil {
		return nil, err
	}

	return productVariant, nil
}

func (r *ProductVariantRepository) GetProductVariantByID(
	ctx context.Context,
	id uuid.UUID,
) (*models.ProductVariant, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("GetProductVariantByID").
			Observe(time.Since(t).Seconds())
	}()

	productVariant, err := r.db.GetProductVariantById(ctx, id)
	if err != nil {
		return nil, err
	}

	return models.ToProductVariant(productVariant), nil
}

func (r *ProductVariantRepository) GetProductVariantsByProductID(
	ctx context.Context,
	productID uuid.UUID,
) ([]*models.ProductVariant, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("GetProductVariantsByProductID").
			Observe(time.Since(t).Seconds())
	}()

	productVariants, err := r.db.GetProductVariantsByProductId(
		ctx,
		productID,
	)
	if err != nil {
		return nil, err
	}

	var result []*models.ProductVariant
	for _, pv := range productVariants {
		result = append(result, models.ToProductVariant(pv))
	}

	return result, nil
}

func (r *ProductVariantRepository) UpdateProductVariant(
	ctx context.Context,
	productVariant *models.ProductVariant,
) (*models.ProductVariant, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("UpdateProductVariant").
			Observe(time.Since(t).Seconds())
	}()

	err := r.db.UpdateProductVariant(ctx, db.UpdateProductVariantParams{
		ID:          productVariant.ID,
		ProductID:   productVariant.ProductID,
		Description: productVariant.Description,
		Color:       productVariant.Color,
		RetailPrice: productVariant.RetailPrice,
		Stock:       productVariant.Stock,
		Status:      productVariant.Status,
		UpdatedAt:   productVariant.UpdatedAt,
		UpdatedBy:   productVariant.UpdatedBy,
	})
	if err != nil {
		return nil, err
	}

	return productVariant, nil
}

func (r *ProductVariantRepository) DeleteProductVariant(
	ctx context.Context,
	productVariant *models.ProductVariant,
) error {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("DeleteProductVariant").
			Observe(time.Since(t).Seconds())
	}()

	err := r.db.DeleteProductVariant(
		ctx,
		db.DeleteProductVariantParams{
			ID:        productVariant.ID,
			DeletedAt: productVariant.DeletedAt,
			DeletedBy: productVariant.DeletedBy,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *ProductVariantRepository) IsProductVariantExist(
	ctx context.Context,
	productID uuid.UUID,
	color string,
) (bool, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("IsProductVariantExist").
			Observe(time.Since(t).Seconds())
	}()

	exist, err := r.db.IsColorExist(ctx, db.IsColorExistParams{
		ProductID: productID,
		Color:     color,
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}
