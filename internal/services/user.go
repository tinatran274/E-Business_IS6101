package services

import (
	"context"

	"github.com/google/uuid"

	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/models"
)

type UserService interface {
	GetMe(ctx context.Context, id uuid.UUID) (*models.User, error)
}
