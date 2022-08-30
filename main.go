package main

import (
	"synergy/web-service-gin/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.New()
	// create group route
	route := app.Group("/api/v1")
	// add child into group route
	routes.AddRoutes(route)
	app.Run("localhost:8082")
}
