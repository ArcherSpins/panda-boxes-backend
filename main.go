package main

import (
	"log"
	"panda-boxes/configs"
	"panda-boxes/db"
	"panda-boxes/internal/handlers"
	"panda-boxes/middleware"

	swaggerFiles "github.com/swaggo/files"

	_ "panda-boxes/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	log.Printf("Starting...")
	config := configs.GetConfig()
	log.Printf("Connecting...")
	db.ConnectDatabase(config)

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/auth/login", handlers.Auth)
	r.POST("/auth/register", handlers.Register)
	r.POST("/auth/password/forgot", handlers.ForgotPassword)

	private := r.Group("/api")
	private.Use(middleware.AuthRequired())

	private.GET("/boxes", handlers.GetBoxes)
	// r.GET("/boxes/:id", handlers.GetBoxByID)           // Получить коробку по ID
	private.POST("/boxes", handlers.CreateBox)
	private.PUT("/boxes", handlers.EditBox)
	private.DELETE("/boxes/:id", handlers.DeleteBox)

	port := config.AppPort

	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s", port)
	r.Run(":" + port)
}
