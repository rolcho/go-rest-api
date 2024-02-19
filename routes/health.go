package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getHealth(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"health": "ok"})
}
