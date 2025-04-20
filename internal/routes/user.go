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
	if r.authMiddleware != nil {
		userRoute.Use(r.authMiddleware)
	}

	userRoute.GET("/me", r.userHandler.GetMe)
	userRoute.GET("", r.userHandler.GetAll)
	userRoute.GET("/:id", r.userHandler.GetByID)
}
