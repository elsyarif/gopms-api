package handler

import (
	"context"
	"github.com/elsyarif/pms-api/internal/applications/usecases"
	"github.com/elsyarif/pms-api/pkg/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ReportHandler struct {
	reportUseCase usecases.ReportUseCase
}

func NewReportHandler(ru usecases.ReportUseCase) ReportHandler {
	return ReportHandler{
		reportUseCase: ru,
	}
}

func (h *ReportHandler) Routes(app *gin.Engine) {
	report := app.Group("/report")
	report.GET("/export-excel/:groupId/:start/:end", h.GetExportExcel)
}

func (h *ReportHandler) GetExportExcel(c *gin.Context) {
	ctx := context.Background()
	groupId := c.Param("groupId")
	start := c.Param("start")
	end := c.Param("end")

	err := h.reportUseCase.ExportExcel(ctx, groupId, start, end)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseJSON.Error("fail", err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, helper.ResponseJSON.SuccessWithMessage("success", "export excel success", nil))
}
