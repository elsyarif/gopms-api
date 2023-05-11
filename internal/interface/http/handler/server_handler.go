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

type ServerHandler struct {
	serverUseCase usecases.ServerUseCase
}

func NewServerHandler(su usecases.ServerUseCase) ServerHandler {
	return ServerHandler{
		serverUseCase: su,
	}
}

func (h *ServerHandler) Routes(app *gin.Engine) {
	server := app.Group("/servers", middleware.Protected())

	server.POST("", h.PostServerHandler)
	server.GET("/:serverId", h.GetServerById)
	server.GET("/:serverId/disks", h.GetServerDisk)
	server.PUT("/:serverId", h.PutServerHandler)
	server.DELETE("/:serverId", h.DeleteServerHandler)
}

func (h *ServerHandler) PostServerHandler(c *gin.Context) {
	ctx := context.Background()
	server := entities.Server{}

	if err := c.ShouldBindJSON(&server); err != nil {
		appError := common.NewError(err, common.ValidationError)
		_ = c.Error(appError)
		return
	}

	newServer, err := h.serverUseCase.AddServer(ctx, server)
	if err != nil {
		appError := common.NewError(err, common.ResourceAlreadyExists)
		_ = c.Error(appError)
		return
	}

	c.JSON(http.StatusCreated, helper.ResponseJSON.Success("success", newServer))
}

func (h *ServerHandler) GetServerById(c *gin.Context) {
	ctx := context.Background()
	serverId := c.Param("serverId")

	server, err := h.serverUseCase.GetServerById(ctx, serverId)
	if err != nil {
		appError := common.NewError(err, common.NotFoundError)
		c.Error(appError)
		return
	}

	c.JSON(http.StatusOK, helper.ResponseJSON.Success("success", server))
}

func (h *ServerHandler) PutServerHandler(c *gin.Context) {
	ctx := context.Background()
	serverId := c.Param("serverId")

	server := entities.Server{}
	if err := c.ShouldBindJSON(&server); err != nil {
		appError := common.NewError(err, common.ValidationError)
		_ = c.Error(appError)
		return
	}

	err := h.serverUseCase.EditServer(ctx, serverId, server)
	if err != nil {
		appError := common.NewError(err, common.NotFoundError)
		_ = c.Error(appError)
		return
	}

	c.JSON(http.StatusOK, helper.ResponseJSON.SuccessWithMessage("success", "server berhasil diperbaharui", nil))
}

func (h *ServerHandler) DeleteServerHandler(c *gin.Context) {
	ctx := context.Background()
	serverId := c.Param("serverId")

	err := h.serverUseCase.DeleteServer(ctx, serverId)
	if err != nil {
		appError := common.NewError(err, common.NotFoundError)
		_ = c.Error(appError)
		return
	}

	c.JSON(http.StatusOK, helper.ResponseJSON.SuccessWithMessage("success", "server berhasil dihapus", nil))
}

func (h *ServerHandler) GetServerDisk(c *gin.Context) {
	ctx := context.Background()
	serverId := c.Param("serverId")

	server, err := h.serverUseCase.GetServerDisk(ctx, serverId)
	if err != nil {
		appError := common.NewError(err, common.NotFoundError)
		c.Error(appError)
		return
	}

	c.JSON(http.StatusOK, helper.ResponseJSON.Success("success", server))
}
