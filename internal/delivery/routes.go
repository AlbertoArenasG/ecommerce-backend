package delivery

import (
	"net/http"

	"github.com/AlbertoArenasG/ecommerce-backend/internal/repository"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/health", func(c *fiber.Ctx) error {
		if err := repository.DB.Exec("SELECT 1").Error; err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"message": "Connection error",
			})
		}
		return c.JSON(fiber.Map{
			"success": true,
			"message": "OK",
		})
	})
}
