package repositories

import (
	"context"
	"time"

	db "10.0.0.50/tuan.quang.tran/aioz-ads/db/generated"
	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/utils/metrics"
	"10.0.0.50/tuan.quang.tran/aioz-ads/models"
	"github.com/google/uuid"
)

type ProductRepository struct {
	db *db.Queries
}

func NewProductRepository(db *db.Queries) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) CreateProduct(
	ctx context.Context,
	product *models.Product,
) (*models.Product, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("CreateProduct").
			Observe(time.Since(t).Seconds())
	}()

	err := r.db.CreateProduct(ctx, db.CreateProductParams{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Brand:       product.Brand,
		Origin:      product.Origin,
		CategoryID:  product.CategoryID,
		UserGuide:   product.UserGuide,
		Status:      product.Status,
		CreatedAt:   product.CreatedAt,
		CreatedBy:   product.CreatedBy,
		UpdatedAt:   product.UpdatedAt,
		UpdatedBy:   product.UpdatedBy,
	})
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (r *ProductRepository) GetProductByID(
	ctx context.Context,
	id uuid.UUID,
) (*models.Product, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("GetProductByID").
			Observe(time.Since(t).Seconds())
	}()

	product, err := r.db.GetProductById(ctx, id)
	if err != nil {
		return nil, err
	}

	category, err := r.db.GetProductCategoryById(ctx, product.CategoryID)
	if err != nil {
		return nil, err
	}

	productVariants, err := r.db.GetProductVariantsByProductId(ctx, id)
	if err != nil {
		return nil, err
	}

	productVariantsModels := make([]*models.ProductVariant, len(productVariants))
	for i, productVariant := range productVariants {
		productVariantsModels[i] = models.ToProductVariant(productVariant)
	}

	productModel := models.ToProduct(product)
	productModel.Category = models.ToProductCategory(category)
	productModel.ProductVariants = productVariantsModels
	return productModel, nil
}

func (r *ProductRepository) GetProducts(
	ctx context.Context,
	cId *uuid.UUID,
	filter models.FilterParams,
) ([]*models.Product, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("GetProducts").
			Observe(time.Since(t).Seconds())
	}()

	var categoryId string
	if cId == nil {
		categoryId = ""
	} else {
		categoryId = cId.String()
	}

	products, err := r.db.GetProducts(ctx, db.GetProductsParams{
		Status:     filter.Status,
		Limit:      filter.Limit,
		Offset:     filter.Offset,
		Keyword:    filter.Keyword,
		SortBy:     filter.SortBy,
		OrderBy:    filter.OrderBy,
		CategoryID: categoryId,
	})
	if err != nil {
		return nil, err
	}

	result := make([]*models.Product, len(products))
	for i, product := range products {
		category, err := r.db.GetProductCategoryById(ctx, product.CategoryID)
		if err != nil {
			return nil, err
		}

		result[i] = models.ToProduct(product)
		result[i].Category = models.ToProductCategory(category)
	}

	return result, nil
}

func (r *ProductRepository) GetMyProducts(
	ctx context.Context,
	cid *uuid.UUID,
	createdBy uuid.UUID,
	filter models.FilterParams,
) ([]*models.Product, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("GetMyProducts").
			Observe(time.Since(t).Seconds())
	}()

	var categoryId string
	if cid == nil {
		categoryId = ""
	} else {
		categoryId = cid.String()
	}

	products, err := r.db.GetMyProducts(ctx, db.GetMyProductsParams{
		Status:     filter.Status,
		Limit:      filter.Limit,
		Offset:     filter.Offset,
		Keyword:    filter.Keyword,
		SortBy:     filter.SortBy,
		OrderBy:    filter.OrderBy,
		CategoryID: categoryId,
		CreatedBy:  &createdBy,
	})
	if err != nil {
		return nil, err
	}

	result := make([]*models.Product, len(products))
	for i, product := range products {
		category, err := r.db.GetProductCategoryById(ctx, product.CategoryID)
		if err != nil {
			return nil, err
		}

		result[i] = models.ToProduct(product)
		result[i].Category = models.ToProductCategory(category)
	}

	return result, nil
}

func (r *ProductRepository) CountProducts(
	ctx context.Context,
	cId *uuid.UUID,
	filter models.FilterParams,
) (int, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("CountProducts").
			Observe(time.Since(t).Seconds())
	}()

	var categoryId string
	if cId == nil {
		categoryId = ""
	} else {
		categoryId = cId.String()
	}

	count, err := r.db.CountProducts(
		ctx,
		db.CountProductsParams{
			Status:     filter.Status,
			Keyword:    filter.Keyword,
			CategoryID: categoryId,
		},
	)
	if err != nil {
		return 0, err
	}

	return int(count), nil
}

func (r *ProductRepository) CountMyProducts(
	ctx context.Context,
	cId *uuid.UUID,
	createdBy uuid.UUID,
	filter models.FilterParams,
) (int, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("CountProducts").
			Observe(time.Since(t).Seconds())
	}()

	var categoryId string
	if cId == nil {
		categoryId = ""
	} else {
		categoryId = cId.String()
	}

	count, err := r.db.CountMyProducts(
		ctx,
		db.CountMyProductsParams{
			Status:     filter.Status,
			Keyword:    filter.Keyword,
			CategoryID: categoryId,
			CreatedBy:  &createdBy,
		},
	)
	if err != nil {
		return 0, err
	}

	return int(count), nil
}

func (r *ProductRepository) GetProductsByCategoryID(
	ctx context.Context,
	id uuid.UUID,
	filter models.FilterParams,
) ([]*models.Product, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("GetProductsByCategoryID").
			Observe(time.Since(t).Seconds())
	}()

	arg := db.GetProductsByCategoryIdParams{
		CategoryID: id,
		Keyword:    filter.Keyword,
		Limit:      filter.Limit,
		Offset:     filter.Offset,
		SortBy:     filter.SortBy,
		OrderBy:    filter.OrderBy,
	}
	products, err := r.db.GetProductsByCategoryId(ctx, arg)
	if err != nil {
		return nil, err
	}

	result := make([]*models.Product, len(products))
	for i, product := range products {
		category, err := r.db.GetProductCategoryById(ctx, product.CategoryID)
		if err != nil {
			return nil, err
		}

		result[i] = models.ToProduct(product)
		result[i].Category = models.ToProductCategory(category)
	}

	return result, nil
}

func (r *ProductRepository) CountProductsByCategoryID(
	ctx context.Context,
	id uuid.UUID,
	filter models.FilterParams,
) (int, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("CountProductsByCategoryID").
			Observe(time.Since(t).Seconds())
	}()

	count, err := r.db.CountProductsByCategoryId(
		ctx,
		db.CountProductsByCategoryIdParams{
			CategoryID: id,
			Keyword:    filter.Keyword,
		},
	)
	if err != nil {
		return 0, err
	}

	return int(count), nil
}

func (r *ProductRepository) UpdateProduct(
	ctx context.Context,
	product *models.Product,
) (*models.Product, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("UpdateProduct").
			Observe(time.Since(t).Seconds())
	}()

	err := r.db.UpdateProduct(ctx, db.UpdateProductParams{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Brand:       product.Brand,
		Origin:      product.Origin,
		UserGuide:   product.UserGuide,
		CategoryID:  product.CategoryID,
		Status:      product.Status,
		UpdatedAt:   product.UpdatedAt,
		UpdatedBy:   product.UpdatedBy,
	})
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (r *ProductRepository) DeleteProduct(
	ctx context.Context,
	product *models.Product,
) error {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("DeleteProduct").
			Observe(time.Since(t).Seconds())
	}()

	err := r.db.DeleteProduct(
		ctx,
		db.DeleteProductParams{
			ID:        product.ID,
			DeletedAt: product.DeletedAt,
			DeletedBy: product.DeletedBy,
		},
	)
	if err != nil {
		return err
	}

	return nil
}
