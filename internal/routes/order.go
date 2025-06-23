package routes

import (
	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/handlers"
	"github.com/labstack/echo/v4"
)

type OrderRouter struct {
	OrderHandler   *handlers.OrderHandler
	authMiddleware echo.MiddlewareFunc
}

func NewOrderRouter(
	OrderHandler *handlers.OrderHandler,
	authMiddleware echo.MiddlewareFunc,
) *OrderRouter {
	return &OrderRouter{
		OrderHandler:   OrderHandler,
		authMiddleware: authMiddleware,
	}
}

func (r *OrderRouter) Register(e *echo.Group) {
	orderRoute := e.Group("/orders")
	if r.authMiddleware != nil {
		orderRoute.Use(r.authMiddleware)
	}

	// orderRoute.POST("", r.OrderHandler.CreateOrder)
	// orderRoute.GET("", r.OrderHandler.GetOrdersByUserID)
	// orderRoute.GET("/:id", r.OrderHandler.GetOrderByID)
	// orderRoute.PUT("/:id", r.OrderHandler.UpdateOrderStatus)
	// orderRoute.DELETE("/:id", r.OrderHandler.DeleteOrder)
}
