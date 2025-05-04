package usecases

import (
	"context"
	"time"

	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/utils/response"
	"10.0.0.50/tuan.quang.tran/aioz-ads/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type ProductUseCase struct {
	productRepo         models.ProductRepository
	productVariantRepo  models.ProductVariantRepository
	productCategoryRepo models.ProductCategoryRepository
}

func NewProductUseCase(
	productRepo models.ProductRepository,
	productVariantRepo models.ProductVariantRepository,
	productCategoryRepo models.ProductCategoryRepository,
) *ProductUseCase {
	return &ProductUseCase{
		productRepo:         productRepo,
		productVariantRepo:  productVariantRepo,
		productCategoryRepo: productCategoryRepo,
	}
}

func (s *ProductUseCase) CreateProduct(
	ctx context.Context,
	payload models.CreateProductRequest,
	authInfo models.AuthenticationInfo,
) (*models.Product, error) {
	productCategoryId, err := uuid.Parse(payload.CategoryID)
	if err != nil {
		return nil, response.NewBadRequestError("Invalid product category ID.")
	}

	_, err = s.productCategoryRepo.GetProductCategoryByID(ctx, productCategoryId)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, response.NewBadRequestError("Product category not found.")
		}

		return nil, response.NewInternalServerError(err)
	}

	product := models.NewProduct(
		payload.Name,
		payload.Description,
		payload.Brand,
		payload.Origin,
		payload.UserGuide,
		productCategoryId,
		&authInfo.User.ID,
	)
	product, err = s.productRepo.CreateProduct(ctx, product)
	if err != nil {
		return nil, response.NewInternalServerError(err)
	}

	for _, variant := range payload.ProductVariants {
		productVariant := models.NewProductVariant(
			product.ID,
			variant.Description,
			variant.Color,
			variant.RetailPrice,
			variant.Stock,
			&authInfo.User.ID,
		)

		_, err = s.productVariantRepo.CreateProductVariant(ctx, productVariant)
		if err != nil {
			return nil, response.NewInternalServerError(err)
		}
	}

	return product, nil
}

func (s *ProductUseCase) GetProductByID(
	ctx context.Context,
	id uuid.UUID,
) (*models.Product, error) {
	product, err := s.productRepo.GetProductByID(ctx, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, response.NewNotFoundError("Product not found.")
		}

		return nil, response.NewInternalServerError(err)
	}

	return product, nil
}

func (s *ProductUseCase) GetProducts(
	ctx context.Context,
	categoryId *uuid.UUID,
	filter models.FilterParams,
) ([]*models.Product, int, error) {
	products, err := s.productRepo.GetProducts(ctx, categoryId, filter)
	if err != nil {
		return nil, 0, response.NewInternalServerError(err)
	}

	total, err := s.productRepo.CountProducts(ctx, categoryId, filter)
	if err != nil {
		return nil, 0, response.NewInternalServerError(err)
	}

	return products, total, nil
}

func (s *ProductUseCase) GetMyProducts(
	ctx context.Context,
	categoryId *uuid.UUID,
	createdBy uuid.UUID,
	filter models.FilterParams,
) ([]*models.Product, int, error) {
	products, err := s.productRepo.GetMyProducts(ctx, categoryId, createdBy, filter)
	if err != nil {
		return nil, 0, response.NewInternalServerError(err)
	}

	total, err := s.productRepo.CountMyProducts(ctx, categoryId, createdBy, filter)
	if err != nil {
		return nil, 0, response.NewInternalServerError(err)
	}

	return products, total, nil
}

func (s *ProductUseCase) UpdateProduct(
	ctx context.Context,
	id uuid.UUID,
	payload models.CreateProductRequest,
	authInfo models.AuthenticationInfo,
) (*models.Product, error) {
	product, err := s.productRepo.GetProductByID(ctx, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, response.NewNotFoundError("Product not found.")
		}

		return nil, response.NewInternalServerError(err)
	}

	productCategoryId, err := uuid.Parse(payload.CategoryID)
	if err != nil {
		return nil, response.NewBadRequestError("Invalid product category ID.")
	}

	_, err = s.productCategoryRepo.GetProductCategoryByID(ctx, productCategoryId)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, response.NewBadRequestError("Product category not found.")

		}

		return nil, response.NewInternalServerError(err)
	}

	if payload.Name != "" && product.Name != payload.Name {
		product.Name = payload.Name
	}

	if payload.Description != nil && product.Description != payload.Description {
		product.Description = payload.Description
	}

	if payload.Brand != nil && product.Brand != payload.Brand {
		product.Brand = payload.Brand
	}

	if payload.Origin != nil && product.Origin != payload.Origin {
		product.Origin = payload.Origin
	}

	if payload.UserGuide != nil && product.UserGuide != payload.UserGuide {
		product.UserGuide = payload.UserGuide
	}

	if product.CategoryID != productCategoryId {
		product.CategoryID = productCategoryId
	}

	newColorNotInOld := make([]*models.CreateProductVariantInProductRequest, 0)
	oldColorNotInNew := make([]*models.ProductVariant, 0)
	oldColorMap := make(map[string]bool)
	for _, oldVariant := range product.ProductVariants {
		oldColorMap[oldVariant.Color] = true
	}

	newColorMap := make(map[string]bool)
	seenColors := make(map[string]bool)
	for _, variant := range payload.ProductVariants {
		if seenColors[variant.Color] {
			return nil, response.NewBadRequestError(
				"Duplicate color found in new product variants.",
			)
		}

		seenColors[variant.Color] = true
		newColorMap[variant.Color] = true
		if !oldColorMap[variant.Color] {
			newColorNotInOld = append(newColorNotInOld, variant)
		}
	}

	for _, oldVariant := range product.ProductVariants {
		if !newColorMap[oldVariant.Color] {
			oldColorNotInNew = append(oldColorNotInNew, oldVariant)
		}
	}

	_, err = s.productRepo.UpdateProduct(ctx, product)
	if err != nil {
		return nil, response.NewInternalServerError(err)
	}

	for _, oldVariant := range oldColorNotInNew {
		// check if the old variant is in a transaction
		err = s.productVariantRepo.DeleteProductVariant(ctx, oldVariant)
		if err != nil {
			return nil, response.NewInternalServerError(err)
		}
	}

	for _, newVariant := range newColorNotInOld {
		productVariant := models.NewProductVariant(
			product.ID,
			newVariant.Description,
			newVariant.Color,
			newVariant.RetailPrice,
			newVariant.Stock,
			&authInfo.User.ID,
		)
		_, err = s.productVariantRepo.CreateProductVariant(ctx, productVariant)
		if err != nil {
			return nil, response.NewInternalServerError(err)
		}
	}

	return product, nil
}

func (s *ProductUseCase) DeleteProduct(
	ctx context.Context,
	id uuid.UUID,
	authInfo models.AuthenticationInfo,
) error {
	product, err := s.productRepo.GetProductByID(ctx, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return response.NewNotFoundError("Product not found.")
		}

		return response.NewInternalServerError(err)
	}

	// Check if the product is in a transaction
	err = s.productRepo.DeleteProduct(ctx, product)
	if err != nil {
		return response.NewInternalServerError(err)
	}

	return nil
}

func (s *ProductUseCase) ApproveProduct(
	ctx context.Context,
	id uuid.UUID,
	authInfo models.AuthenticationInfo,
) error {
	product, err := s.productRepo.GetProductByID(ctx, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return response.NewNotFoundError("Product not found.")
		}

		return response.NewInternalServerError(err)
	}

	product.Status = models.ActiveStatus
	product.UpdatedAt = time.Now().UTC()
	product.UpdatedBy = &authInfo.User.ID
	_, err = s.productRepo.UpdateProduct(ctx, product)
	if err != nil {
		return response.NewInternalServerError(err)
	}

	return nil
}

func (s *ProductUseCase) RejectProduct(
	ctx context.Context,
	id uuid.UUID,
	authInfo models.AuthenticationInfo,
) error {
	product, err := s.productRepo.GetProductByID(ctx, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return response.NewNotFoundError("Product not found.")
		}

		return response.NewInternalServerError(err)
	}

	product.Status = models.RejectedStatus
	product.UpdatedAt = time.Now().UTC()
	product.UpdatedBy = &authInfo.User.ID
	_, err = s.productRepo.UpdateProduct(ctx, product)
	if err != nil {
		return response.NewInternalServerError(err)
	}

	return nil
}
