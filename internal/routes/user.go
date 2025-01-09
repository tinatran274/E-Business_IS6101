package routes

import (
	"github.com/labstack/echo/v4"

	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/handlers"
)

type UserRouter struct {
	userHandler *handlers.UserHandler

	authMiddleware echo.MiddlewareFunc
}

func NewUserRouter(
	userHandler *handlers.UserHandler,

	authMiddleware echo.MiddlewareFunc,
) *UserRouter {
	return &UserRouter{
		userHandler: userHandler,

		authMiddleware: authMiddleware,
	}
}

func (r *UserRouter) Register(e *echo.Group) {
	userRoute := e.Group("/user")
	userRoute.Use(r.authMiddleware)

	e.GET("/me", r.userHandler.GetMe)
}
