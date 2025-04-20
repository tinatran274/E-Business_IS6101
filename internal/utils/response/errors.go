package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Internal Server Error (500)
func NewInternalServerError(err error) *echo.HTTPError {
	return NewHttpError(
		http.StatusInternalServerError,
		err,
		"Internal server error.",
	)
}

// Bad Request (400)
func NewBadRequestError(message string) *echo.HTTPError {
	return NewHttpError(
		http.StatusBadRequest,
		nil,
		message,
	)
}

// Unauthorized (401)
func NewUnauthorizedError(message string) *echo.HTTPError {
	return NewHttpError(
		http.StatusUnauthorized,
		nil,
		message,
	)
}

// Forbidden (403)
func NewForbiddenError(message string) *echo.HTTPError {
	return NewHttpError(
		http.StatusForbidden,
		nil,
		message,
	)
}

// Not Found (404)
func NewNotFoundError(message string) *echo.HTTPError {
	return NewHttpError(
		http.StatusNotFound,
		nil,
		message,
	)
}

// Conflict (409)
func NewConflictError(message string) *echo.HTTPError {
	return NewHttpError(
		http.StatusConflict,
		nil,
		message,
	)
}

// Unprocessable Entity (422)
func NewUnprocessableEntityError(message string) *echo.HTTPError {
	return NewHttpError(
		http.StatusUnprocessableEntity,
		nil,
		message,
	)
}

// Custom Error for any HTTP status
func NewCustomError(statusCode int, message string) *echo.HTTPError {
	return NewHttpError(
		statusCode,
		nil,
		message,
	)
}
