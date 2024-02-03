package main

import (
	"log"
	"os"

	"github.com/AlbertoArenasG/ecommerce-backend/internal/delivery"
	"github.com/AlbertoArenasG/ecommerce-backend/internal/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	repository.ConnectDB()
	defer repository.CloseDB()

	app := fiber.New()

	delivery.SetupRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	app.Listen(":" + port)
}
