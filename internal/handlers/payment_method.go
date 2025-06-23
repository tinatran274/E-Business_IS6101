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

type PaymentMethodHandler struct {
	PaymentMethodUseCase usecases.PaymentMethodUseCase
}

func NewPaymentMethodHandler(
	paymentMethodUseCase usecases.PaymentMethodUseCase,
) *PaymentMethodHandler {
	return &PaymentMethodHandler{
		PaymentMethodUseCase: paymentMethodUseCase,
	}
}

type PaymentMethodResponse struct {
	PaymentMethods []*models.PaymentMethod `json:"payment_methods"`
	Total          int                     `json:"total"`
}

// GetPaymentMethods godoc
//
//	@Summary		Get payment methods
//	@Description	Get payment methods
//	@Tags			payment_method
//	@Security		BasicAuth
//	@Security		Bearer
//	@Accept			json
//	@Produce		json
//	@Produce		json
//	@Param			limit		query		int		false	"Limit (1-100), default is 25"
//	@Param			offset		query		int		false	"Offset for pagination"
//	@Param			sort_by		query		string	false	"Sort direction: asc or desc, default is asc"
//	@Param			order_by	query		string	false	"Field to sort by, allowed: updated_at, default is updated_at"
//	@Success		200			{object}	response.GeneralResponse
//	@Failure		400			{object}	response.GeneralResponse
//	@Failure		404			{object}	response.GeneralResponse
//	@Failure		500			{object}	response.GeneralResponse
//	@Router			/payment_method [get]
func (h *PaymentMethodHandler) GetPaymentMethods(ctx echo.Context) error {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.ApiSum.WithLabelValues("GetPaymentMethods").
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

	paymentMethods, total, err := h.PaymentMethodUseCase.GetPaymentMethods(ctx.Request().Context(), params)
	if err != nil {
		return response.ResponseError(ctx, err)
	}

	return response.ResponseSuccess(
		ctx,
		http.StatusOK,
		PaymentMethodResponse{
			PaymentMethods: paymentMethods,
			Total:          total,
		},
	)
}

// GetPaymentMethodByID godoc
//
//	@Summary		Get payment method by ID
//	@Description	Get payment method by ID
//	@Tags			payment_method
//	@Security		BasicAuth
//	@Security		Bearer
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Payment method ID (UUID)"
//	@Success		200	{object}	response.GeneralResponse
//	@Failure		400	{object}	response.GeneralResponse
//	@Failure		404	{object}	response.GeneralResponse
//	@Failure		500	{object}	response.GeneralResponse
//	@Router			/payment_method/{id} [get]
func (h *PaymentMethodHandler) GetPaymentMethodByID(ctx echo.Context) error {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.ApiSum.WithLabelValues("GetPaymentMethodByID").
			Observe(time.Since(t).Seconds())
	}()

	idParam := ctx.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return response.ResponseError(ctx, response.NewBadRequestError("Invalid payment method ID format."))
	}

	paymentMethod, err := h.PaymentMethodUseCase.GetPaymentMethodByID(ctx.Request().Context(), id)
	if err != nil {
		return response.ResponseError(ctx, err)
	}

	return response.ResponseSuccess(ctx, http.StatusOK, paymentMethod)
}
