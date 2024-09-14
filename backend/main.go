package main

import (
	"backend/db"
	"backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	if err := db.InitDB(); err != nil {
		panic("Failed to connect to database: " + err.Error())
	}
	db.CreateTable() // Ensure table creati
	server := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}

	server.Use(cors.New(corsConfig))

	routes.RegisterRoutes(server)
	server.Run(":8080") // localhost:8080
}
