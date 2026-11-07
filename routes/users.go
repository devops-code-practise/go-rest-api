package routes

import (
	"github.com/gin-gonic/gin"
	"example.com/rest-api/models"
)

func signup(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := user.Save(); err != nil {
		c.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}
	
	c.JSON(201, gin.H{"message": "User created successfully"})
}