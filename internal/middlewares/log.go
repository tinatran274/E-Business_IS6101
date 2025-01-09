package middlewares

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	custom_log "10.0.0.50/tuan.quang.tran/aioz-ads/internal/utils/log"
)

func AddHttpLogContext() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			requestId := uuid.New().String()
			data := []slog.Attr{
				slog.Any("request_id", requestId),
			}

			ctx.SetRequest(ctx.Request().WithContext(context.WithValue(
				ctx.Request().Context(),
				custom_log.SlogFieldsKey,
				data,
			)))

			return next(ctx)
		}
	}
}
