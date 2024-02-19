package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rolcho/go-rest-api/controllers"
	"github.com/rolcho/go-rest-api/db"
)

func main() {
	db.InitDB()

	server := gin.Default()

	eventsGroup := server.Group("/events")
	{
		eventsGroup.GET("", controllers.GetEvents)
		eventsGroup.GET("/:id", controllers.GetEventById)
		eventsGroup.POST("", controllers.CreateEvent)
	}

	server.GET("/health", controllers.GetHealth)
	server.Run(":8080")
}
