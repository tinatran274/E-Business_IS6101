package repositories

import (
	"context"
	"time"

	db "10.0.0.50/tuan.quang.tran/aioz-ads/db/generated"
	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/utils/metrics"
	"10.0.0.50/tuan.quang.tran/aioz-ads/models"
)

type AccountRepository struct {
	db *db.Queries
}

func NewAccountRepository(db *db.Queries) *AccountRepository {
	return &AccountRepository{
		db: db,
	}
}

func (r *AccountRepository) CreateAccount(
	ctx context.Context,
	account *models.Account,
) error {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("CreateAccount").
			Observe(time.Since(t).Seconds())
	}()

	params := db.CreateAccountParams{
		ID:        account.ID,
		UserID:    account.UserID,
		Email:     account.Email,
		Password:  account.Password,
		Status:    account.Status,
		CreatedAt: account.CreatedAt,
		CreatedBy: account.CreatedBy,
		UpdatedAt: account.UpdatedAt,
		UpdatedBy: account.UpdatedBy,
	}
	err := r.db.CreateAccount(ctx, params)
	if err != nil {
		return err
	}
	return nil
}

func (r *AccountRepository) GetAccountByEmail(
	ctx context.Context,
	email string,
) (*models.Account, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("GetAccountByEmail").
			Observe(time.Since(t).Seconds())
	}()

	account, err := r.db.GetAccountByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return models.ToAccount(account), nil
}
