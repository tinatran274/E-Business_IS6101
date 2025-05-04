package usecases

import (
	"context"

	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/utils/response"
	"10.0.0.50/tuan.quang.tran/aioz-ads/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type ProductVariantUseCase struct {
	productVariantRepo models.ProductVariantRepository
	productRepo        models.ProductRepository
}

func NewProductVariantUseCase(
	productVariantRepo models.ProductVariantRepository,
	productRepo models.ProductRepository,
) *ProductVariantUseCase {
	return &ProductVariantUseCase{
		productVariantRepo: productVariantRepo,
		productRepo:        productRepo,
	}
}

func (s *ProductVariantUseCase) CreateProductVariant(
	ctx context.Context,
	payload models.CreateProductVariantRequest,
	authInfo models.AuthenticationInfo,
) (*models.ProductVariant, error) {
	productID, err := uuid.Parse(payload.ProductID)
	if err != nil {
		return nil, response.NewBadRequestError("Invalid product ID.")
	}

	_, err = s.productRepo.GetProductByID(ctx, productID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, response.NewBadRequestError("Product not found.")
		}

		return nil, response.NewInternalServerError(err)
	}

	isExist, err := s.productVariantRepo.IsProductVariantExist(
		ctx,
		productID,
		payload.Color,
	)
	if err != nil {
		return nil, response.NewInternalServerError(err)
	}

	if isExist {
		return nil, response.NewBadRequestError("Product color already exists.")
	}

	productVariant := models.NewProductVariant(
		productID,
		payload.Description,
		payload.Color,
		payload.RetailPrice,
		payload.Stock,
		&authInfo.User.ID,
	)
	_, err = s.productVariantRepo.CreateProductVariant(ctx, productVariant)
	if err != nil {
		return nil, response.NewInternalServerError(err)
	}

	return productVariant, nil
}

func (s *ProductVariantUseCase) GetProductVariantsByProductID(
	ctx context.Context,
	productID uuid.UUID,
) ([]*models.ProductVariant, error) {
	_, err := s.productRepo.GetProductByID(ctx, productID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, response.NewBadRequestError("Product not found.")
		}

		return nil, response.NewInternalServerError(err)
	}

	productVariants, err := s.productVariantRepo.GetProductVariantsByProductID(ctx, productID)
	if err != nil {
		return nil, response.NewInternalServerError(err)
	}

	return productVariants, nil
}

func (s *ProductVariantUseCase) UpdateProductVariant(
	ctx context.Context,
	productVariantID uuid.UUID,
	payload models.CreateProductVariantRequest,
	authInfo models.AuthenticationInfo,
) (*models.ProductVariant, error) {
	productVariant, err := s.productVariantRepo.GetProductVariantByID(ctx, productVariantID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, response.NewBadRequestError("Product variant not found.")
		}

		return nil, response.NewInternalServerError(err)
	}

	productID, err := uuid.Parse(payload.ProductID)
	if err != nil {
		return nil, response.NewBadRequestError("Invalid product ID.")
	}

	_, err = s.productRepo.GetProductByID(ctx, productID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, response.NewBadRequestError("Product not found.")
		}

		return nil, response.NewInternalServerError(err)
	}

	if productVariant.ProductID != productID {
		productVariant.ProductID = productID
	}

	if payload.Description != nil && payload.Description != productVariant.Description {
		productVariant.Description = payload.Description
	}

	if payload.Color != "" && payload.Color != productVariant.Color {
		isExist, err := s.productVariantRepo.IsProductVariantExist(
			ctx,
			productID,
			payload.Color,
		)
		if err != nil {
			return nil, response.NewInternalServerError(err)
		}

		if isExist {
			return nil, response.NewBadRequestError("Product color already exists.")
		}

		productVariant.Color = payload.Color
	}

	if payload.RetailPrice != productVariant.RetailPrice {
		productVariant.RetailPrice = payload.RetailPrice
	}

	if payload.Stock != productVariant.Stock {
		productVariant.Stock = payload.Stock
	}

	_, err = s.productVariantRepo.UpdateProductVariant(ctx, productVariant)
	if err != nil {
		return nil, response.NewInternalServerError(err)
	}

	return productVariant, nil
}

func (s *ProductVariantUseCase) DeleteProductVariant(
	ctx context.Context,
	productVariantID uuid.UUID,
	authInfo models.AuthenticationInfo,
) error {
	productVariant, err := s.productVariantRepo.GetProductVariantByID(ctx, productVariantID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return response.NewBadRequestError("Product variant not found.")
		}

		return response.NewInternalServerError(err)
	}

	err = s.productVariantRepo.DeleteProductVariant(ctx, productVariant)
	if err != nil {
		return response.NewInternalServerError(err)
	}

	return nil
}
