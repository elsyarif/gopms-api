package usecases

import (
	"context"
	"github.com/elsyarif/pms-api/internal/domain/entities"
	"github.com/elsyarif/pms-api/internal/domain/services"
	"github.com/elsyarif/pms-api/pkg/helper"
)

type AuthUseCae struct {
	authService services.AuthService
	userService services.UserService
}

func NewAuthUseCase(as services.AuthService, us services.UserService) AuthUseCae {
	return AuthUseCae{
		authService: as,
		userService: us,
	}
}

func (a *AuthUseCae) Login(ctx context.Context, auth entities.AuthRequest) (*entities.AuthResponse, error) {
	id, err := a.authService.CheckCredential(ctx, auth)
	if err != nil {
		return nil, err
	}

	user, err := a.userService.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	accessToken, _ := helper.GenerateToken(user.Id, user.Username, helper.AccessToken)
	refreshToken, _ := helper.GenerateToken(user.Id, user.Username, helper.RefreshToken)

	err = a.authService.AddToken(ctx, entities.RefreshTokenRequest{RefreshToken: refreshToken})
	if err != nil {
		return nil, err
	}

	token := entities.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return &token, nil
}

func (a *AuthUseCae) RefreshToken(ctx context.Context, token entities.RefreshTokenRequest) (*entities.RefreshTokenResponse, error) {
	decode, err := helper.VerifyToken(token.RefreshToken, helper.RefreshToken)
	if err != nil {
		return nil, err
	}

	err = a.authService.CheckAvailabilityToken(ctx, token)
	if err != nil {
		return nil, err
	}

	accessToken, err := helper.GenerateToken(decode.UserId, decode.Username, helper.AccessToken)
	if err != nil {
		return nil, err
	}

	return &entities.RefreshTokenResponse{AccessToken: accessToken}, nil
}

func (a *AuthUseCae) DeleteToken(ctx context.Context, token entities.RefreshTokenRequest) error {
	err := a.authService.CheckAvailabilityToken(ctx, token)
	if err != nil {
		return err
	}

	err = a.authService.DeleteToken(ctx, token)
	if err != nil {
		return err
	}
	return nil
}
