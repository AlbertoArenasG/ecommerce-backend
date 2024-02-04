package domain

type Product struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	Name        string  `json:"name" validate:"required"`
	Price       float64 `json:"price" validate:"required,min=0"`
	ImageURL    string  `json:"image_url"`
	Description string  `json:"description"`
	IsDeleted   bool    `json:"is_deleted,omitempty" gorm:"default:false"`
}
