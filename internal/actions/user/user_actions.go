package actions

import "10.0.0.50/tuan.quang.tran/aioz-ads/internal/usecases"

type UserActions struct {
	GetMeAction *GetMeAction
}

func NewUserActions(userUseCase usecases.UserUseCase) *UserActions {
	return &UserActions{
		GetMeAction: NewGetMeAction(userUseCase),
	}
}
