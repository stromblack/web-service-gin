package controllers

import (
	"net/http"
	"synergy/web-service-gin/database/gorm"
	"synergy/web-service-gin/models"

	"github.com/gin-gonic/gin"
)

func GormGetAll(c *gin.Context) {
	// open gorm
	user := gorm.GetAllData()
	c.IndentedJSON(http.StatusOK, models.JsonResponse{
		Status: http.StatusOK,
		Data:   user,
	})
}
