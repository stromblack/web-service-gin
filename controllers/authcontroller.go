package controllers

import (
	"net/http"
	authorized "synergy/web-service-gin/common/auth"
	"synergy/web-service-gin/database"
	"synergy/web-service-gin/models"

	"github.com/gin-gonic/gin"
)

func AuthHandler(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusOK, models.JsonResponse{
			Status:  http.StatusNoContent,
			Message: "Invalid Parameter",
		})
		return
	}
	// Verify that the user name and password are correct
	userVerify := database.VerifyUser(user)
	if userVerify {
		tokenString, _ := authorized.GenToken(user.UserName)
		c.IndentedJSON(http.StatusOK, models.JsonResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    tokenString,
		})
		return
	}
	c.IndentedJSON(http.StatusOK, models.JsonResponse{
		Status:  http.StatusOK,
		Message: "Authentication failed",
	})
}
