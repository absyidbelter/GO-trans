package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"GO-Payment/internal/usecase"
)

type UserController struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(router *gin.RouterGroup, cu usecase.UserUsecase) {
	ch := &UserController{userUsecase: cu}
	router.GET("/profile", ch.Profile)
}

func (uc *UserController) Profile(ctx *gin.Context) {
	res, err := uc.userUsecase.GetAllUsers()
	if err != nil {
		if err == usecase.ErrUsecaseNoData {
			FailedJSONResponse(ctx, http.StatusNotFound, "no data")
		} else {
			FailedJSONResponse(ctx, http.StatusInternalServerError,
				"internal server error")
		}

		return
	}

	SuccessJSONResponse(ctx, http.StatusOK, "", res)
}
