package routes

import (
	"synergy/web-service-gin/controllers"

	"github.com/gin-gonic/gin"
)

func AddRoutes(superRoute *gin.RouterGroup) {
	// register auth
	superRoute.POST("auth", controllers.AuthHandler)
	// register group-route
	albumRoutes(superRoute)
	userRoutes(superRoute)
}
