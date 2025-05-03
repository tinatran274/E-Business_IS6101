package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/usecases"
	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/utils/response"
	"10.0.0.50/tuan.quang.tran/aioz-ads/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

func JWTAuthMiddleware(
	jwtSecret []byte,
	userUseCase usecases.UserUseCase,
) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			authHeader := ctx.Request().Header.Get("Authorization")
			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				return response.ResponseFailMessage(ctx, http.StatusUnauthorized, "Missing or invalid token.")
			}

			tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

			token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method")
				}
				return jwtSecret, nil
			})

			if err != nil || !token.Valid {
				return response.ResponseFailMessage(ctx, http.StatusUnauthorized, "Invalid token.")
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok || claims["sub"] == nil {
				return response.ResponseFailMessage(ctx, http.StatusUnauthorized, "Invalid token claims.")
			}

			userID := claims["sub"].(string)
			userUUID, err := uuid.Parse(userID)
			if err != nil {
				return response.ResponseFailMessage(ctx, http.StatusUnauthorized, "Invalid user ID.")
			}
			user, err := userUseCase.GetUserByAccountId(ctx.Request().Context(), userUUID)
			if err != nil {
				if err == pgx.ErrNoRows {
					return response.ResponseFailMessage(ctx, http.StatusUnauthorized, "User not found.")
				}

				return response.NewInternalServerError(err)
			}

			authInfo := models.AuthenticationInfo{User: user}
			ctx.Set(models.AuthInfoKey, authInfo)

			return next(ctx)
		}
	}
}
