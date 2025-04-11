package entities

// User entity without direct dependency on products domain
type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" gorm:"size:100;not null"`
	Email    string `json:"email" gorm:"unique;size:100;not null"`
	Products []uint `json:"-" gorm:"-"` // Only used as a transient property, not persisted
}

// ProductAssociation represents a simplified product for user responses
type ProductAssociation struct {
	ID     uint    `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	UserID uint    `json:"user_id"`
}

// UserWithProducts is a data transfer object that includes product information
type UserWithProducts struct {
	ID       uint                 `json:"id"`
	Name     string               `json:"name"`
	Email    string               `json:"email"`
	Products []ProductAssociation `json:"products,omitempty"`
}

// ToUserWithProducts converts a User to a UserWithProducts DTO
func (u *User) ToUserWithProducts(products []ProductAssociation) *UserWithProducts {
	return &UserWithProducts{
		ID:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		Products: products,
	}
}
