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

type IngredientHandler struct {
	ingredientUseCase usecases.IngredientUseCase
}

func NewIngredientHandler(ingredientUseCase usecases.IngredientUseCase) *IngredientHandler {
	return &IngredientHandler{
		ingredientUseCase: ingredientUseCase,
	}
}

type IngredientResponse struct {
	Inredients []*models.Ingredient `json:"ingredients"`
	Total      int                  `json:"total"`
}

//	@Summary		Get ingredients
//	@Description	Retrieve a list of ingredients with filtering and sorting
//	@Tags			ingredients
//	@Security		BasicAuth
//	@Security		Bearer
//	@Accept			json
//	@Produce		json
//	@Param			limit		query		int		false	"Limit (1-100), default is 25"
//	@Param			offset		query		int		false	"Offset for pagination"
//	@Param			keyword		query		string	false	"Search keyword for name"
//	@Param			sort_by		query		string	false	"Sort direction: asc or desc"
//	@Param			order_by	query		string	false	"Field to sort by, allowed: updated_at"
//	@Success		200			{object}	response.GeneralResponse
//	@Failure		400			{object}	response.GeneralResponse
//	@Failure		500			{object}	response.GeneralResponse
//	@Router			/ingredient [get]
func (h *IngredientHandler) GetIngredients(ctx echo.Context) error {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.ApiSum.WithLabelValues("GetIngredients").
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
		return response.ResponseError(
			ctx,
			response.NewBadRequestError("Invalid sort by value."),
		)
	}

	_, ok = models.ValidOrderBy[params.OrderBy]
	if !ok {
		return response.ResponseError(
			ctx,
			response.NewBadRequestError("Invalid order by value."),
		)
	}

	ingredients, total, err := h.ingredientUseCase.GetIngredients(ctx.Request().Context(), params)
	if err != nil {
		return response.ResponseError(ctx, err)
	}

	return response.ResponseSuccess(ctx, http.StatusOK, IngredientResponse{
		Inredients: ingredients,
		Total:      total,
	})
}

//	@Summary		Get ingredient by ID
//	@Description	Retrieve an ingredient by its ID
//	@Tags			ingredients
//	@Security		BasicAuth
//	@Security		Bearer
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Ingredient ID (UUID)"
//	@Success		200	{object}	response.GeneralResponse
//	@Failure		400	{object}	response.GeneralResponse
//	@Failure		404	{object}	response.GeneralResponse
//	@Failure		500	{object}	response.GeneralResponse
//	@Router			/ingredient/{id} [get]
func (h *IngredientHandler) GetIngredientByID(ctx echo.Context) error {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.ApiSum.WithLabelValues("GetIngredientByID").
			Observe(time.Since(t).Seconds())
	}()

	idParam := ctx.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return response.ResponseError(ctx, response.NewBadRequestError("Invalid ingredient ID format."))
	}

	ingredient, err := h.ingredientUseCase.GetIngredientByID(ctx.Request().Context(), id)
	if err != nil {
		return response.ResponseError(ctx, err)
	}

	return response.ResponseSuccess(ctx, http.StatusOK, ingredient)
}
