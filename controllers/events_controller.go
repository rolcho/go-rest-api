package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rolcho/go-rest-api/models"
)

func GetEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, events)
}

func GetEventById(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not find event id"})
		return
	}
	event, err := models.GetEventById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Could not fetch event id"})
		return
	}

	if event == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		return
	}

	ctx.JSON(http.StatusOK, event)
}

func CreateEvent(ctx *gin.Context) {
	var event models.Event

	if err := ctx.ShouldBindJSON(&event); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}

	if err := event.Save(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Error writing the database"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "event created", "event": &event})
}
