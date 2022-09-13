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
	// create channel
	ch := make(chan models.ChannelResponse)
	go dbjet.GetData(ch)
	// wait channel complete
	var result models.ChannelResponse
	for res := range ch {
		result = res
	}
	c.IndentedJSON(http.StatusOK, result)
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
	// create channel
	ch := make(chan models.ChannelResponse)
	// insert
	go dbjet.InsertData(contactDb, ch)
	// wait channal
	var result models.ChannelResponse
	for res := range ch {
		result = res
	}
	c.IndentedJSON(http.StatusOK, models.JsonResponse{
		Status:  http.StatusOK,
		Data:    result.Result,
		Message: result.Error.Error(),
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
	// add channel & go
	ch := make(chan models.ChannelResponse)
	go dbjet.UpdateData(contactDb, ch)
	// wait channel
	for res := range ch {
		if res.Error != nil {
			c.IndentedJSON(http.StatusOK, models.JsonResponse{
				Status:  statusmodel.ErrorInDatabase,
				Message: fmt.Sprintf("Can't save contact in database %s", res.Error),
			})
		} else {
			c.IndentedJSON(http.StatusOK, models.JsonResponse{
				Status: http.StatusOK,
				Data:   res.Result,
			})
		}
	}
}
