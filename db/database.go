package db

import (
	"log"
	"panda-boxes/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(config *configs.Config) {
	dsn := config.GetDSN()
	log.Println("Соединение с: ", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}
	DB = db
	log.Println("Соединение с базой данных установлено")
}

// package db

// import (
// 	"log"
// 	"os"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// var DB *gorm.DB

// func ConnectDatabase() {
// 	dsn := os.Getenv("DATABASE_URL")
// 	if dsn == "" {
// 		log.Fatal("DATABASE_URL не установлена")
// 	}

// 	log.Println(dsn)
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
// 	}
// 	DB = db
// 	log.Println("Соединение с базой данных установлено")
// }
