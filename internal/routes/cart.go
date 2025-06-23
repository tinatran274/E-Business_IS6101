package routes

import (
	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/handlers"
	"github.com/labstack/echo/v4"
)

type CartRouter struct {
	CartHandler    *handlers.CartHandler
	authMiddleware echo.MiddlewareFunc
}

func NewCartRouter(
	CartHandler *handlers.CartHandler,
	authMiddleware echo.MiddlewareFunc,
) *CartRouter {
	return &CartRouter{
		CartHandler:    CartHandler,
		authMiddleware: authMiddleware,
	}
}

func (r *CartRouter) Register(e *echo.Group) {
	cartRoute := e.Group("/cart")
	if r.authMiddleware != nil {
		cartRoute.Use(r.authMiddleware)
	}

	cartRoute.POST("", r.CartHandler.AddCartItem)
	cartRoute.GET("", r.CartHandler.GetCartItemsByUserID)
	cartRoute.DELETE("/:id", r.CartHandler.DeleteCartItem)
	cartRoute.PUT("/:id", r.CartHandler.UpdateCartItem)
}
