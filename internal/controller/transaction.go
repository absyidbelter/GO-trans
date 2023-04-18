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
	router.POST("/transfer.", tc.Transfer)
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

func (tc *TransactionController) Transfer(c *gin.Context) {
	// Mendapatkan userID dari context
	userID, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get userID from context"})
		return
	}

	var req model.TransferRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaction, err := tc.transactionUsecase.Transfer(userID.(int), req.DestinationID, float64(req.Amount), req.PaymentMethodType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transaction)
}
