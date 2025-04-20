package middlewares

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-redis/redis_rate/v9"
	"github.com/labstack/echo/v4"
)

type RateLimiter interface {
	Allow(ctx context.Context, key string, limit redis_rate.Limit) (*redis_rate.Result, error)
}

type RateLimiterMiddleware struct {
	RedisStore RateLimiter
}

func NewRateLimiter(redisStore RateLimiter) *RateLimiterMiddleware {
	return &RateLimiterMiddleware{
		RedisStore: redisStore,
	}
}

func (r *RateLimiterMiddleware) LimitRequests(limit redis_rate.Limit) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.Request().Context()
			clientIP := c.RealIP()
			key := fmt.Sprintf("rate_limit:%s", clientIP)

			res, err := r.RedisStore.Allow(ctx, key, limit)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": "rate limiter error"})
			}

			if res.Allowed == 0 {
				return c.JSON(http.StatusTooManyRequests, map[string]string{
					"error":       "rate limit exceeded",
					"retry_after": res.RetryAfter.String(),
				})
			}

			return next(c)
		}
	}
}
