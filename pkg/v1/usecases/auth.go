package usecases

import (
	"context"

	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/utils/response"
	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/utils/token"
	"10.0.0.50/tuan.quang.tran/aioz-ads/models"
	"github.com/jackc/pgx/v5"
)

type AuthUseCase struct {
	accountRepo models.AccountRepository
	userRepo    models.UserRepository
}

func NewAuthUseCase(
	accountRepo models.AccountRepository,
	userRepo models.UserRepository,
) *AuthUseCase {
	return &AuthUseCase{
		accountRepo: accountRepo,
		userRepo:    userRepo,
	}
}

func (s *AuthUseCase) SignIn(
	ctx context.Context,
	email, password string,
) (string, error) {
	existingAccount, err := s.accountRepo.GetAccountByEmail(ctx, email)
	if err != nil && err != pgx.ErrNoRows {
		return "", response.NewInternalServerError(err)
	}

	if existingAccount == nil {
		return "", response.NewBadRequestError("Account not found.")
	}

	if err := existingAccount.CheckPassword(password); err != nil {
		return "", response.NewBadRequestError("Invalid password.")
	}

	token, err := token.GenerateJWT(existingAccount.ID.String())
	if err != nil {
		return "", response.NewInternalServerError(err)
	}

	return token, nil
}
