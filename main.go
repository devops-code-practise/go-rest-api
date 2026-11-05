package main

import (
	"net/http"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
	"example.com/rest-api/db"
)

func main() {
	db.InitDB()
	r := gin.Default()
	
	r.GET("/get-events", func(c *gin.Context) {
		events, err := models.GetAllEvents()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

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

		err = newEvent.Save()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "event": newEvent})
	})	

	r.Run(":8080")

}