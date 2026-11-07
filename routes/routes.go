package routes

import (
	"github.com/gin-gonic/gin"
)



func RegisterRoutes(r *gin.Engine) {
	r.GET("/get-events", GetEvents)
	r.GET("/events/:id", GetEvent)
	r.POST("/post", CreateEvent)
	r.PUT("/events/:id", UpdateEvent)
}