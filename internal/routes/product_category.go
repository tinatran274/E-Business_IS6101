package routes

import (
	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/handlers"
	"github.com/labstack/echo/v4"
)

type ProductCategoryRouter struct {
	productCategoryHandler *handlers.ProductCategoryHandler
	authMiddleware         echo.MiddlewareFunc
}

func NewProductCategoryRouter(
	productCategoryHandler *handlers.ProductCategoryHandler,
	authMiddleware echo.MiddlewareFunc,
) *ProductCategoryRouter {
	return &ProductCategoryRouter{
		productCategoryHandler: productCategoryHandler,
		authMiddleware:         authMiddleware,
	}
}

func (r *ProductCategoryRouter) Register(e *echo.Group) {
	productCategoryRoute := e.Group("/product-category")
	if r.authMiddleware != nil {
		productCategoryRoute.Use(r.authMiddleware)
	}

	// productCategoryRoute.POST("", r.productCategoryHandler.CreateProductCategory)
	// productCategoryRoute.GET("", r.productCategoryHandler.GetProductCategories)
	// productCategoryRoute.GET("/:id", r.productCategoryHandler.GetProductCategoryByID)
	// productCategoryRoute.PUT("/:id", r.productCategoryHandler.UpdateProductCategory)
	// productCategoryRoute.DELETE("/:id", r.productCategoryHandler.DeleteProductCategory)
}
