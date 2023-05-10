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

type GroupHandler struct {
	groupUseCase usecases.GroupUseCase
}

func NewGroupHandler(gu usecases.GroupUseCase) GroupHandler {
	return GroupHandler{
		groupUseCase: gu,
	}
}

func (h *GroupHandler) Routes(app *gin.Engine) {
	group := app.Group("/groups", middleware.Protected())

	group.POST("", h.PostGroupHandler)
	group.GET("", h.GetGroupHandler)
	group.GET("/:groupId", h.GetBuyIdGroupHandler)
	group.PUT("/:groupId", h.PutGroupHandler)
	group.DELETE("/:groupId", h.DeleteGroupHandler)
}

func (h *GroupHandler) PostGroupHandler(c *gin.Context) {
	ctx := context.Background()

	group := entities.Group{}
	if err := c.ShouldBindJSON(&group); err != nil {
		appError := common.NewError(err, common.ValidationError)
		_ = c.Error(appError)
		return
	}

	result, err := h.groupUseCase.AddGroup(ctx, group)
	if err != nil {
		c.Error(common.NewError(err, common.ResourceAlreadyExists))
		return
	}

	c.JSON(http.StatusCreated, helper.ResponseJSON.Success("success", result))
}

func (h *GroupHandler) GetGroupHandler(c *gin.Context) {
	ctx := context.Background()
	name := c.Query("name")

	group, err := h.groupUseCase.GetAllGroup(ctx, name)
	if err != nil {
		c.JSON(http.StatusNotFound, helper.ResponseJSON.Error("fail", err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, helper.ResponseJSON.Success("success", group))
}

func (h *GroupHandler) GetBuyIdGroupHandler(c *gin.Context) {
	ctx := context.Background()
	groupId := c.Param("groupId")

	group, err := h.groupUseCase.GetGroupById(ctx, groupId)
	if err != nil {
		c.JSON(http.StatusNotFound, helper.ResponseJSON.Error("fail", err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, helper.ResponseJSON.Success("success", group))
}

func (h *GroupHandler) PutGroupHandler(c *gin.Context) {
	ctx := context.Background()
	groupId := c.Param("groupId")

	group := entities.Group{}

	if err := c.ShouldBindJSON(&group); err != nil {
		appError := common.NewError(err, common.ValidationError)
		_ = c.Error(appError)
		return
	}

	err := h.groupUseCase.EditGroup(ctx, groupId, group)
	if err != nil {
		c.JSON(http.StatusNotFound, helper.ResponseJSON.Error("fail", err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, helper.ResponseJSON.SuccessWithMessage("success", "update group berhasil", nil))
}

func (h *GroupHandler) DeleteGroupHandler(c *gin.Context) {
	ctx := context.Background()
	groupId := c.Param("groupId")

	err := h.groupUseCase.DeleteGroup(ctx, groupId)
	if err != nil {
		c.JSON(http.StatusNotFound, helper.ResponseJSON.Error("fail", err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, helper.ResponseJSON.SuccessWithMessage("success", "group berhasil dihapus", nil))
}
