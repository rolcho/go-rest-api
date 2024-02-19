package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/health", getHealth)

	eventsGroup := server.Group("/events")
	{
		eventsGroup.GET("", getEvents)
		eventsGroup.GET("/:id", getEvent)
		eventsGroup.POST("", createEvent)
		eventsGroup.PUT("/:id", updateEvent)
		eventsGroup.DELETE("/:id", deleteEvent)
	}

	usersGroup := server.Group("/users")
	{
		usersGroup.GET("", getUsers)
		usersGroup.GET("/:id", getUser)
		usersGroup.POST("", createUser)
		usersGroup.PUT("/:id", updateUser)
		usersGroup.DELETE("/:id", deleteUser)
	}
}
