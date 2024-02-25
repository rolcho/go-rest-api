package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rolcho/go-rest-api/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/health", getHealth)

	eventsGroup := server.Group("/events")
	{
		eventsGroup.GET("/", getEvents)
		eventsGroup.GET("/:id", getEvent)
	}

	protectedEventsGroup := server.Group("/events").Use(middlewares.Authenticate)
	{
		protectedEventsGroup.POST("", createEvent)
		protectedEventsGroup.PUT("/:id", updateEvent)
		protectedEventsGroup.DELETE("/:id", deleteEvent)
	}

	usersGroup := server.Group("/users")
	{
		usersGroup.POST("/signup", signupUser)
		usersGroup.POST("/signin", signinUser)
	}

	protectedUsersGroup := server.Group("/users").Use(middlewares.Authenticate)
	{
		protectedUsersGroup.GET("", getUsers)
		protectedUsersGroup.GET("/:id", getUser)
		protectedUsersGroup.PUT("/:id", updateUser)
		protectedUsersGroup.DELETE("/:id", deleteUser)
	}
}
