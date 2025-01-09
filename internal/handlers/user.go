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

type UserHandler struct {
	userActions *user_actions.UserActions
}

func NewUserHanlder(userActions *user_actions.UserActions) *UserHandler {
	return &UserHandler{
		userActions: userActions,
	}
}

// GetMe godoc
//
//	@Summary		Get me
//	@Description	get current user
//	@Tags			users
//	@Security		BasicAuth
//	@Security		Bearer
//	@Accept			json
//	@Produce		json
//	@Accept			x-www-form-urlencoded
//	@Success		200	{object}	models.User
//	@Failure		400	{object}	response.GeneralResponse
//	@Failure		500	{object}	response.GeneralResponse
//	@Router			/user/me [get]
func (h *UserHandler) GetMe(ctx echo.Context) error {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.ApiSum.WithLabelValues("GetMe").
			Observe(time.Since(t).Seconds())
	}()

	authInfo, ok := ctx.Get(models.AuthInfoKey).(models.AuthenticationInfo)
	if !ok {
		return response.ResponseFailMessage(ctx, http.StatusUnauthorized, "unauthorized")
	}

	user, err := h.userActions.GetMeAction.Exec(ctx.Request().Context(), authInfo.User.Id)
	if err != nil {
		return response.ResponseError(ctx, err)
	}

	return response.ResponseSuccess(ctx, http.StatusOK, *user)
}
