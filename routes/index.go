package routes

import (
	"synergy/web-service-gin/controllers"

	"github.com/gin-gonic/gin"
)

func AddRoutes(superRoute *gin.RouterGroup) {
	// register auth
	superRoute.POST("auth", controllers.AuthHandler)
	superRoute.POST("refresh", controllers.RefreshTokenHandler)
	// register group-route
	superRoute.Use(controllers.JWTAuthMiddleware())
	albumRoutes(superRoute)
	userRoutes(superRoute)
	gormRoutes(superRoute)
	jetRoutes(superRoute)
}
