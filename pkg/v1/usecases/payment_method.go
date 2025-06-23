package usecases

import (
	"context"

	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/utils/response"
	"10.0.0.50/tuan.quang.tran/aioz-ads/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type PaymentMethodUseCase struct {
	PaymentMethodRepo models.PaymentMethodRepository
}

func NewPaymentMethodUseCase(
	paymentMethodRepo models.PaymentMethodRepository,
) *PaymentMethodUseCase {
	return &PaymentMethodUseCase{
		PaymentMethodRepo: paymentMethodRepo,
	}
}

func (s *PaymentMethodUseCase) GetPaymentMethods(
	ctx context.Context,
	filter models.FilterParams,
) ([]*models.PaymentMethod, int, error) {
	paymentMethods, err := s.PaymentMethodRepo.GetPaymentMethods(ctx, filter)
	if err != nil {
		return nil, 0, response.NewInternalServerError(err)
	}

	total, err := s.PaymentMethodRepo.CountPaymentMethods(ctx)
	if err != nil {
		return nil, 0, response.NewInternalServerError(err)
	}

	return paymentMethods, total, nil
}

func (s *PaymentMethodUseCase) GetPaymentMethodByID(
	ctx context.Context,
	id uuid.UUID,
) (*models.PaymentMethod, error) {
	paymentMethod, err := s.PaymentMethodRepo.GetPaymentMethodByID(ctx, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, response.NewNotFoundError("Payment method not found.")
		}

		return nil, response.NewInternalServerError(err)
	}

	return paymentMethod, nil
}
