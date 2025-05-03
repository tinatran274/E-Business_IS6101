package handlers

import (
	"fmt"
	"net/http"
	"time"

	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/usecases"
	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/utils/metrics"
	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/utils/response"
	"10.0.0.50/tuan.quang.tran/aioz-ads/models"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type StatisticHandler struct {
	statisticUseCase usecases.StatisticUseCase
}

func NewStatisticHandler(statisticUseCase usecases.StatisticUseCase) *StatisticHandler {
	return &StatisticHandler{
		statisticUseCase: statisticUseCase,
	}
}

// GetStatisticByUserIdAndDate
//	@Summary		Get statistic by user id and date
//	@Description	get statistic by user id and date
//	@Tags			statistic
//	@Security		BasicAuth
//	@Security		Bearer
//	@Accept			json
//	@Produce		json
//	@Param			user_id		path		string	true	"User ID"
//	@Param			updated_at	query		string	true	"Updated At"	Format(date-time)
//	@Success		200			{object}	response.GeneralResponse
//	@Failure		400			{object}	response.GeneralResponse
//	@Failure		500			{object}	response.GeneralResponse
//	@Router			/statistic/{user_id} [get]
func (h *StatisticHandler) GetStatisticByUserIdAndDate(ctx echo.Context) error {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.ApiSum.WithLabelValues("GetStatisticByUserIdAndDate").
			Observe(time.Since(t).Seconds())
	}()

	authInfo, ok := ctx.Get(models.AuthInfoKey).(models.AuthenticationInfo)
	if !ok {
		return response.ResponseFailMessage(ctx, http.StatusUnauthorized, "unauthorized")
	}

	userID := ctx.Param("user_id")
	_, err := uuid.Parse(userID)
	if err != nil {
		return response.ResponseFailMessage(ctx, http.StatusBadRequest, "invalid user id")
	}

	updatedAtStr := ctx.QueryParam("updated_at")
	updatedAt, err := time.Parse(time.DateOnly, updatedAtStr)
	if err != nil {
		return response.ResponseFailMessage(ctx, http.StatusBadRequest, "invalid updated at")
	}

	statistic, err := h.statisticUseCase.GetStatisticByUserIdAndDate(
		ctx.Request().Context(),
		authInfo.User.ID,
		updatedAt,
	)
	if err != nil {
		return response.ResponseError(ctx, err)
	}

	return response.ResponseSuccess(ctx, http.StatusOK, statistic)
}

type UpdateStatisticRequest struct {
	MorningCalories  float64 `json:"morning_calories"`
	LunchCalories    float64 `json:"lunch_calories"`
	DinnerCalories   float64 `json:"dinner_calories"`
	SnackCalories    float64 `json:"snack_calories"`
	ExerciseCalories float64 `json:"exercise_calories"`
}

// UpdateStatisticByUserIdAndDate
//	@Summary		Update statistic by user id and date
//	@Description	update statistic by user id and date
//	@Tags			statistic
//	@Security		BasicAuth
//	@Security		Bearer
//	@Accept			json
//	@Produce		json
//	@Param			user_id	path		string					true	"User ID"
//	@Param			body	body		UpdateStatisticRequest	true	"Statistic"
//	@Success		200		{object}	response.GeneralResponse
//	@Failure		400		{object}	response.GeneralResponse
//	@Failure		500		{object}	response.GeneralResponse
//	@Router			/statistic/{user_id} [put]
func (h *StatisticHandler) UpdateStatisticByUserIdAndDate(ctx echo.Context) error {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.ApiSum.WithLabelValues("UpdateStatisticByUserIdAndDate").
			Observe(time.Since(t).Seconds())
	}()

	authInfo, ok := ctx.Get(models.AuthInfoKey).(models.AuthenticationInfo)
	if !ok {
		return response.ResponseFailMessage(ctx, http.StatusUnauthorized, "Unauthorized.")
	}

	userID := ctx.Param("user_id")
	_, err := uuid.Parse(userID)
	if err != nil {
		return response.ResponseFailMessage(ctx, http.StatusBadRequest, "Invalid user id.")
	}

	now := time.Now().UTC()
	today, err := time.Parse(time.DateOnly, now.Format(time.DateOnly))
	if err != nil {
		return response.ResponseFailMessage(ctx, http.StatusBadRequest, "Invalid date.")
	}

	var payload UpdateStatisticRequest
	if err := ctx.Bind(&payload); err != nil {
		return response.ResponseFailMessage(ctx, http.StatusBadRequest, "Invalid request body.")
	}

	if payload.MorningCalories < 0 ||
		payload.LunchCalories < 0 ||
		payload.DinnerCalories < 0 ||
		payload.SnackCalories < 0 ||
		payload.ExerciseCalories < 0 {
		return response.ResponseFailMessage(
			ctx,
			http.StatusBadRequest,
			"Invalid calories, calories must be greater than or equal to 0.",
		)
	}

	if payload.MorningCalories > models.LimitCalories ||
		payload.LunchCalories > models.LimitCalories ||
		payload.DinnerCalories > models.LimitCalories ||
		payload.SnackCalories > models.LimitCalories ||
		payload.ExerciseCalories > models.LimitCalories {
		return response.ResponseFailMessage(
			ctx,
			http.StatusBadRequest,
			fmt.Sprintf(
				"Invalid calories, calories must be less than or equal to %f.",
				models.LimitCalories,
			),
		)
	}

	statistic := models.NewStatistic(
		today,
		authInfo.User.ID,
		payload.MorningCalories,
		payload.LunchCalories,
		payload.DinnerCalories,
		payload.SnackCalories,
		payload.ExerciseCalories,
	)

	_, err = h.statisticUseCase.UpdateStatisticByUserIdAndDate(
		ctx.Request().Context(),
		statistic,
	)
	if err != nil {
		return response.ResponseError(ctx, err)
	}

	return response.ResponseSuccess(ctx, http.StatusOK, nil)
}
