package routes

import (
	"synergy/web-service-gin/controllers"

	"github.com/gin-gonic/gin"
)

func gormRoutes(superroute *gin.RouterGroup) {
	gormRoutes := superroute.Group("/gorm")
	{
		gormRoutes.GET("", controllers.GormGetAll)
	}
}
