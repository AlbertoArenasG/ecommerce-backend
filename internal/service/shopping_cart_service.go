package service

import (
	"github.com/AlbertoArenasG/ecommerce-backend/internal/delivery/response"
	"github.com/AlbertoArenasG/ecommerce-backend/internal/domain"
	"github.com/AlbertoArenasG/ecommerce-backend/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type ShoppingCartService struct {
	shoppingCartRepository *repository.ShoppingCartRepository
	validator              *validator.Validate
	logger                 *logrus.Logger
}

func NewShoppingCartService(scr *repository.ShoppingCartRepository, logger *logrus.Logger) *ShoppingCartService {
	return &ShoppingCartService{
		shoppingCartRepository: scr,
		validator:              validator.New(),
		logger:                 logger,
	}
}

func (s *ShoppingCartService) CreateCart() (*response.SuccessResponse, *response.ErrorResponse) {
	cart, err := s.shoppingCartRepository.CreateShoppingCart()

	if err != nil {
		s.logger.WithError(err).Error("Failed to create cart", err)
		return nil, &response.ErrorResponse{
			Success: false,
			Message: "Failed to create cart",
		}
	}

	s.logger.Info("Cart created successfully ", cart)

	successResponse := &response.SuccessResponse{
		Success: true,
		Data:    cart,
	}

	return successResponse, nil
}

func (s *ShoppingCartService) AddItemToCart(item *domain.ShoppingCartItem) (*response.SuccessResponse, *response.ErrorResponse) {
	if err := s.validator.Struct(item); err != nil {
		s.logger.WithError(err).Error("Validation failed for quantity")
		return nil, &response.ErrorResponse{
			Success: false,
			Message: "Validation failed",
			Error:   err.Error(),
		}
	}

	productExists, cartExists, err := s.shoppingCartRepository.CheckProductAndCartExistence(item.ProductID, item.CartID)
	if err != nil {
		s.logger.WithError(err).Error("Failed to add product to cart")
		return nil, &response.ErrorResponse{
			Success: false,
			Message: "Failed to add product to cart",
		}
	}

	if !productExists {
		return nil, &response.ErrorResponse{
			Success: false,
			Message: "Product not found",
		}
	}

	if !cartExists {
		return nil, &response.ErrorResponse{
			Success: false,
			Message: "Shopping cart not found",
		}
	}

	if err := s.shoppingCartRepository.AddProductToCart(item); err != nil {
		s.logger.WithError(err).Error("Failed to add product to cart")
		return nil, &response.ErrorResponse{
			Success: false,
			Message: "Failed to add product to cart",
		}
	}

	s.logger.Info("Product added to cart successfully")

	successResponse := &response.SuccessResponse{
		Success: true,
	}

	return successResponse, nil
}
