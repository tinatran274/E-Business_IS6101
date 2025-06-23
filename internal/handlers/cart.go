package handlers

import (
	"net/http"
	"strconv"
	"time"

	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/usecases"
	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/utils/metrics"
	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/utils/response"
	"10.0.0.50/tuan.quang.tran/aioz-ads/models"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type CartHandler struct {
	cartUseCase usecases.CartUseCase
}

func NewCartHandler(
	cartUseCase usecases.CartUseCase,
) *CartHandler {
	return &CartHandler{
		cartUseCase: cartUseCase,
	}
}

// AddCartItem godoc
//
//	@Summary		Add cart item
//	@Description	add cart item
//	@Tags			cart
//	@Security		BasicAuth
//	@Security		Bearer
//	@Accept			json
//	@Produce		json
//	@Param			product_variant_id	query		string	true	"product_variant_id"
//	@Param			quantity			query		int32	true	"quantity"
//	@Success		200					{object}	models.Cart
//	@Failure		400					{object}	response.GeneralResponse
//	@Failure		500					{object}	response.GeneralResponse
//	@Router			/cart [post]
func (h *CartHandler) AddCartItem(ctx echo.Context) error {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.ApiSum.WithLabelValues("AddCartItem").
			Observe(time.Since(t).Seconds())
	}()

	authInfo, ok := ctx.Get("authInfo").(models.AuthenticationInfo)
	if !ok {
		return response.ResponseError(
			ctx,
			response.NewUnauthorizedError("Unauthorized."),
		)
	}

	productVariantID, err := uuid.Parse(ctx.QueryParam("product_variant_id"))
	if err != nil {
		return response.ResponseError(
			ctx,
			response.NewBadRequestError("Invalid product variant ID."),
		)
	}

	quantity, err := strconv.Atoi(ctx.QueryParam("quantity"))
	if err != nil || quantity <= 0 {
		return response.ResponseError(
			ctx,
			response.NewBadRequestError("Invalid quantity."),
		)
	}

	cartItem, err := h.cartUseCase.AddCartItem(
		ctx.Request().Context(),
		productVariantID,
		int32(quantity),
		authInfo,
	)
	if err != nil {
		return response.ResponseError(ctx, err)
	}

	return response.ResponseSuccess(ctx, http.StatusOK, cartItem)
}

type CartResponse struct {
	Carts []*models.Cart `json:"carts"`
	Total int            `json:"total"`
}

// GetCartItemsByUserID godoc
//
//	@Summary		Get cart items by user ID
//	@Description	get cart items by user ID
//	@Tags			cart
//	@Security		BasicAuth
//	@Security		Bearer
//	@Accept			json
//	@Produce		json
//	@Param			limit	query		int	false	"Limit (1-100), default is 25"
//	@Param			offset	query		int	false	"Offset for pagination"
//	@Success		200		{object}	response.GeneralResponse
//	@Failure		400		{object}	response.GeneralResponse
//	@Failure		404		{object}	response.GeneralResponse
//	@Failure		500		{object}	response.GeneralResponse
//	@Router			/cart [get]
func (h *CartHandler) GetCartItemsByUserID(ctx echo.Context) error {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.ApiSum.WithLabelValues("GetCartItemsByUserID").
			Observe(time.Since(t).Seconds())
	}()

	authInfo, ok := ctx.Get("authInfo").(models.AuthenticationInfo)
	if !ok {
		return response.ResponseError(
			ctx,
			response.NewUnauthorizedError("Unauthorized."),
		)
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

	cartItems, total, err := h.cartUseCase.GetCartItemsByUserID(
		ctx.Request().Context(),
		authInfo,
		params,
	)
	if err != nil {
		return response.ResponseError(ctx, err)
	}

	return response.ResponseSuccess(ctx, http.StatusOK, CartResponse{
		Carts: cartItems,
		Total: total,
	})
}

// DeleteCartItem godoc
//
//	@Summary		Delete cart item
//	@Description	delete cart item
//	@Tags			cart
//	@Security		BasicAuth
//	@Security		Bearer
//	@Accept			json
//	@Produce		json
//	@Param			product_variant_id	query		string	true	"product_variant_id"
//	@Success		200					{object}	response.GeneralResponse
//	@Failure		400					{object}	response.GeneralResponse
//	@Failure		404					{object}	response.GeneralResponse
//	@Failure		500					{object}	response.GeneralResponse
//	@Router			/cart [delete]
func (h *CartHandler) DeleteCartItem(ctx echo.Context) error {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.ApiSum.WithLabelValues("DeleteCartItem").
			Observe(time.Since(t).Seconds())
	}()

	authInfo, ok := ctx.Get("authInfo").(models.AuthenticationInfo)
	if !ok {
		return response.ResponseError(
			ctx,
			response.NewUnauthorizedError("Unauthorized."),
		)
	}

	productVariantID, err := uuid.Parse(ctx.QueryParam("product_variant_id"))
	if err != nil {
		return response.ResponseError(
			ctx,
			response.NewBadRequestError("Invalid product variant ID."),
		)
	}

	err = h.cartUseCase.DeleteCartItem(ctx.Request().Context(), productVariantID, authInfo)
	if err != nil {
		return response.ResponseError(ctx, err)
	}

	return response.ResponseSuccess(ctx, http.StatusOK, nil)
}

// UpdateCartItem godoc
//
//	@Summary		Update cart item
//	@Description	update cart item
//	@Tags			cart
//	@Security		BasicAuth
//	@Security		Bearer
//	@Accept			json
//	@Produce		json
//	@Param			product_variant_id	query		string	true	"product_variant_id"
//	@Param			quantity			query		int32	true	"quantity"
//	@Success		200					{object}	response.GeneralResponse
//	@Failure		400					{object}	response.GeneralResponse
//	@Failure		404					{object}	response.GeneralResponse
//	@Failure		500					{object}	response.GeneralResponse
//	@Router			/cart [put]
func (h *CartHandler) UpdateCartItem(ctx echo.Context) error {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.ApiSum.WithLabelValues("UpdateCartItem").
			Observe(time.Since(t).Seconds())
	}()

	authInfo, ok := ctx.Get("authInfo").(models.AuthenticationInfo)
	if !ok {
		return response.ResponseError(
			ctx,
			response.NewUnauthorizedError("Unauthorized."),
		)
	}

	productVariantID, err := uuid.Parse(ctx.QueryParam("product_variant_id"))
	if err != nil {
		return response.ResponseError(
			ctx,
			response.NewBadRequestError("Invalid product variant ID."),
		)
	}

	quantity, err := strconv.Atoi(ctx.QueryParam("quantity"))
	if err != nil || quantity <= 0 {
		return response.ResponseError(
			ctx,
			response.NewBadRequestError("Invalid quantity."),
		)
	}

	cartItem, err := h.cartUseCase.UpdateCartItem(
		ctx.Request().Context(),
		productVariantID,
		int32(quantity),
		authInfo,
	)
	if err != nil {
		return response.ResponseError(ctx, err)
	}

	return response.ResponseSuccess(ctx, http.StatusOK, cartItem)
}
