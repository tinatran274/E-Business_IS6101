package handlers

import "10.0.0.50/tuan.quang.tran/aioz-ads/internal/usecases"

type ProductCategoryHandler struct {
	productCategoryUseCase usecases.ProductCategoryUseCase
}

func NewProductCategoryHandler(
	productCategoryUseCase usecases.ProductCategoryUseCase,
) *ProductCategoryHandler {
	return &ProductCategoryHandler{
		productCategoryUseCase: productCategoryUseCase,
	}
}
