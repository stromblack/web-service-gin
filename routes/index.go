package routes

import (
	"synergy/web-service-gin/controllers"

	"github.com/gin-gonic/gin"
)

func AddRoutes(superRoute *gin.RouterGroup) {
	albumRoutes(superRoute)
	userRoutes(superRoute)
	superRoute.GET("auth", controllers.AuthHandler)
}
