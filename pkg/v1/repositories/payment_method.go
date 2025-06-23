package repositories

import (
	"context"
	"time"

	db "10.0.0.50/tuan.quang.tran/aioz-ads/db/generated"
	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/utils/metrics"
	"10.0.0.50/tuan.quang.tran/aioz-ads/models"
	"github.com/google/uuid"
)

type PaymentMethodRepository struct {
	db *db.Queries
}

func NewPaymentMethodRepository(
	db *db.Queries,
) *PaymentMethodRepository {
	return &PaymentMethodRepository{db: db}
}

func (r *PaymentMethodRepository) GetPaymentMethods(
	ctx context.Context,
	filter models.FilterParams,
) ([]*models.PaymentMethod, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("GetPaymentMethods").
			Observe(time.Since(t).Seconds())
	}()

	paymentMethods, err := r.db.GetPaymentMethods(
		ctx,
		db.GetPaymentMethodsParams{
			Limit:   filter.Limit,
			Offset:  filter.Offset,
			SortBy:  filter.SortBy,
			OrderBy: filter.OrderBy,
		},
	)
	if err != nil {
		return nil, err
	}

	paymentMethodList := make([]*models.PaymentMethod, 0, len(paymentMethods))
	for _, paymentMethod := range paymentMethods {
		paymentMethodList = append(paymentMethodList, models.ToPaymentMethod(paymentMethod))
	}

	return paymentMethodList, nil
}

func (r *PaymentMethodRepository) CountPaymentMethods(
	ctx context.Context,
) (int, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("CountPaymentMethods").
			Observe(time.Since(t).Seconds())
	}()

	count, err := r.db.CountPaymentMethods(ctx)
	if err != nil {
		return 0, err
	}

	return int(count), nil
}

func (r *PaymentMethodRepository) GetPaymentMethodByID(
	ctx context.Context,
	id uuid.UUID,
) (*models.PaymentMethod, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("GetPaymentMethodByID").
			Observe(time.Since(t).Seconds())
	}()

	paymentMethod, err := r.db.GetPaymentMethodById(ctx, id)
	if err != nil {
		return nil, err
	}

	return models.ToPaymentMethod(paymentMethod), nil
}
