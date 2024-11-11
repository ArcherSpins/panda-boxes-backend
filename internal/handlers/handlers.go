package handlers

import (
	"errors"
	"log"
	"net/http"
	"panda-boxes/db"
	"panda-boxes/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetBoxes(c *gin.Context) {
	var items []models.Box
	if err := db.DB.Find(&items).Error; err != nil {
		log.Println(err, "error")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении данных"})
		return
	}
	c.JSON(http.StatusOK, items)
}

func CreateBox(c *gin.Context) {
	var box models.Box

	if err := c.ShouldBindBodyWithJSON(&box); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err := db.DB.Create(&box).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, box)
}

func EditBox(c *gin.Context) {
	var box models.Box
	var boxToEdit models.Box

	if err := c.ShouldBindBodyWithJSON(&box); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err := db.DB.Where("id = ?", box.ID).First(&boxToEdit).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("not found box with this id")})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	if box.Name != "" {
		boxToEdit.Name = box.Name
	}

	if box.Price != 0 {
		boxToEdit.Price = box.Price
	}

	err = db.DB.Save(&boxToEdit).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, boxToEdit)
}

func DeleteBox(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	err := db.DB.Where("id = ?", id).Delete(&models.Box{}).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("not found box with this id")})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, nil)
}
