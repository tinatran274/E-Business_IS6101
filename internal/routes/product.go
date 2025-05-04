package routes

import (
	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/handlers"
	"github.com/labstack/echo/v4"
)

type ProductRouter struct {
	productHandler *handlers.ProductHandler
	authMiddleware echo.MiddlewareFunc
}

func NewProductRouter(
	productHandler *handlers.ProductHandler,
	authMiddleware echo.MiddlewareFunc,
) *ProductRouter {
	return &ProductRouter{
		productHandler: productHandler,
		authMiddleware: authMiddleware,
	}
}

func (r *ProductRouter) Register(e *echo.Group) {
	productRoute := e.Group("/product")
	if r.authMiddleware != nil {
		productRoute.Use(r.authMiddleware)
	}

	productRoute.POST("", r.productHandler.CreateProduct)
	productRoute.GET("", r.productHandler.GetProducts)
	productRoute.GET("/me", r.productHandler.GetMyProducts)
	productRoute.GET("/:id", r.productHandler.GetProductByID)
	productRoute.PUT("/:id", r.productHandler.UpdateProduct)
	productRoute.DELETE("/:id", r.productHandler.DeleteProduct)
	productRoute.POST("/:id/approve", r.productHandler.ApproveProduct)
	productRoute.POST("/:id/reject", r.productHandler.RejectProduct)
}
