package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	AppPort    string
}

func GetConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	return &Config{
		DBHost:     os.Getenv("PGHOST"),
		DBPort:     os.Getenv("PGPORT"),
		DBUser:     os.Getenv("PGUSER"),
		DBPassword: os.Getenv("PGPASSWORD"),
		DBName:     os.Getenv("PGDATABASE"),
		AppPort:    os.Getenv("PORT"),
	}
}

func (c *Config) GetDSN() string {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	dns := os.Getenv("DATABASE_URL")
	log.Println(dns, "dns")
	if dns == "" {
		dns = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName)
	}

	return dns
}
