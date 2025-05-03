package handlers

import (
	"net/http"
	"time"

	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/usecases"
	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/utils/metrics"
	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/utils/response"
	"10.0.0.50/tuan.quang.tran/aioz-ads/models"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type DishHandler struct {
	dishUseCase usecases.DishUseCase
}

func NewDishHandler(dishUseCase usecases.DishUseCase) *DishHandler {
	return &DishHandler{
		dishUseCase: dishUseCase,
	}
}

type DishResponse struct {
	Dishes []*models.Dish `json:"dishes"`
	Total  int            `json:"total"`
}

// @Summary		Get dishes
// @Description	Retrieve a list of dishes with filtering and sorting
// @Tags			dish
// @Security		BasicAuth
// @Security		Bearer
// @Accept			json
// @Produce		json
// @Param			limit		query		int		false	"Limit (1-100), default is 25"
// @Param			offset		query		int		false	"Offset for pagination"
// @Param			keyword		query		string	false	"Search keyword for name"
// @Param			sort_by		query		string	false	"Sort direction: asc or desc"
// @Param			order_by	query		string	false	"Field to sort by, allowed: updated_at"
// @Success		200			{object}	response.GeneralResponse
// @Failure		400			{object}	response.GeneralResponse
// @Failure		500			{object}	response.GeneralResponse
// @Router			/dish [get]
func (h *DishHandler) GetDishs(ctx echo.Context) error {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.ApiSum.WithLabelValues("GetDishs").
			Observe(time.Since(t).Seconds())
	}()

	var params models.FilterParams
	if err := ctx.Bind(&params); err != nil {
		return response.ResponseError(
			ctx,
			response.NewBadRequestError("Invalid input format."),
		)
	}

	if params.Limit <= 0 || params.Limit > 100 {
		params.Limit = 25
	}

	if params.Offset < 0 {
		params.Offset = 0
	}

	_, ok := models.ValidSortBy[params.SortBy]
	if !ok {
		params.SortBy = models.SortByDefault
	}

	_, ok = models.ValidOrderBy[params.OrderBy]
	if !ok {
		params.OrderBy = models.OrderByDefault
	}

	dishes, total, err := h.dishUseCase.GetDishes(ctx.Request().Context(), params)
	if err != nil {
		return response.ResponseError(ctx, err)
	}

	return response.ResponseSuccess(ctx, http.StatusOK, DishResponse{
		Dishes: dishes,
		Total:  total,
	})
}

// @Summary		Get dish by ID
// @Description	Retrieve an dish by its ID
// @Tags			dish
// @Security		BasicAuth
// @Security		Bearer
// @Accept			json
// @Produce		json
// @Param			id	path		string	true	"Dish ID (UUID)"
// @Success		200	{object}	response.GeneralResponse
// @Failure		400	{object}	response.GeneralResponse
// @Failure		404	{object}	response.GeneralResponse
// @Failure		500	{object}	response.GeneralResponse
// @Router			/dish/{id} [get]
func (h *DishHandler) GetDishByID(ctx echo.Context) error {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.ApiSum.WithLabelValues("GetDishByID").
			Observe(time.Since(t).Seconds())
	}()

	idParam := ctx.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return response.ResponseError(ctx, response.NewBadRequestError("Invalid dish ID format."))
	}

	dish, err := h.dishUseCase.GetDishByID(ctx.Request().Context(), id)
	if err != nil {
		return response.ResponseError(ctx, err)
	}

	return response.ResponseSuccess(ctx, http.StatusOK, dish)
}

// @Summary		GetDishesByIngredientID
// @Description	Retrieve a list of dishes by ingredient ID
// @Tags			dish
// @Security		BasicAuth
// @Security		Bearer
// @Accept			json
// @Produce		json
// @Param			limit		query		int		false	"Limit (1-100), default is 25"
// @Param			offset		query		int		false	"Offset for pagination"
// @Param			sort_by		query		string	false	"Sort direction: asc or desc, default is asc"
// @Param			order_by	query		string	false	"Field to sort by, allowed: updated_at, default is updated_at"
// @Param			id			path		string	true	"Ingredient ID (UUID)"
// @Success		200			{object}	response.GeneralResponse
// @Failure		400			{object}	response.GeneralResponse
// @Failure		404			{object}	response.GeneralResponse
// @Failure		500			{object}	response.GeneralResponse
// @Router			/dish/{id}/ingredient [get]
func (h *DishHandler) GetDishesByIngredientID(ctx echo.Context) error {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.ApiSum.WithLabelValues("GetDishesByIngredientID").
			Observe(time.Since(t).Seconds())
	}()

	idParam := ctx.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return response.ResponseError(ctx, response.NewBadRequestError("Invalid ingredient ID format."))
	}

	var params models.FilterParams
	if err := ctx.Bind(&params); err != nil {
		return response.ResponseError(
			ctx,
			response.NewBadRequestError("Invalid input format."),
		)
	}

	if params.Limit <= 0 || params.Limit > 100 {
		params.Limit = 25
	}

	if params.Offset < 0 {
		params.Offset = 0
	}

	_, ok := models.ValidSortBy[params.SortBy]
	if !ok {
		params.SortBy = models.SortByDefault
	}

	_, ok = models.ValidOrderBy[params.OrderBy]
	if !ok {
		params.OrderBy = models.OrderByDefault
	}

	dishes, total, err := h.dishUseCase.GetDishesByIngredientID(
		ctx.Request().Context(),
		id,
		params,
	)
	if err != nil {
		return response.ResponseError(ctx, err)
	}

	return response.ResponseSuccess(ctx, http.StatusOK, DishResponse{
		Dishes: dishes,
		Total:  total,
	})
}

// @Summary		Like dish
// @Description	Like an dish
// @Tags			dish
// @Security		BasicAuth
// @Security		Bearer
// @Accept			json
// @Produce		json
// @Param			id	path		string	true	"Dish ID (UUID)"
// @Success		200	{object}	response.GeneralResponse
// @Failure		400	{object}	response.GeneralResponse
// @Failure		404	{object}	response.GeneralResponse
// @Failure		500	{object}	response.GeneralResponse
// @Router			/dish/{id}/like [post]
func (h *DishHandler) LikeDish(ctx echo.Context) error {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.ApiSum.WithLabelValues("LikeDish").
			Observe(time.Since(t).Seconds())
	}()

	authInfo, ok := ctx.Get(models.AuthInfoKey).(models.AuthenticationInfo)
	if !ok {
		return response.ResponseFailMessage(ctx, http.StatusUnauthorized, "unauthorized")
	}

	idParam := ctx.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return response.ResponseError(ctx, response.NewBadRequestError("Invalid dish ID format."))
	}

	err = h.dishUseCase.LikeDish(ctx.Request().Context(), id, authInfo)
	if err != nil {
		return response.ResponseError(ctx, err)
	}

	return response.ResponseSuccess(ctx, http.StatusOK, "Dish liked successfully.")
}

// @Summary		Unlike dish
// @Description	Unlike an dish
// @Tags			dish
// @Security		BasicAuth
// @Security		Bearer
// @Accept			json
// @Produce		json
// @Param			id	path		string	true	"Dish ID (UUID)"
// @Success		200	{object}	response.GeneralResponse
// @Failure		400	{object}	response.GeneralResponse
// @Failure		404	{object}	response.GeneralResponse
// @Failure		500	{object}	response.GeneralResponse
// @Router			/dish/{id}/unlike [post]
func (h *DishHandler) UnlikeDish(ctx echo.Context) error {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.ApiSum.WithLabelValues("UnlikeDish").
			Observe(time.Since(t).Seconds())
	}()

	authInfo, ok := ctx.Get(models.AuthInfoKey).(models.AuthenticationInfo)
	if !ok {
		return response.ResponseFailMessage(ctx, http.StatusUnauthorized, "Unauthorized.")
	}

	idParam := ctx.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return response.ResponseError(ctx, response.NewBadRequestError("Invalid dish ID format."))
	}

	err = h.dishUseCase.UnlikeDish(ctx.Request().Context(), id, authInfo)
	if err != nil {
		return response.ResponseError(ctx, err)
	}

	return response.ResponseSuccess(ctx, http.StatusOK, "Dish unliked successfully.")
}
