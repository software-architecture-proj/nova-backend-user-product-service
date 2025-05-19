package repos

import (
	"gorm.io/gorm" 
    "github.com/software-architecture-proj/nova-backend-user-product-service/internal/models"
)

type UserRepository interface {
	CreateUser(user *User) error
	GetUserByID(id uint) (*User, error)
	ListUsers() ([]User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) UserRepository {
	return &userRepo{db}
}

func (r *userRepo) CreateUser(user *User) error {
	return r.db.Create(user).Error
}

func (r *userRepo) GetUserByID(id uint) (*User, error) {
	var user User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) ListUsers() ([]User, error) {
	var users []User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
