package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rolcho/go-rest-api/utils"
)

func Authenticate(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")

	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "You should login first"})
		return
	}

	userId, err := utils.ValidateToken(token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "You can't access"})
		return
	}

	ctx.Set("userId", userId)
	ctx.Next()
}
