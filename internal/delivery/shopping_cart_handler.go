package delivery

import (
	"net/http"
	"strconv"

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

func (h *ShoppingCartHandler) GetCartContents(c *fiber.Ctx) error {
	cartID := c.Params("id")
	id, err := strconv.ParseUint(cartID, 10, 32)
	if err != nil {
		h.logger.WithError(err).Error("Invalid cart ID")
		return c.Status(http.StatusBadRequest).JSON(&response.ErrorResponse{
			Success: false,
			Message: "Invalid cart ID",
			Error:   err.Error(),
		})
	}

	successResponse, errorResponse := h.shoppingCartService.GetCartContents(uint(id))
	if errorResponse != nil {
		statusCode := http.StatusInternalServerError
		if errorResponse.Message == "Cart not found" {
			statusCode = http.StatusNotFound
		}
		return c.Status(statusCode).JSON(errorResponse)
	}

	return c.Status(http.StatusOK).JSON(successResponse)
}

func (h *ShoppingCartHandler) CreateCart(c *fiber.Ctx) error {
	successResponse, errorResponse := h.shoppingCartService.CreateCart()

	if errorResponse != nil {
		return c.Status(http.StatusInternalServerError).JSON(errorResponse)
	}

	return c.Status(http.StatusCreated).JSON(successResponse)
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
