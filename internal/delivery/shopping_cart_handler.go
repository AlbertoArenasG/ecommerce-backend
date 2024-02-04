package delivery

import (
	"net/http"

	"github.com/AlbertoArenasG/ecommerce-backend/internal/delivery/response"
	"github.com/AlbertoArenasG/ecommerce-backend/internal/domain"
	"github.com/AlbertoArenasG/ecommerce-backend/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ShoppingCartHandler struct {
	shoppingCartService *service.ShoppingCartService
	logger              *logrus.Logger
}

func NewShoppingCartHandler(scs *service.ShoppingCartService, logger *logrus.Logger) *ShoppingCartHandler {
	return &ShoppingCartHandler{
		shoppingCartService: scs,
		logger:              logger,
	}
}

func (h *ShoppingCartHandler) AddItemToCart(c *fiber.Ctx) error {
	var item domain.ShoppingCartItem
	if err := c.BodyParser(&item); err != nil {
		h.logger.WithError(err).Error("Invalid request body")
		return c.Status(http.StatusBadRequest).JSON(&response.ErrorResponse{
			Success: false,
			Message: "Invalid request body",
		})
	}

	successResponse, errorResponse := h.shoppingCartService.AddItemToCart(&item)

	if errorResponse != nil {
		statusCode := http.StatusInternalServerError
		if errorResponse.Message == "Validation failed" {
			statusCode = http.StatusBadRequest
		}
		if errorResponse.Message == "Product not found" || errorResponse.Message == "Shopping cart not found" {
			statusCode = http.StatusNotFound
		}
		return c.Status(statusCode).JSON(errorResponse)
	}

	return c.Status(http.StatusOK).JSON(successResponse)
}
