package models

import (
	"context"
	"time"

	db "10.0.0.50/tuan.quang.tran/aioz-ads/db/generated"
	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/utils/response"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AccountRepository interface {
	CreateAccount(
		ctx context.Context,
		account *Account,
	) error
	GetAccountByEmail(
		ctx context.Context,
		email string,
	) (*Account, error)
}

type Account struct {
	ID        uuid.UUID  `json:"id"`
	UserID    uuid.UUID  `json:"user_id"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	Status    string     `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	CreatedBy *uuid.UUID `json:"created_by"`
	UpdatedAt time.Time  `json:"updated_at"`
	UpdatedBy *uuid.UUID `json:"updated_by"`
	DeletedAt *time.Time `json:"deleted_at"`
	DeletedBy *uuid.UUID `json:"deleted_by"`
}

func NewAccount(
	email, password string,
	userId uuid.UUID,
) *Account {
	return &Account{
		ID:        uuid.New(),
		UserID:    userId,
		Email:     email,
		Password:  password,
		Status:    ActiveStatus,
		CreatedAt: time.Now().UTC(),
		CreatedBy: nil,
		UpdatedAt: time.Now().UTC(),
		UpdatedBy: nil,
	}
}

func ToAccount(a db.Account) *Account {
	return &Account{
		ID:        a.ID,
		UserID:    a.UserID,
		Email:     a.Email,
		Password:  a.Password,
		CreatedAt: a.CreatedAt,
		CreatedBy: a.CreatedBy,
		UpdatedAt: a.UpdatedAt,
		UpdatedBy: a.UpdatedBy,
		DeletedAt: a.DeletedAt,
		DeletedBy: a.DeletedBy,
	}
}

func (a *Account) CheckPassword(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(password)); err != nil {
		return response.NewBadRequestError("Invalid password.")
	}

	return nil
}
