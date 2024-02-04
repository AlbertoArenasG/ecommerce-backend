package delivery

import (
	"github.com/AlbertoArenasG/ecommerce-backend/internal/service"
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
