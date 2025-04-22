package routes

import (
	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/handlers"
	"github.com/labstack/echo/v4"
)

type IngredientRouter struct {
	ingredientHandler *handlers.IngredientHandler
	authMiddleware    echo.MiddlewareFunc
}

func NewIngredientRouter(
	ingredientHandler *handlers.IngredientHandler,
	authMiddleware echo.MiddlewareFunc,
) *IngredientRouter {
	return &IngredientRouter{
		ingredientHandler: ingredientHandler,
		authMiddleware:    authMiddleware,
	}
}

func (r *IngredientRouter) Register(e *echo.Group) {
	ingredientRoute := e.Group("/ingredient")
	if r.authMiddleware != nil {
		ingredientRoute.Use(r.authMiddleware)
	}

	ingredientRoute.GET("", r.ingredientHandler.GetIngredients)
	ingredientRoute.GET("/:id", r.ingredientHandler.GetIngredientByID)
}
