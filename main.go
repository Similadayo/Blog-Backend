package main

import (
	"log"
	"os"

	"github.com/Similadayo/myBlog/database"
	"github.com/Similadayo/myBlog/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	database.Connect()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env files in main")
	}
	port := os.Getenv("PORT")
	app := fiber.New()
	routes.Setup(app)
	app.Listen(":" + port)
}
