package routes

import (
	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/handlers"
	"github.com/labstack/echo/v4"
)

type PaymentMethodRouter struct {
	PaymentMethodHandler *handlers.PaymentMethodHandler
	authMiddleware       echo.MiddlewareFunc
}

func NewPaymentMethodRouter(
	paymentMethodHandler *handlers.PaymentMethodHandler,
	authMiddleware echo.MiddlewareFunc,
) *PaymentMethodRouter {
	return &PaymentMethodRouter{
		PaymentMethodHandler: paymentMethodHandler,
		authMiddleware:       authMiddleware,
	}
}

func (r *PaymentMethodRouter) Register(e *echo.Group) {
	paymentMethodRoute := e.Group("/payment-methods")
	if r.authMiddleware != nil {
		paymentMethodRoute.Use(r.authMiddleware)
	}

	paymentMethodRoute.GET("", r.PaymentMethodHandler.GetPaymentMethods)
	paymentMethodRoute.GET("/:id", r.PaymentMethodHandler.GetPaymentMethodByID)
}
