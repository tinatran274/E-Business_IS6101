package response

import (
	"fmt"
	"log/slog"
	"net/http"
	"reflect"

	"github.com/labstack/echo/v4"
	"github.com/mdobak/go-xerrors"
)

type GeneralResponse struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

type StackFrame struct {
	File     string `json:"file"`
	Line     int    `json:"line"`
	Function string `json:"function"`
}

func ResponseSuccess(
	ctx echo.Context,
	code int,
	data interface{},
) error {
	if data == nil {
		return ctx.JSON(code, GeneralResponse{
			Status:  "success",
			Message: "success",
		})
	}

	if reflect.ValueOf(data).Type().Kind() == reflect.String {
		return ctx.JSON(code, GeneralResponse{
			Status:  "success",
			Message: data.(string),
		})
	}

	return ctx.JSON(code, GeneralResponse{
		Status: "success",
		Data:   data,
	})
}

func NewHttpError(
	code int,
	internal error,
	message ...interface{},
) *echo.HTTPError {
	if len(message) == 0 && internal != nil {
		message = append(message, internal.Error())
	}

	rs := echo.NewHTTPError(code, message...)
	if internal != nil {
		rs.Internal = xerrors.New(internal)
	}

	return rs
}

func ResponseError(
	ctx echo.Context,
	err error,
) error {
	httpErr, ok := err.(*echo.HTTPError)
	if !ok {
		httpErr = NewHttpError(http.StatusInternalServerError, err, "Internal server error.")
	}

	status := "error"
	if httpErr.Code >= 400 && httpErr.Code < 500 {
		status = "fail"
	}

	var message string
	if reflect.ValueOf(httpErr.Message).Type().Kind() == reflect.String {
		message = httpErr.Message.(string)
	} else if httpErr.Internal != nil {
		message = httpErr.Internal.Error()
	}

	if httpErr.Internal != nil && status == "error" {
		data := []any{
			slog.Any("data", httpErr.Internal),
			slog.Any("message", httpErr.Message),
		}
		trace := xerrors.StackTrace(httpErr.Internal)
		if len(trace) != 0 {
			frames := trace.Frames()
			stackFrames := make([]StackFrame, 0, len(frames))
			for i := 2; i < len(frames)-3; i++ {
				stackFrames = append(stackFrames, StackFrame{
					File:     frames[i].File,
					Line:     frames[i].Line,
					Function: frames[i].Function,
				})
			}

			data = append(data, slog.Any("stack", stackFrames))
		} else {
			data = append(data, slog.Any("stack", "no stack trace"))
		}

		slog.ErrorContext(
			ctx.Request().Context(),
			httpErr.Internal.Error(),
			data...,
		)
	} else {
		slog.ErrorContext(
			ctx.Request().Context(),
			fmt.Sprint(httpErr.Message),
			slog.Any("data", "logic error"),
			slog.Any("message", httpErr.Message),
		)
	}

	return ctx.JSON(httpErr.Code, GeneralResponse{
		Status:  status,
		Message: message,
	})
}

func ResponseFailMessage(
	ctx echo.Context,
	code int,
	message string,
) error {
	status := "fail"
	if code >= 400 && code < 500 {
		status = "fail"
	}
	return ctx.JSON(code, GeneralResponse{
		Status:  status,
		Message: message,
	})
}
