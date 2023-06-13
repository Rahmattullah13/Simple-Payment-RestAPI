package controller

import (
	"net/http"
	"simple-payment/delivery/middleware"
	"simple-payment/usecase"

	"github.com/gin-gonic/gin"
)

// @Summary Get all histories
// @Tags logs history
// @success 200 {object} model.LogHistoryResponse{data=[]model.LogHistory}
// @Router /api/logs/history [get]
type LogHistoryController struct {
	rg                *gin.RouterGroup
	logHistoryUseCase usecase.LogHistoryUsaCase
	tokenMdw          middleware.AuthTokenMiddleware
}

func (lhc *LogHistoryController) getLogHistory(ctx *gin.Context) {
	histories, err := lhc.logHistoryUseCase.LogHistory()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to get histories",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":   "Get all histories",
		"histories": histories,
	})
}

func NewLogHistoryController(rg *gin.RouterGroup, logHistoryUseCase usecase.LogHistoryUsaCase, tokenMdw middleware.AuthTokenMiddleware) *LogHistoryController {
	controller := LogHistoryController{
		rg:                rg,
		logHistoryUseCase: logHistoryUseCase,
		tokenMdw:          tokenMdw,
	}

	protectedGroup := rg.Group("", tokenMdw.RequiredToken())

	protectedGroup.GET("/logs/history", controller.getLogHistory)

	return &controller
}
