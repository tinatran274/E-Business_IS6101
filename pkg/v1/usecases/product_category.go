package usecases

import (
	"context"

	"10.0.0.50/tuan.quang.tran/aioz-ads/models"
)

type ProductCategoryUseCase struct {
	productCategoryRepo models.ProductCategoryRepository
}

func NewProductCategoryUseCase(
	productCategoryRepo models.ProductCategoryRepository,
) *ProductCategoryUseCase {
	return &ProductCategoryUseCase{
		productCategoryRepo: productCategoryRepo,
	}
}

func (u *ProductCategoryUseCase) GetProductCategories(
	ctx context.Context,
) ([]*models.ProductCategory, error) {
	productCategories, err := u.productCategoryRepo.GetProductCategories(ctx)
	if err != nil {
		return nil, err
	}

	return productCategories, nil
}
