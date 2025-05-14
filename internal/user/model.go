package user

import (
	"nova-backend-user-product-service/internal/countrycode"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID               `gorm:"type:char(36);primaryKey"`
	Email    string                  `gorm:"unique;not null"`
	Username string                  `gorm:"unique;not null"`
	Phone    *string                 `gorm:"unique"`
	CodeID   uuid.UUID               `gorm:"type:char(36);not null"` // FK field
	Code     countrycode.CountryCode `gorm:"foreignKey:CodeID;references:ID"`

	FirstName string    `gorm:"not null"`
	LastName  string    `gorm:"not null"`
	Birthdate time.Time `gorm:"not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}
