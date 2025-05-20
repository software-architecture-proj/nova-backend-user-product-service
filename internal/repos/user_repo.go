package repos

import (
	"gorm.io/gorm" 
    "github.com/software-architecture-proj/nova-backend-user-product-service/internal/models"
    "github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(user *User) error
    GetUserByID(id uuid.UUID) (*User, error)
	UpdateUser(user *User) error
	DeleteUserById(id uuid.UUID) error
    ListUsers() ([]User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{db}
}

func (r *userRepo) CreateUser(user *User) error {
	return r.db.Create(user).Error
}

func (r *userRepo) UpdateUser(user *User) error {
    return r.db.Model(&User{}).Where("id = ?", user.ID).Updates(user).Error
}

func (r *userRepo) DeleteUserById(id uuid.UUID) error {
    return r.db.Delete(&models.User{}, "id = ?", id).Error
}

func (r *userRepo) GetUserByID(id uuid.UUID) (*User, error) {
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
