package handler

import (
	"context"
	"github.com/elsyarif/pms-api/internal/applications/usecases"
	"github.com/elsyarif/pms-api/internal/domain/entities"
	"github.com/elsyarif/pms-api/pkg/common"
	"github.com/elsyarif/pms-api/pkg/helper"
	"github.com/elsyarif/pms-api/pkg/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	userUseCase usecases.UserUseCase
}

func NewUserHandler(addUser usecases.UserUseCase) UserHandler {
	return UserHandler{userUseCase: addUser}
}

func (h *UserHandler) Routes(app *gin.Engine) {
	user := app.Group("/users")
	user.POST("", h.PostUserHandler)
	user.GET("/profile", middleware.Protected(), func(c *gin.Context) {
		user, _ := c.Get("users")
		c.JSON(http.StatusOK, gin.H{
			"user": user,
		})
	})
}

func (h *UserHandler) PostUserHandler(c *gin.Context) {
	ctx := context.Background()

	user := entities.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		appError := common.NewError(err, common.ValidationError)
		c.Error(appError)
		return
	}

	result, err := h.userUseCase.AddUser(ctx, user)
	if err != nil {
		c.Error(common.NewError(err, common.ResourceAlreadyExists))
		return
	}
	ss := entities.UserToResponse(result)
	c.JSON(http.StatusCreated, helper.ResponseJSON.Success("success", ss))
}
