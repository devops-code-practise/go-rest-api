package routes

import (
	"example.com/rest-api/middlewares"
	"github.com/gin-gonic/gin"
)



func RegisterRoutes(r *gin.Engine) {
	r.GET("/get-events", getEvents)
	r.GET("/events/:id", getEvent)

	authenticated := r.Group("/")
	authenticated.Use(middlewares.Authenticate)

	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerEvent)
	authenticated.DELETE("/events/:id/register", unregisterEvent)


	r.POST("/signup", signup)
	r.POST("/login", login)
}