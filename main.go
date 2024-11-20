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
	log.Printf("Connecting...")
	db.ConnectDatabase(config)

	r := gin.Default()

	r.POST("/auth", handlers.Auth)
	r.POST("/register", handlers.Register)
	r.POST("/forgot-password", handlers.ForgotPassword)

	r.GET("/boxes", handlers.GetBoxes)
	r.DELETE("/boxes/:id", handlers.DeleteBox)
	r.POST("/boxes", handlers.CreateBox)
	r.PUT("/boxes/edit", handlers.EditBox)

	port := config.AppPort

	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s", port)
	r.Run(":" + port)
}
