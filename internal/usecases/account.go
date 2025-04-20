package usecases

import "context"

type AccountUseCase interface {
	SignIn(
		ctx context.Context,
		email, password string,
	) (string, error)
}
