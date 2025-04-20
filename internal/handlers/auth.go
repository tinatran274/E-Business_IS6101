package handlers

import (
	"net/http"
	"net/mail"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"

	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/usecases"
	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/utils/metrics"
	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/utils/response"
	"10.0.0.50/tuan.quang.tran/aioz-ads/models"
)

type AuthHandler struct {
	userUseCase usecases.UserUseCase
	authUseCase usecases.AccountUseCase
}

func NewAuthHandler(
	userUseCase usecases.UserUseCase,
	authUseCase usecases.AccountUseCase,
) *AuthHandler {
	return &AuthHandler{
		userUseCase: userUseCase,
		authUseCase: authUseCase,
	}
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SignUp godoc
//
//	@Summary		Sign up
//	@Description	sign up new user
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Accept			x-www-form-urlencoded
//	@Param			payload	body		RegisterRequest	true	"input"
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

	var req RegisterRequest
	if err := ctx.Bind(&req); err != nil {
		return response.ResponseError(
			ctx,
			response.NewBadRequestError("Invalid input format."),
		)
	}

	req.Email = strings.TrimSpace(req.Email)
	req.Password = strings.TrimSpace(req.Password)
	if req.Email == "" || req.Password == "" {
		return response.ResponseError(
			ctx,
			response.NewBadRequestError("Email and password are required."),
		)
	}

	_, err := mail.ParseAddress(req.Email)
	if err != nil {
		return response.ResponseError(
			ctx,
			response.NewBadRequestError("Invalid email format."),
		)
	}

	if len(req.Password) < 6 {
		return response.ResponseError(
			ctx,
			response.NewBadRequestError("Password must be at least 6 characters."),
		)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return response.ResponseError(ctx, response.NewInternalServerError(err))
	}

	user := models.NewUser(
		nil,
		nil,
		nil,
		&models.DefaultAge,
		&models.DefaultHeight,
		&models.DefaultWeight,
		&models.DefaultGender,
		&models.DefaultExerciseLevel,
		&models.DefaultAim,
	)
	account := models.NewAccount(
		req.Email,
		string(hashedPassword),
		user.ID,
	)
	if err := h.userUseCase.CreateUser(ctx.Request().Context(), account, user); err != nil {
		return response.ResponseError(ctx, err)
	}

	return response.ResponseSuccess(ctx, http.StatusOK, user)
}

// SignIn godoc
//
//	@Summary		Sign in
//	@Description	sign in user
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Accept			x-www-form-urlencoded
//	@Param			payload	body		RegisterRequest	true	"input"
//	@Success		200		{object}	models.User
//	@Failure		400		{object}	response.GeneralResponse
//	@Failure		500		{object}	response.GeneralResponse
//	@Router			/auth/sign-in [post]
func (h *AuthHandler) SignIn(ctx echo.Context) error {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.ApiSum.WithLabelValues("SignIn").
			Observe(time.Since(t).Seconds())
	}()

	var req RegisterRequest
	if err := ctx.Bind(&req); err != nil {
		return response.ResponseError(
			ctx,
			response.NewBadRequestError("Invalid input format."),
		)
	}

	req.Email = strings.TrimSpace(req.Email)
	req.Password = strings.TrimSpace(req.Password)
	if req.Email == "" || req.Password == "" {
		return response.ResponseError(
			ctx,
			response.NewBadRequestError("Email and password are required."),
		)
	}

	_, err := mail.ParseAddress(req.Email)
	if err != nil {
		return response.ResponseError(
			ctx,
			response.NewBadRequestError("Invalid email format."),
		)
	}

	if len(req.Password) < 6 {
		return response.ResponseError(
			ctx,
			response.NewBadRequestError("Password must be at least 6 characters."),
		)
	}

	accessToken, err := h.authUseCase.SignIn(
		ctx.Request().Context(),
		req.Email,
		req.Password,
	)
	if err != nil {
		return response.ResponseError(ctx, err)
	}

	return response.ResponseSuccess(ctx, http.StatusOK, accessToken)
}
