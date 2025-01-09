package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewInternalServerError(err error) *echo.HTTPError {
	return NewHttpError(
		http.StatusInternalServerError,
		err,
		"Internal server error.",
	)
}
