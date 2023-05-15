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

type InspectionHandler struct {
	inspectionUseCase usecases.InspectionUseCase
}

func NewInspectionHandler(iu usecases.InspectionUseCase) InspectionHandler {
	return InspectionHandler{
		inspectionUseCase: iu,
	}
}

func (h *InspectionHandler) Routes(app *gin.Engine) {
	inspection := app.Group("/inspections")
	inspection.POST("", h.PostInspectionHandler)
}

func (h *InspectionHandler) PostInspectionHandler(c *gin.Context) {
	ctx := context.Background()
	inspection := entities.InspectionRequest{}

	if err := c.ShouldBindJSON(&inspection); err != nil {
		appError := common.NewError(err, common.ValidationError)
		c.Error(appError)
		return
	}

	id, err := h.inspectionUseCase.AddInspection(ctx, inspection)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseJSON.Error("fail", err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, helper.ResponseJSON.SuccessWithMessage("success", "inspection berhasil disimpan", gin.H{
		"inspection_id": id,
	}))
}
