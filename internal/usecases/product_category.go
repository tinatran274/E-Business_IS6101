package usecases

import (
	"context"

	"10.0.0.50/tuan.quang.tran/aioz-ads/models"
)

type ProductCategoryUseCase interface {
	GetProductCategories(
		ctx context.Context,
	) ([]*models.ProductCategory, error)
}
