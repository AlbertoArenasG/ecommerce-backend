package service

import (
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
