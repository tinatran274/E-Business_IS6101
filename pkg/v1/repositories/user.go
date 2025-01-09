package repositories

import (
	"context"
	"time"

	"github.com/google/uuid"

	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/models"
	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/utils/metrics"
	"10.0.0.50/tuan.quang.tran/aioz-ads/pkg/v1/db"
)

type UserRepository struct {
	q *db.Queries
}

func NewUserRepository(q *db.Queries) models.UserRepository {
	return &UserRepository{
		q: q,
	}
}

func (r *UserRepository) Create(
	ctx context.Context,
	user *models.User,
) error {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("CreateUser").
			Observe(time.Since(t).Seconds())
	}()

	panic("implement me")
}

func (r *UserRepository) GetUserById(
	ctx context.Context,
	id uuid.UUID,
) (*models.User, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("GetUserById").
			Observe(time.Since(t).Seconds())
	}()

	user, err := r.q.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	return models.ToUser(user), nil
}

func (r *UserRepository) UpdateUser(
	ctx context.Context,
	user *models.User,
) error {
	panic("implement me")
}
