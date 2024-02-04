package repository

import "gorm.io/gorm"

type ShoppingCartRepository struct {
	db *gorm.DB
}

func NewShoppingCartRepository(db *gorm.DB) *ShoppingCartRepository {
	return &ShoppingCartRepository{db}
}
