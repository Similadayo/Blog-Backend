package database

import (
	"log"
	"os"

	"github.com/Similadayo/myBlog/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .envfile in database")
	}
	dsn := os.Getenv("DSN")
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error connecting to database")
	} else {
		log.Println("Connected Successfully")
	}
	DB = database
	database.AutoMigrate(
		&models.User{},
		&models.Posts{},
	)
}
