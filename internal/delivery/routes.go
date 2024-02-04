package delivery

import (
	"net/http"

	"github.com/AlbertoArenasG/ecommerce-backend/internal/delivery/response"
	"github.com/AlbertoArenasG/ecommerce-backend/internal/repository"
	"github.com/AlbertoArenasG/ecommerce-backend/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func SetupRoutes(app *fiber.App, productService *service.ProductService, logger *logrus.Logger) {
	productHandler := NewProductHandler(productService, logger)

	app.Get("/health", func(c *fiber.Ctx) error {
		if err := repository.DB.Exec("SELECT 1").Error; err != nil {
			logger.WithError(err).Error("Connection error")
			return c.Status(http.StatusInternalServerError).JSON(&response.ErrorResponse{
				Success: false,
				Message: "Connection error",
			})
		}
		return c.JSON(&response.SuccessResponse{
			Success: false,
			Message: "OK",
		})
	})

	// Products routes
	app.Post("/products", productHandler.CreateProduct)
	app.Put("/products/:id", productHandler.EditProduct)
	app.Delete("/products/:id", productHandler.DeleteProduct)
}
