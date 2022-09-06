package controllers

import (
	"fmt"
	"net/http"
	statusmodel "synergy/web-service-gin/common/http"
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
	r, _ := dbjet.InsertData(contactDb)

	c.IndentedJSON(http.StatusOK, models.JsonResponse{
		Status: http.StatusOK,
		Data:   r,
	})
}

func JetUpdateContactDAta(c *gin.Context) {
	var data models.Contacts
	if err := c.BindJSON(&data); err != nil {
		c.IndentedJSON(http.StatusOK, models.JsonResponse{
			Status:  statusmodel.InvalidParameter,
			Message: fmt.Sprintf("Invalid Parameter %v", err),
		})
	}
	contactDb := model.Contacts{ContactID: data.ContactID, CustomerID: &data.CustomerID, ContactName: data.ContactName, Phone: &data.Phone, Email: &data.Email}
	r, err := dbjet.UpdateData(contactDb)
	if err != nil {
		c.IndentedJSON(http.StatusOK, models.JsonResponse{
			Status:  statusmodel.ErrorInDatabase,
			Message: fmt.Sprintf("Can't save contact in database %s", err),
		})
	} else {
		c.IndentedJSON(http.StatusOK, models.JsonResponse{
			Status: http.StatusOK,
			Data:   r,
		})
	}

}
