package service

import (
	"github.com/AlbertoArenasG/ecommerce-backend/internal/delivery/response"
	"github.com/AlbertoArenasG/ecommerce-backend/internal/domain"
	"github.com/AlbertoArenasG/ecommerce-backend/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type ProductService struct {
	productRepository *repository.ProductRepository
	validator         *validator.Validate
	logger            *logrus.Logger
}

func NewProductService(pr *repository.ProductRepository, logger *logrus.Logger) *ProductService {
	return &ProductService{
		productRepository: pr,
		validator:         validator.New(),
		logger:            logger,
	}
}

func (ps *ProductService) AddProduct(product *domain.Product) (*response.SuccessResponse, *response.ErrorResponse) {
	if err := ps.validateProductData(product); err != nil {
		return nil, &response.ErrorResponse{
			Success: false,
			Message: "Validation failed",
			Error:   err.Error(),
		}
	}

	if err := ps.productRepository.AddProduct(product); err != nil {
		ps.logger.WithError(err).Error("Failed to create product", err)
		return nil, &response.ErrorResponse{
			Success: false,
			Message: "Failed to create product",
		}
	}

	ps.logger.Info("Product created successfully ", product)

	successResponse := &response.SuccessResponse{
		Success: true,
		Data:    product,
	}

	return successResponse, nil
}

func (ps *ProductService) EditProduct(id uint, updatedProduct *domain.Product) (*response.SuccessResponse, *response.ErrorResponse) {
	if err := ps.validateProductData(updatedProduct); err != nil {
		return nil, &response.ErrorResponse{
			Success: false,
			Message: "Validation failed",
			Error:   err.Error(),
		}
	}

	product, errorResponse := ps.getProductByID(id)
	if errorResponse != nil {
		return nil, errorResponse
	}

	product.Name = updatedProduct.Name
	product.Price = updatedProduct.Price
	product.ImageURL = updatedProduct.ImageURL
	product.Description = updatedProduct.Description

	if err := ps.productRepository.UpdateProduct(product); err != nil {
		ps.logger.WithError(err).Error("Failed to update product")
		return nil, &response.ErrorResponse{
			Success: false,
			Message: "Failed to update product",
		}
	}

	ps.logger.Info("Product updated successfully")

	successResponse := &response.SuccessResponse{
		Success: true,
		Data:    product,
	}

	return successResponse, nil
}

func (ps *ProductService) DeleteProductByID(id uint) (*response.SuccessResponse, *response.ErrorResponse) {
	product, errorResponse := ps.getProductByID(id)
	if errorResponse != nil {
		return nil, errorResponse
	}

	product.IsDeleted = true
	if err := ps.productRepository.UpdateProduct(product); err != nil {
		ps.logger.WithError(err).Error("Failed to delete product")
		return nil, &response.ErrorResponse{
			Success: false,
			Message: "Failed to delete product",
			Error:   err.Error(),
		}
	}

	ps.logger.Info("Product deleted successfully")

	successResponse := &response.SuccessResponse{
		Success: true,
		Message: "Product deleted successfully",
	}

	return successResponse, nil
}

func (ps *ProductService) getProductByID(id uint) (*domain.Product, *response.ErrorResponse) {
	existingProduct, err := ps.productRepository.GetProductByID(id)
	if err != nil {
		ps.logger.WithError(err).Error("Failed to get product")
		return nil, &response.ErrorResponse{
			Success: false,
			Message: "Product not found",
		}
	}
	return existingProduct, nil
}

func (ps *ProductService) validateProductData(data *domain.Product) error {
	if err := ps.validator.Struct(data); err != nil {
		ps.logger.WithError(err).Error("Validation failed", err)
		return err
	}
	return nil
}
