package handlers

import (
	"net/http"
	"panda-boxes/db"
	"panda-boxes/models"
	"panda-boxes/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func ForgotPassword(c *gin.Context) {
	var req struct {
		Email string `json:"email"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var user models.User

	if err := db.DB.Where("Email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User with this email does not exist"})
		return
	}

	token, err := utils.GenerateJWT(user, time.Now().Add(15*time.Minute).Unix())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	err = utils.SendEmail(req.Email, "Password Reset", "Click here to reset your password: https://example.com/reset-password?token="+token)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating reset password link", "info": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password reset link sent"})
}

// @Summary Post users
// @Description Create a user
// @Tags users
// @Param body body models.UserRegister true "User register"
// @Success 200 {object} map[string]interface{}
// @Router /auth/register [post]
func Register(c *gin.Context) {
	var req models.UserRegister

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}

	user := models.User{Username: req.Username, Email: req.Email, Password: hashedPassword}

	if err := db.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Already exist"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created"})
}

// @Summary Auth user
// @Description Login
// @Tags users
// @Success 200 {object} map[string]interface{}
// @Param body body models.UserAuth true "User auth"
// @Router /auth/login [post]
func Auth(c *gin.Context) {
	var req models.UserAuth

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var user models.User

	if err := db.DB.Where("username = ? OR email = ?", req.Username, req.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username/email or password", "info": err})
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}

	if !utils.CheckPasswordHash(req.Password, hashedPassword) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username/email or password"})
		return
	}

	token, err := utils.GenerateJWT(user, 0)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error with genereting token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user, "token": token})
}
