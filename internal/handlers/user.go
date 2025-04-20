package handlers

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/usecases"
	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/utils/metrics"
	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/utils/response"
	"10.0.0.50/tuan.quang.tran/aioz-ads/models"
)

type UserHandler struct {
	userUseCase usecases.UserUseCase
}

func NewUserHanlder(userUseCase usecases.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

// @Summary		Get me
// @Description	get current user
// @Tags			users
// @Security		BasicAuth
// @Security		Bearer
// @Accept			json
// @Produce		json
// @Accept			x-www-form-urlencoded
// @Success		200	{object}	response.GeneralResponse
// @Failure		400	{object}	response.GeneralResponse
// @Failure		500	{object}	response.GeneralResponse
// @Router			/user/me [get]
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

	user, err := h.userUseCase.GetMe(ctx.Request().Context(), authInfo.User.ID)
	if err != nil {
		return response.ResponseError(ctx, err)
	}

	return response.ResponseSuccess(ctx, http.StatusOK, *user)
}

type UserResponse struct {
	UserList []models.User `json:"user"`
}

// @Summary		Get all
// @Description	get all user
// @Tags			users
// @Accept			json
// @Produce		json
// @Accept			x-www-form-urlencoded
// @Success		200	{object}	response.GeneralResponse
// @Failure		500	{object}	response.GeneralResponse
// @Failure		500	{object}	response.GeneralResponse
// @Router			/user [get]
func (h *UserHandler) GetAll(ctx echo.Context) error {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.ApiSum.WithLabelValues("GetAll").
			Observe(time.Since(t).Seconds())
	}()

	users, err := h.userUseCase.GetAll(ctx.Request().Context())
	if err != nil {
		return response.ResponseError(ctx, err)
	}

	var userList []models.User
	for _, user := range users {
		userList = append(userList, *user)
	}

	responsePayload := UserResponse{
		UserList: userList,
	}

	return response.ResponseSuccess(ctx, http.StatusOK, responsePayload)
}

// @Summary		Get user by ID
// @Description	Get a user by their ID
// @Tags			users
// @Accept			json
// @Produce		json
// @Accept			x-www-form-urlencoded
// @Param			id	path		string	true	"User ID"
// @Success		200	{object}	response.GeneralResponse
// @Failure		400	{object}	response.GeneralResponse
// @Failure		404	{object}	response.GeneralResponse
// @Failure		500	{object}	response.GeneralResponse
// @Router			/user/{id} [get]
func (h *UserHandler) GetByID(ctx echo.Context) error {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.ApiSum.WithLabelValues("GetByID").
			Observe(time.Since(t).Seconds())
	}()

	idParam := ctx.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return response.ResponseError(ctx, response.NewBadRequestError("invalid user ID format"))
	}

	user, err := h.userUseCase.GetUserById(ctx.Request().Context(), id)
	if err != nil {
		return response.ResponseError(ctx, response.NewNotFoundError("user not found"))
	}

	responsePayload := UserResponse{
		UserList: []models.User{*user},
	}

	return response.ResponseSuccess(ctx, http.StatusOK, responsePayload)
}
