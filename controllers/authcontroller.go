package controllers

import (
	"net/http"
	"strings"
	authorized "synergy/web-service-gin/common/auth"
	"synergy/web-service-gin/database"
	"synergy/web-service-gin/models"

	"github.com/gin-gonic/gin"
)

func AuthHandler(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, models.JsonResponse{
			Status:  2002,
			Message: "Invalid Parameter",
		})
		return
	}
	// Verify that the user name and password are correct
	userVerify := database.VerifyUser(user)
	if userVerify {
		tokenString, _ := authorized.GenToken(user)
		c.IndentedJSON(http.StatusOK, models.JsonResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    tokenString,
		})
		return
	}
	c.IndentedJSON(http.StatusUnauthorized, models.JsonResponse{
		Status:  http.StatusOK,
		Message: "Authentication failed",
	})
}

// JWT authmiddleware authentication middleware based on JWT
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// There are three ways for the client to carry a Token. 1 Put in request header 2 Put in the request body 3 Put in URI
		// Here, it is assumed that the Token is placed in the Authorization of the Header and starts with Bearer
		// The specific implementation method here should be determined according to your actual business situation
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.IndentedJSON(http.StatusOK, models.JsonResponse{
				Status:  2003,
				Message: "Request header auth Empty",
			})
			c.Abort()
			return
		}
		// Split by space
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.IndentedJSON(http.StatusOK, models.JsonResponse{
				Status:  2004,
				Message: "Request header auth Incorrect format",
			})
			c.Abort()
			return
		}
		// parts[1] is the obtained tokenString. We use the previously defined function to parse JWT to parse it
		mc, err := authorized.GetClaimsFromToken(parts[1])
		if err != nil {
			c.IndentedJSON(http.StatusUnauthorized, models.JsonResponse{
				Status:  2005,
				Message: "invalid Token",
			})
			c.Abort()
			return
		}
		// Save the currently requested username information to the requested context c
		c.Set("username", mc.UserInfo.UserName)
		c.Next() // Subsequent processing functions can use c.Get("username") to obtain the currently requested user information
	}
}
