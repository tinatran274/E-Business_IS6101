package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	// cors config
	server.Use(
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

	// check health
	server.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	// metrics
	server.GET("/metrics", func(c echo.Context) error {
		promhttp.Handler().ServeHTTP(c.Response(), c.Request())
		return nil
	})

	// swagger docs
	server.GET(
		"/swagger/*",
		echoSwagger.WrapHandler,
	)

	// router register
	api := server.Group("/api")
	api.Use(middlewares.AddHttpLogContext())

	// v1 router
	v1 := api.Group("/v1")
	authRouter.Register(v1)
	userRouter.Register(v1)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()

	go func() {
		if err := server.Start(fmt.Sprintf(":%s", appConfig.Port)); err != nil {
			if err != http.ErrServerClosed {
				slog.Error(
					"can not start server",
					slog.Any("port", appConfig.Port),
					slog.Any("error", err),
				)
				os.Exit(1)
			}
		}
	}()

	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("shutdown server error", slog.Any("error", err))
		os.Exit(1)
	}

	pool.Close()

	slog.Info("server gracefully shutdown")
}

//	@title			AIOZ ADS SERVICE
//	@version		1.0
//	@description	AIOZ ADS SERVICE API DOCUMENTATION
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@BasePath	/api/v1

//	@externalDocs.description	Find out more about Swagger
//	@externalDocs.url			https://swagger.io/resources/open-api/

//	@securityDefinitions.basic	BasicAuth

//	@securityDefinitions.apikey	Bearer
//	@in							header
//	@name						Authorization
//	@description				Type "Bearer" followed by a space and JWT token.
