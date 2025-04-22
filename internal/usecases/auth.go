package usecases

import "context"

type AuthUseCase interface {
	SignIn(
		ctx context.Context,
		email, password string,
	) (string, error)
}
