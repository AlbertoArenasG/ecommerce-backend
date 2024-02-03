package repository

import (
	"github.com/AlbertoArenasG/ecommerce-backend/internal/domain"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db}
}

func (pr *ProductRepository) AddProduct(product *domain.Product) error {
	return pr.db.Create(product).Error
}
