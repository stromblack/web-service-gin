package routes

import (
	"synergy/web-service-gin/controllers"

	"github.com/gin-gonic/gin"
)

func jetRoutes(superroute *gin.RouterGroup) {
	jet := superroute.Group("/jet")
	{
		jet.GET("", controllers.JetGetData)
		jet.POST("/contact", controllers.JetAddContactData)
		jet.PUT("/contact", controllers.JetUpdateContactDAta)
	}
}
