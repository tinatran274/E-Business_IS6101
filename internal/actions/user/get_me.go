package user_actions

import (
	"context"

	"github.com/google/uuid"

	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/models"
	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/services"
)

type GetMeAction struct {
	userService services.UserService
}

func NewGetMeAction(userService services.UserService) *GetMeAction {
	return &GetMeAction{
		userService: userService,
	}
}

func (a *GetMeAction) Exec(ctx context.Context, id uuid.UUID) (*models.User, error) {
	return a.userService.GetMe(ctx, id)
}
