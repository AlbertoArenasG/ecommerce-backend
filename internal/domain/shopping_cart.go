package domain

import "time"

type ShoppingCart struct {
	ID        uint       `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	IsDeleted bool       `json:"is_deleted"`
	Items     []CartItem `json:"items"`
}

type CartItem struct {
	Product   Product `json:"product"`
	Quantity  int     `json:"quantity"`
	IsDeleted bool    `json:"is_deleted"`
}
