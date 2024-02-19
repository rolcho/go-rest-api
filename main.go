package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/health", getHealth)

	server.Run(":8080")
}

func getHealth(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"health": "ok"})
}
