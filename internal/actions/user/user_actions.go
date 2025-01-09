package user_actions

import "10.0.0.50/tuan.quang.tran/aioz-ads/internal/services"

type UserActions struct {
	GetMeAction *GetMeAction
}

func NewUserActions(userService services.UserService) *UserActions {
	return &UserActions{
		GetMeAction: NewGetMeAction(userService),
	}
}
