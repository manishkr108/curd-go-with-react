package routes

import (
	"backend/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(Server *gin.Engine) {
	Server.GET("/events/", getEvents)
	Server.GET("/events/:id", getEvent)

	authenticated := Server.Group("/")

	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	//SignUp route
	Server.POST("/signup", SignUp)
	Server.POST("/login", login)
}
