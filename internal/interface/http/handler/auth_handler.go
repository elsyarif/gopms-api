package handler

import (
	"context"
	"github.com/elsyarif/pms-api/internal/applications/usecases"
	"github.com/elsyarif/pms-api/internal/domain/entities"
	"github.com/elsyarif/pms-api/pkg/common"
	"github.com/elsyarif/pms-api/pkg/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthHandler struct {
	authUseCase usecases.AuthUseCae
}

func NewAuthHandler(auth usecases.AuthUseCae) AuthHandler {
	return AuthHandler{authUseCase: auth}
}

func (h *AuthHandler) Routes(app *gin.Engine) {
	auth := app.Group("/authentication")

	auth.POST("", h.PostAuthHandler)
	auth.PUT("", h.PutAuthHandler)
	auth.DELETE("", h.DeleteAuthHandler)
}

func (h *AuthHandler) PostAuthHandler(c *gin.Context) {
	ctx := context.Background()

	auth := entities.AuthRequest{}
	if err := c.ShouldBindJSON(&auth); err != nil {
		appError := common.NewError(err, common.ValidationError)
		c.Error(appError)
		return
	}

	response, err := h.authUseCase.Login(ctx, auth)
	if err != nil {
		c.Error(common.NewError(err, common.NotAuthenticatedError))
		return
	}

	c.JSON(http.StatusOK, helper.ResponseJSON.Success("success", response))
}

func (h *AuthHandler) PutAuthHandler(c *gin.Context) {
	ctx := context.Background()

	auth := entities.RefreshTokenRequest{}
	if err := c.ShouldBindJSON(&auth); err != nil {
		appError := common.NewError(err, common.ValidationError)
		c.Error(appError)
		return
	}

	token, err := h.authUseCase.RefreshToken(ctx, auth)
	if err != nil {
		c.Error(common.NewError(err, common.InvalidTokenError))
		return
	}

	c.JSON(http.StatusOK, helper.ResponseJSON.Success("success", token))
}

func (h *AuthHandler) DeleteAuthHandler(c *gin.Context) {
	ctx := context.Background()

	auth := entities.RefreshTokenRequest{}
	if err := c.ShouldBindJSON(&auth); err != nil {
		appError := common.NewError(err, common.ValidationError)
		c.Error(appError)
		return
	}

	err := h.authUseCase.DeleteToken(ctx, auth)
	if err != nil {
		c.Error(common.NewError(err, common.InvalidTokenError))
		return
	}

	c.JSON(http.StatusOK, helper.Response{
		Status:  "success",
		Message: "logout berhasil",
	})
}
