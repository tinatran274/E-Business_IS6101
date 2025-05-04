package handlers

import (
	"net/http"
	"strings"
	"time"

	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/usecases"
	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/utils/metrics"
	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/utils/response"
	"10.0.0.50/tuan.quang.tran/aioz-ads/models"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ProductVariantHandler struct {
	productVariantUseCase usecases.ProducetVariantUseCase
}

func NewProductVariantHandler(
	productVariantUseCase usecases.ProducetVariantUseCase,
) *ProductVariantHandler {
	return &ProductVariantHandler{
		productVariantUseCase: productVariantUseCase,
	}
}

//	@Summary		Create a new product variant
//	@Description	Create a new product variant
//	@Tags			product-variant
//	@Security		BasicAuth
//	@Security		Bearer
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		models.CreateProductVariantRequest	true	"Product variant data"
//	@Success		200		{object}	response.GeneralResponse
//	@Failure		400		{object}	response.GeneralResponse
//	@Failure		401		{object}	response.GeneralResponse
//	@Failure		403		{object}	response.GeneralResponse
//	@Failure		404		{object}	response.GeneralResponse
//	@Failure		500		{object}	response.GeneralResponse
//	@Router			/product-variant [post]
func (h *ProductVariantHandler) CreateProductVariant(ctx echo.Context) error {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.ApiSum.WithLabelValues("CreateProductVariant").
			Observe(time.Since(t).Seconds())
	}()

	authInfo, ok := ctx.Get(models.AuthInfoKey).(models.AuthenticationInfo)
	if !ok {
		return response.ResponseFailMessage(ctx, http.StatusUnauthorized, "Unauthorized.")
	}

	var variant models.CreateProductVariantRequest
	if err := ctx.Bind(&variant); err != nil {
		return response.ResponseFailMessage(ctx, http.StatusBadRequest, err.Error())
	}

	variant.ProductID = strings.TrimSpace(variant.ProductID)
	if len(variant.ProductID) == 0 {
		return response.ResponseError(
			ctx,
			response.NewBadRequestError("Product ID is required."),
		)
	}

	if variant.Description != nil {
		description := strings.TrimSpace(*variant.Description)
		if len(description) > models.LimitLongTextLength {
			return response.ResponseError(
				ctx,
				response.NewBadRequestError("Product variant description is too long."),
			)
		}

		var newDescription *string
		if len(description) > 0 {
			newDescription = &description
		}

		variant.Description = newDescription
	}

	variant.Color = strings.TrimSpace(variant.Color)
	if len(variant.Color) == 0 {
		return response.ResponseError(
			ctx,
			response.NewBadRequestError("Product variant color is required."),
		)
	}

	if len(variant.Color) > models.LimitTextLength {
		return response.ResponseError(
			ctx,
			response.NewBadRequestError("Product variant color is too long."),
		)
	}

	if variant.RetailPrice <= 0 {
		return response.ResponseError(
			ctx,
			response.NewBadRequestError("Product variant retail price must be greater than 0."),
		)
	}

	if variant.Stock < 0 {
		return response.ResponseError(
			ctx,
			response.NewBadRequestError("Product variant stock must be greater than or equal to 0."),
		)
	}

	productVariant, err := h.productVariantUseCase.CreateProductVariant(
		ctx.Request().Context(),
		variant,
		authInfo,
	)
	if err != nil {
		return response.ResponseError(ctx, err)
	}

	return response.ResponseSuccess(ctx, http.StatusCreated, productVariant)
}

//	@Summary		Get product variants by product ID
//	@Description	Get product variants by product ID
//	@Tags			product-variant
//	@Security		BasicAuth
//	@Security		Bearer
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Product ID (UUID)"
//	@Success		200	{object}	response.GeneralResponse
//	@Failure		400	{object}	response.GeneralResponse
//	@Failure		401	{object}	response.GeneralResponse
//	@Failure		403	{object}	response.GeneralResponse
//	@Failure		404	{object}	response.GeneralResponse
//	@Failure		500	{object}	response.GeneralResponse
//	@Router			/product-variant/product/{id} [get]
func (h *ProductVariantHandler) GetProductVariantsByProductID(ctx echo.Context) error {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.ApiSum.WithLabelValues("GetProductVariantsByProductID").
			Observe(time.Since(t).Seconds())
	}()

	id := ctx.Param("id")
	id = strings.TrimSpace(id)
	if len(id) == 0 {
		return response.ResponseError(
			ctx,
			response.NewBadRequestError("Product ID is required."),
		)
	}

	productID, err := uuid.Parse(id)
	if err != nil {
		return response.ResponseError(
			ctx,
			response.NewBadRequestError("Invalid product ID."),
		)
	}

	productVariants, err := h.productVariantUseCase.GetProductVariantsByProductID(
		ctx.Request().Context(),
		productID,
	)
	if err != nil {
		return response.ResponseError(ctx, err)
	}

	return response.ResponseSuccess(ctx, http.StatusOK, productVariants)
}

//	@Summary		Update a product variant
//	@Description	Update a product variant
//	@Tags			product-variant
//	@Security		BasicAuth
//	@Security		Bearer
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string								true	"Product variant ID (UUID)"
//	@Param			payload	body		models.CreateProductVariantRequest	true	"Product variant data"
//	@Success		200		{object}	models.ProductVariant
//	@Failure		400		{object}	response.GeneralResponse
//	@Failure		401		{object}	response.GeneralResponse
//	@Failure		403		{object}	response.GeneralResponse
//	@Failure		404		{object}	response.GeneralResponse
//	@Failure		500		{object}	response.GeneralResponse
//	@Router			/product-variant/{id} [put]
func (h *ProductVariantHandler) UpdateProductVariant(ctx echo.Context) error {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.ApiSum.WithLabelValues("UpdateProductVariant").
			Observe(time.Since(t).Seconds())
	}()

	authInfo, ok := ctx.Get(models.AuthInfoKey).(models.AuthenticationInfo)
	if !ok {
		return response.ResponseFailMessage(ctx, http.StatusUnauthorized, "Unauthorized.")
	}

	id := ctx.Param("id")
	id = strings.TrimSpace(id)
	if len(id) == 0 {
		return response.ResponseError(
			ctx,
			response.NewBadRequestError("Product variant ID is required."),
		)
	}

	productVariantID, err := uuid.Parse(id)
	if err != nil {
		return response.ResponseError(
			ctx,
			response.NewBadRequestError("Invalid product variant ID."),
		)
	}

	var payload models.CreateProductVariantRequest
	if err := ctx.Bind(&payload); err != nil {
		return response.ResponseFailMessage(ctx, http.StatusBadRequest, err.Error())
	}

	payload.ProductID = strings.TrimSpace(payload.ProductID)
	if len(payload.ProductID) == 0 {
		return response.ResponseError(
			ctx,
			response.NewBadRequestError("Product ID is required."),
		)
	}

	if payload.Description != nil {
		description := strings.TrimSpace(*payload.Description)
		if len(description) > models.LimitLongTextLength {
			return response.ResponseError(
				ctx,
				response.NewBadRequestError("Product variant description is too long."),
			)
		}

		var newDescription *string
		if len(description) > 0 {
			newDescription = &description
		}

		payload.Description = newDescription
	}

	payload.Color = strings.TrimSpace(payload.Color)
	if len(payload.Color) == 0 {
		return response.ResponseError(
			ctx,
			response.NewBadRequestError("Product variant color is required."),
		)
	}

	if len(payload.Color) > models.LimitTextLength {

		return response.ResponseError(
			ctx,
			response.NewBadRequestError("Product variant color is too long."),
		)
	}

	if payload.RetailPrice <= 0 {
		return response.ResponseError(
			ctx,
			response.NewBadRequestError("Product variant retail price must be greater than 0."),
		)
	}

	if payload.Stock < 0 {
		return response.ResponseError(
			ctx,
			response.NewBadRequestError("Product variant stock must be greater than or equal to 0."),
		)
	}

	productVariant, err := h.productVariantUseCase.UpdateProductVariant(
		ctx.Request().Context(),
		productVariantID,
		payload,
		authInfo,
	)
	if err != nil {
		return response.ResponseError(ctx, err)
	}

	return response.ResponseSuccess(ctx, http.StatusOK, productVariant)
}

//	@Summary		Delete a product variant
//	@Description	Delete a product variant
//	@Tags			product-variant
//	@Security		BasicAuth
//	@Security		Bearer
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Product variant ID (UUID)"
//	@Success		200	{object}	response.GeneralResponse
//	@Failure		400	{object}	response.GeneralResponse
//	@Failure		401	{object}	response.GeneralResponse
//	@Failure		403	{object}	response.GeneralResponse
//	@Failure		404	{object}	response.GeneralResponse
//	@Failure		500	{object}	response.GeneralResponse
//	@Router			/product-variant/{id} [delete]
func (h *ProductVariantHandler) DeleteProductVariant(ctx echo.Context) error {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.ApiSum.WithLabelValues("DeleteProductVariant").
			Observe(time.Since(t).Seconds())
	}()

	authInfo, ok := ctx.Get(models.AuthInfoKey).(models.AuthenticationInfo)
	if !ok {
		return response.ResponseFailMessage(ctx, http.StatusUnauthorized, "Unauthorized.")
	}

	id := ctx.Param("id")
	id = strings.TrimSpace(id)
	if len(id) == 0 {
		return response.ResponseError(
			ctx,
			response.NewBadRequestError("Product variant ID is required."),
		)
	}

	productVariantID, err := uuid.Parse(id)
	if err != nil {
		return response.ResponseError(
			ctx,
			response.NewBadRequestError("Invalid product variant ID."),
		)
	}

	err = h.productVariantUseCase.DeleteProductVariant(
		ctx.Request().Context(),
		productVariantID,
		authInfo,
	)
	if err != nil {
		return response.ResponseError(ctx, err)
	}

	return response.ResponseSuccess(ctx, http.StatusOK, nil)
}
