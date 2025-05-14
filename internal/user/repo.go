package user

import (
	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(user *User) error
	GetUserByID(id uint) (*User, error)
	ListUsers() ([]User, error)
}

type repo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repo{db}
}

func (r *repo) CreateUser(user *User) error {
	return r.db.Create(user).Error
}

func (r *repo) GetUserByID(id uint) (*User, error) {
	var user User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repo) ListUsers() ([]User, error) {
	var users []User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
