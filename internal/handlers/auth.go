package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	user_actions "10.0.0.50/tuan.quang.tran/aioz-ads/internal/actions/user"
	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/models"
	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/utils/metrics"
	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/utils/response"
)

type AuthHandler struct {
	userActions *user_actions.UserActions
}

func NewAuthHandler(userActions *user_actions.UserActions) *AuthHandler {
	return &AuthHandler{
		userActions: userActions,
	}
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type CreateUserResponse struct {
	User *models.User `json:"user"`
}

// SignUp godoc
//
//	@Summary		Sign up
//	@Description	sign up new user
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Accept			x-www-form-urlencoded
//	@Param			payload	body		CreateUserRequest	true	"create org input"
//	@Success		201		{object}	models.User
//	@Failure		400		{object}	response.GeneralResponse
//	@Failure		500		{object}	response.GeneralResponse
//	@Router			/auth/sign-up [post]
func (h *AuthHandler) SignUp(ctx echo.Context) error {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.ApiSum.WithLabelValues("SignUp").
			Observe(time.Since(t).Seconds())
	}()

	panic("implement me")

	return response.ResponseSuccess(ctx, http.StatusOK, CreateUserResponse{
		User: &models.User{},
	})
}
