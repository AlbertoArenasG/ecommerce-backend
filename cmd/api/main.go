package main

import (
	"os"

	"github.com/AlbertoArenasG/ecommerce-backend/internal/delivery"
	"github.com/AlbertoArenasG/ecommerce-backend/internal/repository"
	"github.com/AlbertoArenasG/ecommerce-backend/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env file")
	}

	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)

	db := repository.ConnectDB()
	defer repository.CloseDB()

	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository, logger)

	app := fiber.New()

	delivery.SetupRoutes(app, productService, logger)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	app.Listen(":" + port)
}
