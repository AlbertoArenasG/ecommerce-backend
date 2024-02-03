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
	if err := ps.validator.Struct(product); err != nil {
		ps.logger.WithError(err).Error("Validation failed", err)
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
