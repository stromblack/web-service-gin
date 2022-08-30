package routes

import (
	"synergy/web-service-gin/controllers"

	"github.com/gin-gonic/gin"
)

func userRoutes(superRoute *gin.RouterGroup) {
	userRoute := superRoute.Group("user")
	{
		userRoute.GET("", controllers.GetUser)
		userRoute.POST("", controllers.InsertUser)
		userRoute.PUT("", controllers.UpdateUser)
		userRoute.DELETE("", controllers.DeleteUser)
	}
}
