package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rolcho/go-rest-api/models"
)

func registerForEvent(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
	}

	if err = event.Register(userId); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register user for event"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Registered"})
}

func cancelRegistration(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	var event models.Event

	event.Id = eventId

	if err = event.CancelRegistration(userId); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel registration for event"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Registration cancelled"})
}
