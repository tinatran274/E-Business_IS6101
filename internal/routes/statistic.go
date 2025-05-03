package routes

import (
	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/handlers"
	"github.com/labstack/echo/v4"
)

type StatisticRouter struct {
	statisticHandler *handlers.StatisticHandler
	authMiddleware   echo.MiddlewareFunc
}

func NewStatisticRouter(
	statisticHandler *handlers.StatisticHandler,
	authMiddleware echo.MiddlewareFunc,
) *StatisticRouter {
	return &StatisticRouter{
		statisticHandler: statisticHandler,
		authMiddleware:   authMiddleware,
	}
}

func (r *StatisticRouter) Register(e *echo.Group) {
	statisticRoute := e.Group("/statistic")
	if r.authMiddleware != nil {
		statisticRoute.Use(r.authMiddleware)
	}

	statisticRoute.GET("/:user_id", r.statisticHandler.GetStatisticByUserIdAndDate)
	statisticRoute.GET("/:user_id/range", r.statisticHandler.GetStatisticByUserIdAndDateRange)
	statisticRoute.PUT("/:user_id", r.statisticHandler.UpdateStatisticByUserIdAndDate)
}
