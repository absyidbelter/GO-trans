package middleware

import (
	"GO-Payment/internal/model/dto"
	"GO-Payment/pkg/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type CustomClaims struct {
	UserID int64 `json:"userID"`
	jwt.StandardClaims
}

func ValidateToken(secret []byte) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHdr := ctx.GetHeader("Authorization")
		tokStr := strings.Replace(authHdr, "Bearer ", "", 1)

		userID, err := utils.ValidateToken(tokStr, secret)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.NewApiResponseFailed("invalid token"))
			return
		}

		ctx.Set("user_id", userID)
		ctx.Next()
	}
}
