package main

import (
	"backend/db"
	"backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := db.InitDB(); err != nil {
		panic("Failed to connect to database: " + err.Error())
	}
	db.CreateTable() // Ensure table creati
	server := gin.Default()

	routes.RegisterRoutes(server)
	server.Run(":8080") // localhost:8080
}
