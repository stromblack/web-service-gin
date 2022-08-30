package routes

import (
	"synergy/web-service-gin/controllers"

	"github.com/gin-gonic/gin"
)

func albumRoutes(superRoute *gin.RouterGroup) {
	albumRouter := superRoute.Group("/album")
	{
		albumRouter.GET("/album", controllers.GetAlbums)
		albumRouter.POST("/album", controllers.InsertAlbum)
		albumRouter.GET("/album/:id", controllers.FindAlbum)
	}
}
