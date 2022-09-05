package controllers

import (
	"fmt"
	"net/http"
	dbjet "synergy/web-service-gin/database/jet"
	"synergy/web-service-gin/models"
	"synergy/web-service-gin/postgres/public/model"

	"github.com/gin-gonic/gin"
)

func JetGetData(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, dbjet.GetData())
}

func JetAddContactData(c *gin.Context) {
	var data models.Contacts
	if err := c.BindJSON(&data); err != nil {
		c.IndentedJSON(http.StatusOK, models.JsonResponse{
			Status:  2002,
			Message: fmt.Sprintf("Invalid Parameter %s", err),
		})
		return
	}
	fmt.Printf("# JSON: %v \n", data)
	contactDb := model.Contacts{CustomerID: &data.CustomerID, ContactName: data.ContactName, Phone: &data.Phone, Email: &data.Email}
	// insert
	r := dbjet.InsertData(contactDb)

	c.IndentedJSON(http.StatusOK, models.JsonResponse{
		Status: http.StatusOK,
		Data:   r,
	})
}
