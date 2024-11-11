package main

import (
	"panda-boxes/configs"
	"panda-boxes/db"
	"panda-boxes/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	config := configs.GetConfig()
	db.ConnectDatabase(config)

	r := gin.Default()
	r.GET("/boxes", handlers.GetBoxes)
	r.POST("/boxes", handlers.CreateBox)
	r.PUT("/boxes/edit", handlers.EditBox)

	r.Run(":8080")
}
