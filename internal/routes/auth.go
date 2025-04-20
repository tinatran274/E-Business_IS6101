package routes

import (
	"github.com/labstack/echo/v4"

	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/handlers"
)

type AuthRouter struct {
	authHandler *handlers.AuthHandler
}

func NewAuthRouter(
	authHandler *handlers.AuthHandler,
) *AuthRouter {
	return &AuthRouter{
		authHandler: authHandler,
	}
}

func (r *AuthRouter) Register(e *echo.Group) {
	authRoute := e.Group("/auth")

	authRoute.POST("/sign-up", r.authHandler.SignUp)
	authRoute.POST("/sign-in", r.authHandler.SignIn)
}
