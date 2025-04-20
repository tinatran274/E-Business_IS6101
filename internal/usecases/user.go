package usecases

import (
	"context"

	"github.com/google/uuid"

	"10.0.0.50/tuan.quang.tran/aioz-ads/models"
)

type UserUseCase interface {
	GetMe(context.Context, uuid.UUID) (*models.User, error)
	GetAll(context.Context) ([]*models.User, error)
	GetUserById(context.Context, uuid.UUID) (*models.User, error)
	CreateUser(context.Context, *models.Account, *models.User) error
}
