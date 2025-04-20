package actions

import (
	"context"

	"github.com/google/uuid"

	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/usecases"
	"10.0.0.50/tuan.quang.tran/aioz-ads/models"
)

type GetMeAction struct {
	userUseCase usecases.UserUseCase
}

func NewGetMeAction(userUseCase usecases.UserUseCase) *GetMeAction {
	return &GetMeAction{
		userUseCase: userUseCase,
	}
}

func (a *GetMeAction) Exec(ctx context.Context, id uuid.UUID) (*models.User, error) {
	return a.userUseCase.GetMe(ctx, id)
}
