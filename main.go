package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/health", getHealt)

	server.Run(":8080")
}

func getHealt(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"health": "ok"})
}
