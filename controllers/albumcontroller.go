package controllers

import (
	"log"
	"net/http"
	"strconv"
	models "synergy/web-service-gin/models"

	"github.com/gin-gonic/gin"
)

var albums = []models.Album{
	{ID: 1, Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: 2, Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responds with the list of all albums as JSON.
func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
func InsertAlbum(c *gin.Context) {
	var newAlbum models.Album
	// & goes in front of a variable when you want to get that variable's memory address.
	if err := c.BindJSON(&newAlbum); err != nil {
		log.Fatalln("error: invalid json", err)
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
func FindAlbum(c *gin.Context) {
	// convert id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatalln("error convert string to int")
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "server error",
		})
		return
	}
	// find
	for _, item := range albums {
		if item.ID == id {
			c.IndentedJSON(http.StatusOK, item)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
