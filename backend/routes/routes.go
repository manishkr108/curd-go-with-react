package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(Server *gin.Engine) {
	Server.GET("/events", getEvents)
	Server.GET("/events/:id", getEvent)
	Server.POST("/events", createEvent)
	Server.PUT("/events/:id", updateEvent)
}
