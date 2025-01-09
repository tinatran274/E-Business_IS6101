package services

import (
	"context"

	"github.com/google/uuid"

	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/models"
	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/utils/response"
)

type UserService struct {
	userRepo models.UserRepository
}

func NewUserService(userRepo models.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) GetMe(ctx context.Context, id uuid.UUID) (*models.User, error) {
	user, err := s.userRepo.GetUserById(ctx, id)
	if err != nil {
		return nil, response.NewInternalServerError(err)
	}

	return user, nil
}
