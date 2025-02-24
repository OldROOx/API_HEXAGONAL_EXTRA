package entities

import "API_HEXAGONAL_RECU/src/products/domain/entities"

type User struct {
	ID       uint               `json:"id" gorm:"primaryKey"`
	Name     string             `json:"name" gorm:"size:100;not null"`
	Email    string             `json:"email" gorm:"unique;size:100;not null"`
	Products []entities.Product `json:"products,omitempty" gorm:"foreignKey:UserID"`
}
