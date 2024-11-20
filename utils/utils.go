package utils

import (
	"log"
	"net/smtp"
	"panda-boxes/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("your_secret_key")

func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func SendEmail(to, subject, body string) error {
	from := "sendemailtest222@gmail.com"
	password := "imuxndrhziyyysdp"

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	auth := smtp.PlainAuth("", from, password, smtpHost)
	log.Println(from, to, msg)
	return smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(msg))
}

func GenerateJWT(user models.User, exp int64) (string, error) {
	if exp == 0 {
		exp = time.Now().Add(time.Hour * 72).Unix()
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     exp,
	})

	return token.SignedString(jwtSecret)
}
