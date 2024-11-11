package main

import (
	"log"
	"panda-boxes/configs"
	"panda-boxes/db"
	"panda-boxes/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Printf("Starting...")
	config := configs.GetConfig()
	db.ConnectDatabase()

	r := gin.Default()
	r.GET("/boxes", handlers.GetBoxes)
	r.POST("/boxes", handlers.CreateBox)
	r.PUT("/boxes/edit", handlers.EditBox)

	port := config.AppPort

	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s", port)
	r.Run(":" + port)
}
