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

type ProductHandler struct {
	productService *service.ProductService
	logger         *logrus.Logger
}

func NewProductHandler(ps *service.ProductService, logger *logrus.Logger) *ProductHandler {
	return &ProductHandler{
		productService: ps,
		logger:         logger,
	}
}

func (ph *ProductHandler) GetProductByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		ph.logger.WithError(err).Error("Invalid product ID")
		return c.Status(http.StatusBadRequest).JSON(&response.ErrorResponse{
			Success: false,
			Message: "Invalid product ID",
			Error:   err.Error(),
		})
	}

	successResponse, errorResponse := ph.productService.GetProductDetailByID(uint(id))
	if errorResponse != nil {
		statusCode := http.StatusInternalServerError
		if errorResponse.Message == "Product not found" {
			statusCode = http.StatusNotFound
		}
		return c.Status(statusCode).JSON(errorResponse)
	}

	return c.JSON(successResponse)
}

func (ph *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	var product domain.Product
	if err := c.BodyParser(&product); err != nil {
		ph.logger.WithError(err).Error("Invalid request body")
		return c.Status(http.StatusBadRequest).JSON(&response.ErrorResponse{
			Success: false,
			Message: "Invalid request body",
		})
	}

	successResponse, errorResponse := ph.productService.AddProduct(&product)

	if errorResponse != nil {
		statusCode := http.StatusInternalServerError
		if errorResponse.Message == "Validation failed" {
			statusCode = http.StatusBadRequest
		}
		return c.Status(statusCode).JSON(errorResponse)
	}

	return c.Status(http.StatusCreated).JSON(successResponse)
}

func (ph *ProductHandler) EditProduct(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		ph.logger.WithError(err).Error("Invalid product ID")
		return c.Status(http.StatusBadRequest).JSON(&response.ErrorResponse{
			Success: false,
			Message: "Invalid product ID",
			Error:   err.Error(),
		})
	}

	var updatedProduct domain.Product
	if err := c.BodyParser(&updatedProduct); err != nil {
		ph.logger.WithError(err).Error("Invalid request body")
		return c.Status(http.StatusBadRequest).JSON(&response.ErrorResponse{
			Success: false,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	successResponse, errorResponse := ph.productService.EditProduct(uint(id), &updatedProduct)

	if errorResponse != nil {
		statusCode := http.StatusInternalServerError
		if errorResponse.Message == "Product not found" {
			statusCode = http.StatusNotFound
		}
		return c.Status(statusCode).JSON(errorResponse)
	}

	return c.JSON(successResponse)
}

func (ph *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		ph.logger.WithError(err).Error("Invalid product ID")
		return c.Status(http.StatusBadRequest).JSON(&response.ErrorResponse{
			Success: false,
			Message: "Invalid product ID",
			Error:   err.Error(),
		})
	}

	successResponse, errorResponse := ph.productService.DeleteProductByID(uint(id))

	if errorResponse != nil {
		statusCode := http.StatusInternalServerError
		if errorResponse.Message == "Product not found" {
			statusCode = http.StatusNotFound
		}
		return c.Status(statusCode).JSON(errorResponse)
	}

	return c.JSON(successResponse)
}
