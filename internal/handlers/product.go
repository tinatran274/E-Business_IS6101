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

type ProductHandler struct {
	productUseCase usecases.ProductUseCase
}

func NewProductHandler(
	productUseCase usecases.ProductUseCase,
) *ProductHandler {
	return &ProductHandler{
		productUseCase: productUseCase,
	}
}

// Summary 			Create a new product
// Description 		Create a new product with the provided information
//
//	@Tags		product
//	@Security	BasicAuth
//	@Security	Bearer
//	@Accept		json
//	@Produce	json
//	@Param		product	body		models.CreateProductRequest	true	"Product information"
//	@Success	201		{object}	response.GeneralResponse
//	@Failure	400		{object}	response.GeneralResponse
//	@Failure	500		{object}	response.GeneralResponse
//	@Router		/product [post]
func (h *ProductHandler) CreateProduct(ctx echo.Context) error {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.ApiSum.WithLabelValues("CreateProduct").
			Observe(time.Since(t).Seconds())
	}()

	authInfo, ok := ctx.Get(models.AuthInfoKey).(models.AuthenticationInfo)
	if !ok {
		return response.ResponseFailMessage(ctx, http.StatusUnauthorized, "Unauthorized.")
	}

	var payload models.CreateProductRequest
	if err := ctx.Bind(&payload); err != nil {
		return response.ResponseError(
			ctx,
			response.NewBadRequestError("Invalid input format."),
		)
	}

	payload.Name = strings.TrimSpace(payload.Name)
	if len(payload.Name) == 0 {
		return response.ResponseError(
			ctx,
			response.NewBadRequestError("Product name is required."),
		)
	}

	if len(payload.Name) > models.LimitTextLength {
		return response.ResponseError(
			ctx,
			response.NewBadRequestError("Product name is too long."),
		)
	}

	if payload.Description != nil {
		description := strings.TrimSpace(*payload.Description)
		if len(description) > models.LimitLongTextLength {
			return response.ResponseError(
				ctx,
				response.NewBadRequestError("Product description is too long."),
			)
		}

		var newDescription *string
		if len(description) > 0 {
			newDescription = &description
		}

		payload.Description = newDescription
	}

	if payload.Brand != nil {
		brand := strings.TrimSpace(*payload.Brand)
		if len(brand) > models.LimitTextLength {
			return response.ResponseError(
				ctx,
				response.NewBadRequestError("Product brand is too long."),
			)
		}

		var newBrand *string
		if len(brand) > 0 {
			newBrand = &brand
		}

		payload.Brand = newBrand
	}

	if payload.Origin != nil {
		origin := strings.TrimSpace(*payload.Origin)
		if len(origin) > models.LimitTextLength {
			return response.ResponseError(
				ctx,
				response.NewBadRequestError("Product origin is too long."),
			)
		}

		var newOrigin *string
		if len(origin) > 0 {
			newOrigin = &origin
		}

		payload.Origin = newOrigin
	}

	if payload.UserGuide != nil {
		userGuide := strings.TrimSpace(*payload.UserGuide)
		if len(userGuide) > models.LimitLongTextLength {
			return response.ResponseError(
				ctx,
				response.NewBadRequestError("Product user guide is too long."),
			)
		}

		var newUserGuide *string
		if len(userGuide) > 0 {
			newUserGuide = &userGuide
		}

		payload.UserGuide = newUserGuide
	}

	payload.CategoryID = strings.TrimSpace(payload.CategoryID)
	if len(payload.CategoryID) == 0 {
		return response.ResponseError(
			ctx,
			response.NewBadRequestError("Product category ID is required."),
		)
	}

	if len(payload.ProductVariants) == 0 {
		return response.ResponseError(
			ctx,
			response.NewBadRequestError("Product variants are required, at least one variant is needed."),
		)
	}

	for _, variant := range payload.ProductVariants {
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
	}

	product, err := h.productUseCase.CreateProduct(
		ctx.Request().Context(),
		payload,
		authInfo,
	)
	if err != nil {
		return response.ResponseError(ctx, err)
	}

	return response.ResponseSuccess(ctx, http.StatusCreated, product)
}

// @Summary		Get product by ID
// @Description	Retrieve a product by its ID
// @Tags			product
// @Security		BasicAuth
// @Security		Bearer
// @Accept			json
// @Produce		json
// @Param			id	path		string	true	"Product ID (UUID)"
// @Success		200	{object}	response.GeneralResponse
// @Failure		400	{object}	response.GeneralResponse
// @Failure		404	{object}	response.GeneralResponse
// @Failure		500	{object}	response.GeneralResponse
// @Router			/product/{id} [get]
func (h *ProductHandler) GetProductByID(ctx echo.Context) error {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.ApiSum.WithLabelValues("GetProductByID").
			Observe(time.Since(t).Seconds())
	}()

	id := ctx.Param("id")
	id = strings.TrimSpace(id)
	productId, err := uuid.Parse(id)
	if err != nil {
		return response.ResponseError(
			ctx,
			response.NewBadRequestError("Invalid product ID."),
		)
	}

	product, err := h.productUseCase.GetProductByID(
		ctx.Request().Context(),
		productId,
	)
	if err != nil {
		return response.ResponseError(ctx, err)
	}

	return response.ResponseSuccess(ctx, http.StatusOK, product)
}

type ProductResponse struct {
	Products []*models.Product `json:"products"`
	Total    int               `json:"total"`
}

// @Summary		Get products
// @Description	Retrieve a list of products with optional filters
// @Tags			product
// @Security		BasicAuth
// @Security		Bearer
// @Accept			json
// @Produce		json
// @Param			category_id	query		string	false	"Category ID (UUID)"
// @Param			status		query		string	false	"Product status, allowed: pending, active, rejected"
// @Param			limit		query		int		false	"Limit (1-100), default is 25"
// @Param			offset		query		int		false	"Offset for pagination"
// @Param			keyword		query		string	false	"Search keyword for name"
// @Param			sort_by		query		string	false	"Sort direction: asc or desc"
// @Param			order_by	query		string	false	"Field to sort by, allowed: updated_at"
// @Success		200			{object}	response.GeneralResponse
// @Failure		400			{object}	response.GeneralResponse
// @Failure		500			{object}	response.GeneralResponse
// @Router			/product [get]
func (h *ProductHandler) GetProducts(ctx echo.Context) error {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.ApiSum.WithLabelValues("GetProducts").
			Observe(time.Since(t).Seconds())
	}()

	var params models.FilterParams
	if err := ctx.Bind(&params); err != nil {
		return response.ResponseError(
			ctx,
			response.NewBadRequestError("Invalid input format."),
		)
	}

	var categoryId *uuid.UUID
	categoryIdParam := ctx.QueryParam("category_id")
	categoryIdParam = strings.TrimSpace(categoryIdParam)
	if len(categoryIdParam) > 0 {
		categoryIdParsed, err := uuid.Parse(categoryIdParam)
		if err != nil {
			return response.ResponseError(
				ctx,
				response.NewBadRequestError("Invalid category ID."),
			)
		}

		categoryId = &categoryIdParsed
	}

	params.Status = strings.TrimSpace(params.Status)
	if len(params.Status) > 0 {
		if _, ok := models.ValidStatus[params.Status]; !ok {
			return response.ResponseError(
				ctx,
				response.NewBadRequestError("Invalid product status."),
			)
		}
	}

	if params.Limit <= 0 || params.Limit > 100 {
		params.Limit = 25
	}

	if params.Offset < 0 {
		params.Offset = 0
	}

	params.Keyword = strings.TrimSpace(params.Keyword)
	if len(params.Keyword) > 0 {
		if len(params.Keyword) > models.LimitTextLength {
			return response.ResponseError(
				ctx,
				response.NewBadRequestError("Keyword is too long."),
			)
		}
	}

	_, ok := models.ValidSortBy[params.SortBy]
	if !ok {
		params.SortBy = models.SortByDefault
	}

	_, ok = models.ValidOrderBy[params.OrderBy]
	if !ok {
		params.OrderBy = models.OrderByDefault
	}

	product, total, err := h.productUseCase.GetProducts(
		ctx.Request().Context(),
		categoryId,
		params,
	)
	if err != nil {
		return response.ResponseError(ctx, err)
	}

	return response.ResponseSuccess(
		ctx,
		http.StatusOK,
		ProductResponse{
			Products: product,
			Total:    total,
		},
	)
}

// @Summary		Get my products
// @Description	Retrieve a list of products created by the authenticated user
// @Tags			product
// @Security		BasicAuth
// @Security		Bearer
// @Accept			json
// @Produce		json
// @Param			category_id	query		string	false	"Category ID (UUID)"
// @Param			status		query		string	false	"Product status, allowed: pending, active, rejected"
// @Param			limit		query		int		false	"Limit (1-100), default is 25"
// @Param			offset		query		int		false	"Offset for pagination"
// @Param			keyword		query		string	false	"Search keyword for name"
// @Param			sort_by		query		string	false	"Sort direction: asc or desc"
// @Param			order_by	query		string	false	"Field to sort by, allowed: updated_at"
// @Success		200			{object}	response.GeneralResponse
// @Failure		400			{object}	response.GeneralResponse
// @Failure		500			{object}	response.GeneralResponse
// @Router			/product/me [get]
func (h *ProductHandler) GetMyProducts(ctx echo.Context) error {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.ApiSum.WithLabelValues("GetMyProducts").
			Observe(time.Since(t).Seconds())
	}()

	authInfo, ok := ctx.Get(models.AuthInfoKey).(models.AuthenticationInfo)
	if !ok {
		return response.ResponseFailMessage(ctx, http.StatusUnauthorized, "Unauthorized.")
	}

	var params models.FilterParams
	if err := ctx.Bind(&params); err != nil {
		return response.ResponseError(
			ctx,
			response.NewBadRequestError("Invalid input format."),
		)
	}

	var categoryId *uuid.UUID
	categoryIdParam := ctx.QueryParam("category_id")
	categoryIdParam = strings.TrimSpace(categoryIdParam)
	if len(categoryIdParam) > 0 {
		categoryIdParsed, err := uuid.Parse(categoryIdParam)
		if err != nil {
			return response.ResponseError(
				ctx,
				response.NewBadRequestError("Invalid category ID."),
			)
		}

		categoryId = &categoryIdParsed
	}

	params.Status = strings.TrimSpace(params.Status)
	if len(params.Status) > 0 {
		if _, ok := models.ValidStatus[params.Status]; !ok {
			return response.ResponseError(
				ctx,
				response.NewBadRequestError("Invalid product status."),
			)
		}
	}

	if params.Limit <= 0 || params.Limit > 100 {
		params.Limit = 25
	}

	if params.Offset < 0 {
		params.Offset = 0
	}

	params.Keyword = strings.TrimSpace(params.Keyword)
	if len(params.Keyword) > 0 {
		if len(params.Keyword) > models.LimitTextLength {
			return response.ResponseError(
				ctx,
				response.NewBadRequestError("Keyword is too long."),
			)
		}
	}

	_, ok = models.ValidSortBy[params.SortBy]
	if !ok {
		params.SortBy = models.SortByDefault
	}

	_, ok = models.ValidOrderBy[params.OrderBy]
	if !ok {
		params.OrderBy = models.OrderByDefault
	}

	product, total, err := h.productUseCase.GetMyProducts(
		ctx.Request().Context(),
		categoryId,
		authInfo.User.ID,
		params,
	)
	if err != nil {
		return response.ResponseError(ctx, err)
	}

	return response.ResponseSuccess(
		ctx,
		http.StatusOK,
		ProductResponse{
			Products: product,
			Total:    total,
		},
	)
}

// @Summary		Update product
// @Description	Update a product by its ID
// @Tags			product
// @Security		BasicAuth
// @Security		Bearer
// @Accept			json
// @Produce		json
// @Param			id		path		string						true	"Product ID (UUID)"
// @Param			product	body		models.CreateProductRequest	true	"Product information"
// @Success		200		{object}	response.GeneralResponse
// @Failure		400		{object}	response.GeneralResponse
// @Failure		404		{object}	response.GeneralResponse
// @Failure		500		{object}	response.GeneralResponse
// @Router			/product/{id} [put]
func (h *ProductHandler) UpdateProduct(ctx echo.Context) error {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.ApiSum.WithLabelValues("UpdateProduct").
			Observe(time.Since(t).Seconds())
	}()

	authInfo, ok := ctx.Get(models.AuthInfoKey).(models.AuthenticationInfo)
	if !ok {
		return response.ResponseFailMessage(ctx, http.StatusUnauthorized, "Unauthorized.")
	}

	id := ctx.Param("id")
	id = strings.TrimSpace(id)
	productId, err := uuid.Parse(id)
	if err != nil {
		return response.ResponseError(
			ctx,
			response.NewBadRequestError("Invalid product ID."),
		)
	}

	var payload models.CreateProductRequest
	if err := ctx.Bind(&payload); err != nil {
		return response.ResponseError(
			ctx,
			response.NewBadRequestError("Invalid input format."),
		)
	}

	payload.Name = strings.TrimSpace(payload.Name)
	if len(payload.Name) == 0 {
		return response.ResponseError(
			ctx,
			response.NewBadRequestError("Product name is required."),
		)
	}

	if len(payload.Name) > models.LimitTextLength {
		return response.ResponseError(
			ctx,
			response.NewBadRequestError("Product name is too long."),
		)
	}

	if payload.Description != nil {
		description := strings.TrimSpace(*payload.Description)
		if len(description) > models.LimitLongTextLength {
			return response.ResponseError(
				ctx,
				response.NewBadRequestError("Product description is too long."),
			)
		}

		var newDescription *string
		if len(description) > 0 {
			newDescription = &description
		}

		payload.Description = newDescription
	}

	if payload.Brand != nil {
		brand := strings.TrimSpace(*payload.Brand)
		if len(brand) > models.LimitTextLength {
			return response.ResponseError(
				ctx,
				response.NewBadRequestError("Product brand is too long."),
			)
		}

		var newBrand *string
		if len(brand) > 0 {
			newBrand = &brand
		}

		payload.Brand = newBrand
	}

	if payload.Origin != nil {
		origin := strings.TrimSpace(*payload.Origin)
		if len(origin) > models.LimitTextLength {
			return response.ResponseError(
				ctx,
				response.NewBadRequestError("Product origin is too long."),
			)
		}

		var newOrigin *string
		if len(origin) > 0 {
			newOrigin = &origin
		}

		payload.Origin = newOrigin
	}

	if payload.UserGuide != nil {
		userGuide := strings.TrimSpace(*payload.UserGuide)
		if len(userGuide) > models.LimitLongTextLength {
			return response.ResponseError(
				ctx,
				response.NewBadRequestError("Product user guide is too long."),
			)
		}

		var newUserGuide *string
		if len(userGuide) > 0 {
			newUserGuide = &userGuide
		}

		payload.UserGuide = newUserGuide
	}

	payload.CategoryID = strings.TrimSpace(payload.CategoryID)
	if len(payload.CategoryID) == 0 {
		return response.ResponseError(
			ctx,
			response.NewBadRequestError("Product category ID is required."),
		)
	}

	if len(payload.CategoryID) > models.LimitTextLength {
		return response.ResponseError(
			ctx,
			response.NewBadRequestError("Product category ID is too long."),
		)
	}

	if len(payload.ProductVariants) == 0 {
		return response.ResponseError(
			ctx,
			response.NewBadRequestError("Product variants are required, at least one variant is needed."),
		)
	}

	for _, variant := range payload.ProductVariants {
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
	}

	product, err := h.productUseCase.UpdateProduct(
		ctx.Request().Context(),
		productId,
		payload,
		authInfo,
	)
	if err != nil {
		return response.ResponseError(ctx, err)
	}

	return response.ResponseSuccess(ctx, http.StatusOK, product)
}

// @Summary		Delete product
// @Description	Delete product by id
// @Tag			product
// @Security		BasicAuth
// @Security		Bearer
// @Accept			json
// @Produce		json
// @Param			id	path		string	true	"Product ID (UUID)"
// @Success		200	{object}	response.GeneralResponse
// @Failure		400	{object}	response.GeneralResponse
// @Failure		404	{object}	response.GeneralResponse
// @Failure		500	{object}	response.GeneralResponse
// @Router			/product/{id} [delete]
func (h *ProductHandler) DeleteProduct(ctx echo.Context) error {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.ApiSum.WithLabelValues("DeleteProduct").
			Observe(time.Since(t).Seconds())
	}()

	authInfo, ok := ctx.Get(models.AuthInfoKey).(models.AuthenticationInfo)
	if !ok {
		return response.ResponseFailMessage(ctx, http.StatusUnauthorized, "Unauthorized.")
	}

	id := ctx.Param("id")
	id = strings.TrimSpace(id)
	productId, err := uuid.Parse(id)
	if err != nil {
		return response.ResponseError(
			ctx,
			response.NewBadRequestError("Invalid product ID."),
		)
	}

	err = h.productUseCase.DeleteProduct(
		ctx.Request().Context(),
		productId,
		authInfo,
	)
	if err != nil {
		return response.ResponseError(ctx, err)
	}

	return response.ResponseSuccess(ctx, http.StatusOK, nil)
}

// @Summary		Approve product
// @Description	Approve product by id
// @Tag			product
// @Security		BasicAuth
// @Security		Bearer
// @Accept			json
// @Produce		json
// @Param			id	path		string	true	"Product ID (UUID)"
// @Success		200	{object}	response.GeneralResponse
// @Failure		400	{object}	response.GeneralResponse
// @Failure		404	{object}	response.GeneralResponse
// @Failure		500	{object}	response.GeneralResponse
// @Router			/product/{id}/approve [put]
func (h *ProductHandler) ApproveProduct(ctx echo.Context) error {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.ApiSum.WithLabelValues("ApproveProduct").
			Observe(time.Since(t).Seconds())
	}()

	authInfo, ok := ctx.Get(models.AuthInfoKey).(models.AuthenticationInfo)
	if !ok {
		return response.ResponseFailMessage(ctx, http.StatusUnauthorized, "Unauthorized.")
	}

	id := ctx.Param("id")
	id = strings.TrimSpace(id)
	productId, err := uuid.Parse(id)
	if err != nil {
		return response.ResponseError(
			ctx,
			response.NewBadRequestError("Invalid product ID."),
		)
	}

	err = h.productUseCase.ApproveProduct(
		ctx.Request().Context(),
		productId,
		authInfo,
	)
	if err != nil {
		return response.ResponseError(ctx, err)
	}

	return response.ResponseSuccess(ctx, http.StatusOK, nil)
}

// @Summary		Reject product
// @Description	Reject product by id
// @Tag			product
// @Security		BasicAuth
// @Security		Bearer
// @Accept			json
// @Produce		json
// @Param			id	path		string	true	"Product ID (UUID)"
// @Success		200	{object}	response.GeneralResponse
// @Failure		400	{object}	response.GeneralResponse
// @Failure		404	{object}	response.GeneralResponse
// @Failure		500	{object}	response.GeneralResponse
// @Router			/product/{id}/reject [put]
func (h *ProductHandler) RejectProduct(ctx echo.Context) error {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.ApiSum.WithLabelValues("RejectProduct").
			Observe(time.Since(t).Seconds())
	}()

	authInfo, ok := ctx.Get(models.AuthInfoKey).(models.AuthenticationInfo)
	if !ok {
		return response.ResponseFailMessage(ctx, http.StatusUnauthorized, "Unauthorized.")
	}

	id := ctx.Param("id")
	id = strings.TrimSpace(id)
	productId, err := uuid.Parse(id)
	if err != nil {
		return response.ResponseError(
			ctx,
			response.NewBadRequestError("Invalid product ID."),
		)
	}

	err = h.productUseCase.RejectProduct(
		ctx.Request().Context(),
		productId,
		authInfo,
	)
	if err != nil {
		return response.ResponseError(ctx, err)
	}

	return response.ResponseSuccess(ctx, http.StatusOK, nil)
}
