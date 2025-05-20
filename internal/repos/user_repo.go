package repos

import (
	"gorm.io/gorm" 
    "github.com/software-architecture-proj/nova-backend-user-product-service/internal/models"
    "github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(user *models.User) error
    GetUserById(id uuid.UUID) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUserById(id uuid.UUID) error
    ListUsers() ([]models.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{db}
}

func (r *userRepo) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepo) UpdateUser(user *models.User) error {
    return r.db.Model(&models.User{}).Where("id = ?", user.ID).Updates(user).Error
}

func (r *userRepo) DeleteUserById(id uuid.UUID) error {
    return r.db.Delete(&models.models.User{}, "id = ?", id).Error
}

func (r *userRepo) GetUserById(id uuid.UUID) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) ListUsers() ([]models.User, error) {
	var users []models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepo) ListCountryCodes() ([]models.CountryCode, error) {
	var codes []models.CountryCode
	if err := r.db.Find(&codes).Error; err != nil {
		return nil, err
	}
	return codes, nil
}

