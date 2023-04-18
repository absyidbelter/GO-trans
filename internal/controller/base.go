package controller

import (
	"GO-Payment/internal/model/dto"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SuccessJSONResponse(ctx *gin.Context, code int, msg string, data any) {
	ctx.JSON(code, dto.NewApiResponseSuccess(msg, data))
}

func FailedJSONResponse(ctx *gin.Context, code int, msg string) {
	ctx.JSON(code, dto.NewApiResponseFailed(msg))
}

func GetUserId(ctx *gin.Context) (int, error) {
	user_id, ok := ctx.Get("user_id")
	if !ok {
		FailedJSONResponse(ctx, http.StatusBadRequest, "invalid user id")
		return -1, errors.New("invalid user id")
	}

	id, err := strconv.Atoi(user_id.(string))
	if err != nil {
		FailedJSONResponse(ctx, http.StatusBadRequest, "invalid user id")
		return -1, errors.New("invalid user id")
	}

	return id, nil
}
