package routes

import (
	"net/http"

	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func NewRouter(
	handler *echo.Echo,
	authRouter *AuthRouter,
	userRouter *UserRouter,
	ingredientRouter *IngredientRouter,
	dishRouter *DishRouter,
	statisticRouter *StatisticRouter,
) {
	handler.Use(
		middleware.CORSWithConfig(
			middleware.CORSConfig{
				AllowOrigins: []string{"*"},
				AllowHeaders: []string{
					echo.HeaderOrigin,
					echo.HeaderContentType,
					echo.HeaderAccept,
					echo.HeaderAuthorization,
				},
				AllowMethods: []string{
					http.MethodGet,
					http.MethodPost,
					http.MethodPut,
					http.MethodPatch,
					http.MethodDelete,
					http.MethodOptions,
				},
			},
		),
	)

	handler.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	handler.GET("/metrics", func(c echo.Context) error {
		promhttp.Handler().ServeHTTP(c.Response(), c.Request())
		return nil
	})

	handler.GET(
		"/swagger/*",
		echoSwagger.WrapHandler,
	)

	h := handler.Group("/api")
	h.Use(middlewares.AddHttpLogContext())

	v1 := h.Group("/v1")
	userRouter.Register(v1)
	authRouter.Register(v1)
	ingredientRouter.Register(v1)
	dishRouter.Register(v1)
	statisticRouter.Register(v1)
}
