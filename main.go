package main

import (
	"net/http"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	
	r.GET("/ping", func(c *gin.Context) {
		events := models.GetAllEvents()

		c.JSON(http.StatusOK, events)
	})

	r.POST("/post", func(c *gin.Context) {
		var newEvent models.Event
		err := c.ShouldBindJSON(&newEvent); 

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newEvent.ID = 1
		newEvent.UserID = 1

		newEvent.Save()

		c.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "event": newEvent})
	})	

	r.Run(":8080")

}