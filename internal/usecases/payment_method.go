package usecases

import (
	"context"

	"10.0.0.50/tuan.quang.tran/aioz-ads/models"
	"github.com/google/uuid"
)

type PaymentMethodUseCase interface {
	GetPaymentMethods(
		ctx context.Context,
		filter models.FilterParams,
	) ([]*models.PaymentMethod, int, error)
	GetPaymentMethodByID(
		ctx context.Context,
		id uuid.UUID,
	) (*models.PaymentMethod, error)
}
