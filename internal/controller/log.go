package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"GO-Payment/internal/usecase"
)

type LogController struct {
	logUsecase usecase.LogUsecase
}

func NewLogController(router *gin.RouterGroup, lu usecase.LogUsecase) {
	ch := &LogController{logUsecase: lu}
	router.GET("/log", ch.GetAllLogs)
}

func (lc *LogController) GetAllLogs(ctx *gin.Context) {
	logs, err := lc.logUsecase.GetAllLogs()
	if err != nil {
		FailedJSONResponse(ctx, http.StatusInternalServerError, "internal server error")
		return
	}

	SuccessJSONResponse(ctx, http.StatusOK, "logs retrieved successfully", logs)
}
