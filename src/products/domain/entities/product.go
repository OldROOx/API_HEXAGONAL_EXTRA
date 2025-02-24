package entities

type Product struct {
	ID     uint    `json:"id" gorm:"primaryKey"`
	Name   string  `json:"name" gorm:"size:100;not null"`
	Price  float64 `json:"price" gorm:"type:decimal(10,2);not null"`
	UserID uint    `json:"user_id" gorm:"not null"`
}
