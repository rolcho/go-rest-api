package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rolcho/go-rest-api/models"
)

func main() {
	server := gin.Default()

	server.GET("/health", getHealth)
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getHealth(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"health": "ok"})
}

func getEvents(ctx *gin.Context) {
	events := models.GetAllEvents()
	ctx.JSON(http.StatusOK, events)
}

func createEvent(ctx *gin.Context) {
	var event models.Event

	if err := ctx.ShouldBindJSON(&event); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}

	event.Save()
	ctx.JSON(http.StatusCreated, gin.H{"message": "event created", "event": &event})
}
