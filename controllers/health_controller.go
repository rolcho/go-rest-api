package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHealth(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"health": "ok"})
}
