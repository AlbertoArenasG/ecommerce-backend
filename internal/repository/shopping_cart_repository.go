package repository

import (
	"github.com/AlbertoArenasG/ecommerce-backend/internal/domain"
	"gorm.io/gorm"
)

type ShoppingCartRepository struct {
	db *gorm.DB
}

func NewShoppingCartRepository(db *gorm.DB) *ShoppingCartRepository {
	return &ShoppingCartRepository{db}
}

func (r *ShoppingCartRepository) CreateShoppingCart() (*domain.ShoppingCart, error) {
	cart := &domain.ShoppingCart{}
	if err := r.db.Create(cart).Error; err != nil {
		return nil, err
	}
	return cart, nil
}

func (r *ShoppingCartRepository) AddProductToCart(item *domain.ShoppingCartItem) error {
	var existingItem domain.ShoppingCartItem
	result := r.db.Where("cart_id = ? AND product_id = ?", item.CartID, item.ProductID).First(&existingItem)
	if result.RowsAffected > 0 {
		existingItem.Quantity = item.Quantity
		return r.db.Where("cart_id = ? AND product_id = ?", item.CartID, item.ProductID).Save(&existingItem).Error
	}

	return r.db.Create(item).Error
}

func (r *ShoppingCartRepository) GetCartContents(cartID uint) (*domain.ShoppingCart, error) {
	var cart domain.ShoppingCart
	if err := r.db.Preload("Items.Product").First(&cart, cartID).Error; err != nil {
		return nil, err
	}

	filteredItems := make([]domain.ShoppingCartItem, 0)
	for _, item := range cart.Items {
		if !item.Product.IsDeleted {
			filteredItems = append(filteredItems, item)
		}
	}
	cart.Items = filteredItems

	return &cart, nil
}

func (r *ShoppingCartRepository) RemoveProductFromCart(cartID, productID uint) error {
	return r.db.Delete(&domain.ShoppingCartItem{}, "cart_id = ? AND product_id = ?", cartID, productID).Error
}

func (r *ShoppingCartRepository) CheckProductAndCartExistence(productID, cartID uint) (bool, bool, error) {
	var productCount, cartCount int64
	err := r.db.Model(&domain.Product{}).Where("id = ? AND is_deleted = ?", productID, false).Count(&productCount).Error
	if err != nil {
		return false, false, err
	}

	err = r.db.Model(&domain.ShoppingCart{}).Where("id = ? AND is_deleted = ?", cartID, false).Count(&cartCount).Error
	if err != nil {
		return false, false, err
	}

	return productCount > 0, cartCount > 0, nil
}
