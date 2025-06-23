package handlers

import "10.0.0.50/tuan.quang.tran/aioz-ads/internal/usecases"

type OrderHandler struct {
	OrderUseCase usecases.OrderUseCase
}

func NewOrderHandler(
	OrderUseCase usecases.OrderUseCase,
) *OrderHandler {
	return &OrderHandler{
		OrderUseCase: OrderUseCase,
	}
}
