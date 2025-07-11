package usecases

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/utils/response"
	"10.0.0.50/tuan.quang.tran/aioz-ads/models"
)

type UserUseCase struct {
	userRepo    models.UserRepository
	accountRepo models.AccountRepository
}

func NewUserUseCase(
	userRepo models.UserRepository,
	accountRepo models.AccountRepository,
) *UserUseCase {
	return &UserUseCase{
		userRepo:    userRepo,
		accountRepo: accountRepo,
	}
}

func (s *UserUseCase) GetMe(ctx context.Context, id uuid.UUID) (*models.User, error) {
	user, err := s.userRepo.GetUserById(ctx, id)
	if err != nil {
		return nil, response.NewInternalServerError(err)
	}

	return user, nil
}

func (s *UserUseCase) GetAll(ctx context.Context) ([]*models.User, error) {
	user, err := s.userRepo.GetAllUser(ctx)
	if err != nil {
		return nil, response.NewInternalServerError(err)
	}

	return user, nil
}

func (s *UserUseCase) GetUserById(ctx context.Context, id uuid.UUID) (*models.User, error) {
	user, err := s.userRepo.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserUseCase) GetUserByAccountId(ctx context.Context, id uuid.UUID) (*models.User, error) {
	user, err := s.userRepo.GetUserByAccountId(ctx, id)
	if err != nil {
		return nil, response.NewInternalServerError(err)
	}

	return user, nil
}

func (s *UserUseCase) CreateUser(ctx context.Context, account *models.Account, user *models.User) error {
	existingAccount, err := s.accountRepo.GetAccountByEmail(ctx, account.Email)
	if err != nil && err != pgx.ErrNoRows {
		return response.NewInternalServerError(err)
	}

	if existingAccount != nil {
		return response.NewBadRequestError("Email already in use.")
	}

	err = s.userRepo.CreateUser(ctx, user)
	if err != nil {
		return response.NewInternalServerError(err)
	}

	err = s.accountRepo.CreateAccount(ctx, account)
	if err != nil {
		return response.NewInternalServerError(err)
	}

	return nil
}

func (s *UserUseCase) UpdateUser(
	ctx context.Context,
	id uuid.UUID,
	payload models.UpdateUserRequest,
) (*models.User, error) {
	user, err := s.userRepo.GetUserById(ctx, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, response.NewNotFoundError("User not found.")
		}

		return nil, response.NewInternalServerError(err)
	}

	user.FirstName = &payload.FirstName
	user.LastName = &payload.LastName
	user.Username = &payload.Username
	age := int32(payload.Age)
	user.Age = &age
	height := int32(payload.Height)
	user.Height = &height
	weight := int32(payload.Weight)
	user.Weight = &weight
	user.UpdatedAt = time.Now().UTC()
	err = s.userRepo.UpdateUser(ctx, user)
	if err != nil {
		return nil, response.NewInternalServerError(err)
	}

	return user, nil
}
