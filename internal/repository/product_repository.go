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

func (pr *ProductRepository) GetProducts(sortField, sortOrder, searchQuery string, page, limit int) ([]domain.Product, int, error) {
	var products []domain.Product
	var totalProducts int64

	query := pr.db.Model(&domain.Product{}).Where("is_deleted = ?", false)

	if searchQuery != "" {
		query = query.Where("name LIKE ?", "%"+searchQuery+"%")
	}

	err := query.Count(&totalProducts).Error
	if err != nil {
		return nil, 0, err
	}

	if sortField != "" {
		query = query.Order(sortField + " " + sortOrder)
	}

	offset := (page - 1) * limit
	err = query.Offset(offset).Limit(limit).Find(&products).Error
	if err != nil {
		return nil, 0, err
	}

	return products, int(totalProducts), nil
}
