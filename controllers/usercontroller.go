package controllers

import (
	"fmt"
	"net/http"
	"synergy/web-service-gin/common"
	"synergy/web-service-gin/database"
	"synergy/web-service-gin/models"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	fmt.Println("# Call GetUser")
	userList := database.GetUsers()
	res := models.JsonResponse{
		Status: http.StatusOK,
		Data:   userList,
	}
	c.IndentedJSON(http.StatusOK, res)
}

func InsertUser(c *gin.Context) {
	var newuser models.User
	if err := c.BindJSON(&newuser); err != nil {
		common.CheckErrMessage("insertuser: ", err)
	}
	fmt.Println("# Call InsertUser BindJSON", newuser)
	// call user
	insertuser := database.InsertUser(newuser)
	res := models.JsonResponse{
		Status: http.StatusOK,
		Data:   insertuser,
	}
	c.IndentedJSON(http.StatusOK, res)
}

func UpdateUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		common.CheckErrMessage("updateuser: ", err)
	}
	fmt.Println("# Call UpdateUser BindJSON", user)
	var isUpdate bool = database.UpdateUser(user)
	var status int
	if isUpdate {
		status = http.StatusOK
	} else {
		status = http.StatusNotFound
	}
	res := models.JsonResponse{
		Status: int64(status),
		Data:   user,
	}
	c.IndentedJSON(http.StatusOK, res)
}
func DeleteUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		common.CheckErrMessage("deleteuser: ", err)
	}
	fmt.Println("# Call DeleteUser BindJSON", user)
	var isUpdate bool = database.DeleteUser(int(user.UserID))
	var status int
	if isUpdate {
		status = http.StatusOK
	} else {
		status = http.StatusNotFound
	}
	res := models.JsonResponse{
		Status:  int64(status),
		Message: "Delete Success",
	}
	c.IndentedJSON(http.StatusOK, res)
}
