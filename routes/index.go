package routes

import "github.com/gin-gonic/gin"

func AddRoutes(superRoute *gin.RouterGroup) {
	albumRoutes(superRoute)
	userRoutes(superRoute)
}
