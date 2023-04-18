package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"GO-Payment/config"
	"GO-Payment/internal/model/dto"
	model "GO-Payment/internal/model/entity"
	"GO-Payment/internal/usecase"
	"GO-Payment/pkg/utils"
)

type AuthController struct {
	authUsecase   usecase.AuthUsecase
	loggedInUsers map[uint]string
	secret        config.Secret
}

func NewAuthHandler(router *gin.RouterGroup, au usecase.AuthUsecase) {
	ac := AuthController{
		authUsecase:   au,
		loggedInUsers: make(map[uint]string),
	}
	router.POST("/login", ac.HandleLogin)
	router.POST("/logout", ac.Logout)
}

func (ac *AuthController) HandleLogin(ctx *gin.Context) {
	var loginRequest dto.ApiloginRequest
	if err := ctx.BindJSON(&loginRequest); err != nil {
		FailedJSONResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user, err := ac.authUsecase.LoginUser(loginRequest.Name, loginRequest.Password)
	if err != nil {
		switch err {
		case usecase.ErrUsecaseInvalidAuth:
			FailedJSONResponse(ctx, http.StatusUnauthorized, "Invalid credentials")
		default:
			FailedJSONResponse(ctx, http.StatusInternalServerError, err.Error())
		}
		return
	}

	expInSeconds := ac.secret.Exp.Seconds()
	token, err := utils.GenerateToken(ac.secret.Key, int64(user), time.Duration(expInSeconds)*time.Second)
	if err != nil {
		FailedJSONResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	responseData := gin.H{"token": token, "expired": ac.secret.Exp.String()}
	SuccessJSONResponse(ctx, http.StatusOK, "Login success", responseData)
}

func (ah *AuthController) Logout(ctx *gin.Context) {
	user := ctx.MustGet("user").(*model.User)
	userID := user.ID

	err := ah.authUsecase.LogoutUser(int64(userID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	delete(ah.loggedInUsers, uint(userID))

	ctx.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}
