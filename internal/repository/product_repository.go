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

func (pr *ProductRepository) GetProductByID(id uint) (*domain.Product, error) {
	var product domain.Product
	err := pr.db.Where("id = ? AND is_deleted = ?", id, false).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (pr *ProductRepository) UpdateProduct(product *domain.Product) error {
	return pr.db.Save(product).Error
}
