package routes

import (
	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/handlers"
	"github.com/labstack/echo/v4"
)

type ProductVariantRouter struct {
	productVariantHandler *handlers.ProductVariantHandler
	authMiddleware        echo.MiddlewareFunc
}

func NewProductVariantRouter(
	productVariantHandler *handlers.ProductVariantHandler,
	authMiddleware echo.MiddlewareFunc,
) *ProductVariantRouter {
	return &ProductVariantRouter{
		productVariantHandler: productVariantHandler,
		authMiddleware:        authMiddleware,
	}
}

func (r *ProductVariantRouter) Register(e *echo.Group) {
	productVariantRoute := e.Group("/product-variant")
	if r.authMiddleware != nil {
		productVariantRoute.Use(r.authMiddleware)
	}

	productVariantRoute.POST("", r.productVariantHandler.CreateProductVariant)
	productVariantRoute.GET("/product/:id", r.productVariantHandler.GetProductVariantsByProductID)
	productVariantRoute.PUT("/:id", r.productVariantHandler.UpdateProductVariant)
	productVariantRoute.DELETE("/:id", r.productVariantHandler.DeleteProductVariant)
	
}
