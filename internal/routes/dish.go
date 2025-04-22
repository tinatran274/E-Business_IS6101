package routes

import (
	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/handlers"
	"github.com/labstack/echo/v4"
)

type DishRouter struct {
	dishHandler    *handlers.DishHandler
	authMiddleware echo.MiddlewareFunc
}

func NewDishRouter(
	dishHandler *handlers.DishHandler,
	authMiddleware echo.MiddlewareFunc,
) *DishRouter {
	return &DishRouter{
		dishHandler:    dishHandler,
		authMiddleware: authMiddleware,
	}
}

func (r *DishRouter) Register(e *echo.Group) {
	dishRoute := e.Group("/dish")
	if r.authMiddleware != nil {
		dishRoute.Use(r.authMiddleware)
	}

	dishRoute.GET("", r.dishHandler.GetDishs)
	dishRoute.GET("/:id", r.dishHandler.GetDishByID)
}
