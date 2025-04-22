package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"10.0.0.50/tuan.quang.tran/aioz-ads/config"
	db "10.0.0.50/tuan.quang.tran/aioz-ads/db/generated"
	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/handlers"
	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/middlewares"
	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/routes"
	custom_log "10.0.0.50/tuan.quang.tran/aioz-ads/internal/utils/log"
	"10.0.0.50/tuan.quang.tran/aioz-ads/pkg/v1/postgres"
	"10.0.0.50/tuan.quang.tran/aioz-ads/pkg/v1/redis"
	"10.0.0.50/tuan.quang.tran/aioz-ads/pkg/v1/repositories"
	"10.0.0.50/tuan.quang.tran/aioz-ads/pkg/v1/usecases"

	"github.com/go-redis/redis_rate/v9"
	"github.com/jackc/pgx/v5/tracelog"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	config *config.AppConfig
	pg     *postgres.Postgres
	server *echo.Echo
}

func New(appConfig *config.AppConfig) *Server {
	return &Server{
		config: appConfig,
		server: echo.New(),
	}
}

func (s *Server) Start() {
	s.initDatabase()
	if s.pg == nil {
		slog.Error("Database connection is nil, unable to proceed with server start.")
		os.Exit(1)
	}
	defer s.pg.Pool.Close()

	redisStore := redis.NewRedisClient(s.config.Addr, s.config.Password, s.config.DB)
	defer redisStore.Close()

	rateLimiter := middlewares.NewRateLimiter(redisStore)
	s.server.Use(rateLimiter.LimitRequests(redis_rate.PerMinute(50)))

	q := db.New(s.pg.Pool)
	userRepo := repositories.NewUserRepository(q)
	accountRepo := repositories.NewAccountRepository(q)
	ingredientRepo := repositories.NewIngredientRepository(q)
	dishRepo := repositories.NewDishRepository(q)

	userUseCase := usecases.NewUserUseCase(userRepo, accountRepo)
	authUseCase := usecases.NewAuthUseCase(accountRepo, userRepo)
	ingredientUseCase := usecases.NewIngredientUseCase(ingredientRepo)
	dishUseCase := usecases.NewDishUseCase(dishRepo)

	authHandler := handlers.NewAuthHandler(userUseCase, authUseCase)
	userHandler := handlers.NewUserHanlder(userUseCase)
	ingredientHandler := handlers.NewIngredientHandler(ingredientUseCase)
	dishHandler := handlers.NewDishHandler(dishUseCase)

	authMiddleware := middlewares.JWTAuthMiddleware(
		[]byte(s.config.JwtSecret),
		userUseCase,
	)
	authRouter := routes.NewAuthRouter(authHandler)
	userRouter := routes.NewUserRouter(userHandler, authMiddleware)
	ingredientRouter := routes.NewIngredientRouter(
		ingredientHandler,
		authMiddleware,
	)
	dishRouter := routes.NewDishRouter(
		dishHandler,
		authMiddleware,
	)

	routes.NewRouter(
		s.server,
		authRouter,
		userRouter,
		ingredientRouter,
		dishRouter,
	)

	s.configLogger()
	addr := fmt.Sprintf(":%s", s.config.Port)
	slog.Info("Server is starting", "addrress", addr)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()

	go func() {
		if err := s.server.Start(fmt.Sprintf(":%s", s.config.Port)); err != nil {
			if err != http.ErrServerClosed {
				slog.Error(
					"can not start server",
					slog.Any("port", s.config.Port),
					slog.Any("error", err),
				)
				os.Exit(1)
			}
		}
	}()

	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		slog.Error("shutdown server error", slog.Any("error", err))
		os.Exit(1)
	}

	s.pg.Pool.Close()

	slog.Info("server gracefully shutdown")
}

func (s *Server) initDatabase() error {
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

	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		s.config.PostgresUser,
		s.config.PostgresPassword,
		s.config.PostgresHost,
		s.config.PostgresPort,
		s.config.PostgresDBName,
	)

	pg, err := postgres.New(connString, postgres.WithTracer(logTracer))
	if err != nil {
		slog.Error("Failed to initialize database", "error", err)
		return fmt.Errorf("failed to initialize database: %w", err)
	}

	s.pg = pg
	return nil
}

func (s *Server) configLogger() {
	ExcludedPaths := map[string]bool{
		"/metrics": true,
	}

	s.server.Use(
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
