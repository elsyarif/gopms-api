package usecases

import (
	"context"
	"github.com/elsyarif/pms-api/internal/domain/entities"
	"github.com/elsyarif/pms-api/internal/domain/services"
)

type UserUseCase struct {
	userService services.UserService
}

func NewUserUseCase(us services.UserService) UserUseCase {
	return UserUseCase{
		userService: us,
	}
}

func (u *UserUseCase) AddUser(ctx context.Context, user entities.User) (*entities.User, error) {
	createUser, err := u.userService.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return createUser, nil
}
