package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rolcho/go-rest-api/models"
)

func getEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, events)
}

func getEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch event id"})
		return
	}
	event, err := models.GetEventById(eventId)
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

func createEvent(ctx *gin.Context) {
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

func updateEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch event id"})
		return
	}
	_, err = models.GetEventById(eventId)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Could not fetch event id"})
		return
	}

	var updatedEvent models.Event

	if err = ctx.ShouldBindJSON(&updatedEvent); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}

	updatedEvent.Id = eventId
	if err = updatedEvent.Update(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Event updated successfully"})

}

func deleteEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch event id"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Could not fetch event id"})
		return
	}

	if err := event.Delete(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Event deleted"})
}
