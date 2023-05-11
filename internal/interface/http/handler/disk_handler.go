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

type DiskHandler struct {
	diskUseCase usecases.DiskUseCase
}

func NewDiskHandler(du usecases.DiskUseCase) DiskHandler {
	return DiskHandler{
		diskUseCase: du,
	}
}
func (h *DiskHandler) Routes(app *gin.Engine) {
	disk := app.Group("/disks", middleware.Protected())

	disk.POST("/", h.PostDiskHandler)
	disk.GET("/:diskId", h.GetDiskById)
	disk.PUT("/:diskId", h.EditDisk)
	disk.DELETE("/:diskId", h.DeleteDisk)
}

func (h *DiskHandler) PostDiskHandler(c *gin.Context) {
	ctx := context.Background()
	disk := entities.Disk{}

	if err := c.ShouldBindJSON(&disk); err != nil {
		appError := common.NewError(err, common.ValidationError)
		_ = c.Error(appError)
		return
	}

	newDisk, err := h.diskUseCase.AddDisk(ctx, disk)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseJSON.Error("fail", err.Error(), nil))
		return
	}

	c.JSON(http.StatusCreated, helper.ResponseJSON.Success("success", newDisk))
}

func (h *DiskHandler) GetDiskById(c *gin.Context) {
	ctx := context.Background()
	diskId := c.Param("diskId")

	disk, err := h.diskUseCase.GetDiskById(ctx, diskId)
	if err != nil {
		c.JSON(http.StatusNotFound, helper.ResponseJSON.Error("fail", err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, helper.ResponseJSON.Success("success", disk))
}

func (h *DiskHandler) EditDisk(c *gin.Context) {
	ctx := context.Background()
	diskId := c.Param("diskId")
	disk := entities.Disk{}

	if err := c.ShouldBindJSON(&disk); err != nil {
		appError := common.NewError(err, common.ValidationError)
		_ = c.Error(appError)
		return
	}

	err := h.diskUseCase.EditDisk(ctx, diskId, disk)
	if err != nil {
		c.JSON(http.StatusNotFound, helper.ResponseJSON.Error("fail", err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, helper.ResponseJSON.SuccessWithMessage("success", "disk berhasil diperbaharui", nil))
}

func (h *DiskHandler) DeleteDisk(c *gin.Context) {
	ctx := context.Background()
	diskId := c.Param("diskId")

	err := h.diskUseCase.DeleteDisk(ctx, diskId)
	if err != nil {
		c.JSON(http.StatusNotFound, helper.ResponseJSON.Error("fail", err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, helper.ResponseJSON.SuccessWithMessage("success", "disk berhasil dihapus", nil))
}
