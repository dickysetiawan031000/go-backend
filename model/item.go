package model

type Item struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	Name        string  `json:"name" gorm:"not null"`
	Description string  `json:"description"`
	Price       float64 `json:"price" gorm:"not null"`
	Stock       int     `json:"stock" gorm:"not null"`
	CreatedAt   int64   `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   int64   `json:"updated_at" gorm:"autoUpdateTime"`
}
