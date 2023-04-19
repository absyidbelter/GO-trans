package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"GO-Payment/config"
	"GO-Payment/internal/middleware"
	model "GO-Payment/internal/model/entity"
	"GO-Payment/internal/usecase"
)

type TransactionController struct {
	transactionUsecase usecase.TransactionUsecase
}

func NewTransactionController(router *gin.RouterGroup, tu usecase.TransactionUsecase, cfg *config.Secret) {
	tc := TransactionController{transactionUsecase: tu}
	authMiddleware := middleware.ValidateToken(cfg.Key)

	router.GET("/transactions", tc.GetTransactions)
	router.POST("/auth-mid/transfer", authMiddleware, tc.Transfer)
	router.POST("/transfer", tc.Transfer)
}

func (tc *TransactionController) GetTransactions(ctx *gin.Context) {
	res, err := tc.transactionUsecase.GetAllTransactions()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to get transactions",
		})
		return
	}

	if len(res) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "succes",
			"message": "No transactions found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   res,
	})
}

func (tc *TransactionController) Transfer(ctx *gin.Context) {
	var transferRequest model.TransferRequest
	err := ctx.ShouldBindJSON(&transferRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Failed to parse request",
		})
		return
	}

	destinationID := transferRequest.DestinationID
	destinationWallet := &model.Wallet{Number: destinationID}

	amount := int(transferRequest.Amount)
	transaction, err := tc.transactionUsecase.Transfer(int(transferRequest.UserID), destinationWallet, amount, transferRequest.History)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Transfer successful",
		"data":    transaction,
	})
}
