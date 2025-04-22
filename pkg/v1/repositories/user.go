package repositories

import (
	"context"
	"time"

	"github.com/google/uuid"

	db "10.0.0.50/tuan.quang.tran/aioz-ads/db/generated"
	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/utils/metrics"
	"10.0.0.50/tuan.quang.tran/aioz-ads/models"
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

func (r *UserRepository) GetUserByAccountId(
	ctx context.Context,
	id uuid.UUID,
) (*models.User, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("GetUserByAccountId").
			Observe(time.Since(t).Seconds())
	}()

	user, err := r.q.GetUserByAccountId(ctx, id)
	if err != nil {
		return nil, err
	}

	return models.ToUser(user), nil
}

func (r *UserRepository) GetAllUser(
	ctx context.Context) ([]*models.User, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("GetAllUser").
			Observe(time.Since(t).Seconds())
	}()

	rows, err := r.q.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	userList := make([]*models.User, 0, len(rows))
	for i := range rows {
		user := models.ToUser(rows[i])
		userList = append(userList, user)
	}

	return userList, nil
}

func (r *UserRepository) CreateUser(
	ctx context.Context,
	user *models.User,
) error {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("CreateUser").
			Observe(time.Since(t).Seconds())
	}()

	params := db.CreateUserParams{
		ID:     user.ID,
		Status: "active",
	}

	err := r.q.CreateUser(ctx, params)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) UpdateUser(
	ctx context.Context,
	user *models.User,
) error {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("UpdateUser").
			Observe(time.Since(t).Seconds())
	}()
	params := db.UpdateUserParams{
		ID:            user.ID,
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		Username:      user.Username,
		Age:           user.Age,
		Height:        user.Height,
		Weight:        user.Weight,
		Gender:        user.Gender,
		ExerciseLevel: user.ExerciseLevel,
		Aim:           user.Aim,
		Status:        user.Status,
	}
	err := r.q.UpdateUser(ctx, params)
	if err != nil {
		return err
	}

	return nil
}
