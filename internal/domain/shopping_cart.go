package domain

import "time"

type ShoppingCart struct {
	ID        uint               `json:"id"`
	CreatedAt time.Time          `json:"created_at"`
	IsDeleted bool               `json:"is_deleted,omitempty"`
	Items     []ShoppingCartItem `json:"items,omitempty" gorm:"foreignKey:CartID"`
}

type ShoppingCartItem struct {
	CartID    uint `json:"cart_id"`
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}
