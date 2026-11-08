package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"example.com/rest-api/utils"
)

func Authenticate(ctx *gin.Context)  {
	token := ctx.Request.Header.Get("Authorization")

	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization header"})
		return 
		}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return 
	}

	ctx.Set("userId", userId)

	ctx.Next()
}