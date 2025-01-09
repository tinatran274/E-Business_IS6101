package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	user_actions "10.0.0.50/tuan.quang.tran/aioz-ads/internal/actions/user"
	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/config"
	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/handlers"
	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/models"
	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/routes"
	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/services"
	custom_log "10.0.0.50/tuan.quang.tran/aioz-ads/internal/utils/log"
	"10.0.0.50/tuan.quang.tran/aioz-ads/pkg/v1/db"
	"10.0.0.50/tuan.quang.tran/aioz-ads/pkg/v1/repositories"
	v1 "10.0.0.50/tuan.quang.tran/aioz-ads/pkg/v1/services"
)

var (
	appConfig    *config.AppConfig
	pool         *pgxpool.Pool
	defaultQuery *db.Queries

	userRepo models.UserRepository

	userService services.UserService

	userActions *user_actions.UserActions

	authHandler *handlers.AuthHandler
	userHandler *handlers.UserHandler

	authRouter *routes.AuthRouter
	userRouter *routes.UserRouter

	server *echo.Echo
)

func init() {
	// set timezone
	if os.Setenv("TZ", "UTC") != nil {
		panic("failed to set timezone")
	}

	// load app config
	env := os.Getenv("ENV")
	if env == "" {
		env = "app"
	}

	appConfig = config.MustNewAppConfig(fmt.Sprintf("%s.env", env))

	// logger config
	defaultLogger := slog.New(
		custom_log.NewHandler(&slog.HandlerOptions{}, nil),
	)
	slog.SetDefault(defaultLogger)
	debugLogger := slog.New(
		custom_log.NewHandler(&slog.HandlerOptions{
			Level: slog.LevelDebug,
		}, nil),
	)

	logTracer := custom_log.NewLogTracer(
		custom_log.NewLogger(debugLogger),
		tracelog.LogLevelWarn,
		true,
		1*time.Second,
	)

	// init database
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbConfig, err := pgxpool.ParseConfig("some-connection-string")
	if err != nil {
		// panic("parse db config error")
		slog.Warn("should panic here")
	}

	dbConfig.ConnConfig.Tracer = logTracer
	pool, err = pgxpool.NewWithConfig(
		ctx,
		dbConfig,
	)
	if err != nil {
		// panic("failed to connect to database")
		slog.Warn("should panic here")
	}

	defaultQuery = db.New(pool)

	// init repositories
	userRepo = repositories.NewUserRepository(defaultQuery)

	// init services
	userService = v1.NewUserService(userRepo)

	// init actions
	userActions = user_actions.NewUserActions(userService)

	// init handlers
	authHandler = handlers.NewAuthHandler(userActions)
	userHandler = handlers.NewUserHanlder(userActions)

	// init routers
	authRouter = routes.NewAuthRouter(authHandler)
	userRouter = routes.NewUserRouter(userHandler, nil)

	server = echo.New()
	configLogger(server)
}

func configLogger(server *echo.Echo) {
	ExcludedPaths := map[string]bool{
		"/metrics": true,
	}

	server.Use(
		middleware.RequestLoggerWithConfig(
			middleware.RequestLoggerConfig{
				LogStatus:   true,
				LogURI:      true,
				LogError:    true,
				LogMethod:   true,
				LogLatency:  true,
				HandleError: true,
				LogValuesFunc: func(
					c echo.Context,
					v middleware.RequestLoggerValues,
				) error {
					if ExcludedPaths[c.Path()] {
						return nil
					}

					d := v.Latency
					var latencyString string
					switch {
					case d >= time.Second:
						latencyString = fmt.Sprintf(
							"%03ds",
							int64(d.Seconds()),
						)
					case d >= time.Millisecond:
						latencyString = fmt.Sprintf(
							"%03dms",
							int64(d.Milliseconds()),
						)
					case d >= time.Microsecond:
						latencyString = fmt.Sprintf(
							"%03dµs",
							int64(d.Microseconds()),
						)
					default:
						latencyString = fmt.Sprintf(
							"%03dns",
							d.Nanoseconds(),
						)
					}

					var method string
					switch v.Method {
					case http.MethodGet:
						method = custom_log.GetColor
					case http.MethodPost:
						method = custom_log.PostColor
					case http.MethodPut:
						method = custom_log.PutColor
					case http.MethodPatch:
						method = custom_log.PatchColor
					case http.MethodDelete:
						method = custom_log.DeleteColor
					default:
						return nil
					}

					var status string
					switch {
					case v.Status >= 200 && v.Status < 400:
						status = fmt.Sprintf(custom_log.SuccessColor, v.Status)
					case v.Status >= 400 && v.Status < 500:
						status = fmt.Sprintf(custom_log.FailColor, v.Status)
					default:
						status = fmt.Sprintf(custom_log.ErrorColor, v.Status)
					}

					slog.InfoContext(
						c.Request().Context(),
						fmt.Sprintf(
							"%s status: %s latency: %s %s",
							method,
							status,
							latencyString,
							v.URI,
						),
					)

					return nil
				},
			},
		),
	)
}
