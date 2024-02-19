package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rolcho/go-rest-api/db"
	"github.com/rolcho/go-rest-api/routes"
)

func main() {
	db.InitDB()

	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run(":3000")
}
