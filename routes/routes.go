package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/health", GetHealth)

	eventsGroup := server.Group("/events")
	{
		eventsGroup.GET("", getEvents)
		eventsGroup.GET("/:id", getEventById)
		eventsGroup.POST("", createEvent)
	}

}
