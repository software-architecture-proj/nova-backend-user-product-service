package repos

import (
    "gorm.io/gorm"
    "github.com/software-architecture-proj/nova-backend-user-product-service/internal/models"
    "github.com/google/uuid"
)

type FavoriteRepository interface {
    CreateFavorite(favorite *Favorite) error
    GetFavoritesByUserID(user_id uuid.UUID) ([]Favorite, error)
    UpdateFavorite(favorite *Favorite) error
    DeleteFavoriteByID(id uuid.UUID) error
}

type favoriteRepo struct {
    db *gorm.DB
}

func NewFavoriteRepository(db *gorm.DB) FavoriteRepository {
	return &favoriteRepo{db}
}

func (r *favoriteRepo) CreateFavorite(favorite *Favorite) error {
	return r.db.Create(favorite).Error
}

func (r *favoriteRepo) GetFavoritesByUserID(user_id uuid.UUID) ([]Favorite, error) {
	var favorites []Favorite
	if err := r.db.Where("user_id = ?", user_id).Find(&favorites).Error; err != nil {
		return nil, err
	}
	return favorites, nil
}

func (r *favoriteRepo) UpdateFavorite(favorite *Favorite) error {
    return r.db.Model(&Favorite{}).Where("id = ?", favorite.ID).Updates(favorite).Error
}

func (r *favoriteRepo) DeleteFavoriteByID(id uuid.UUID) error {
    return r.db.Where("id = ?", id).Delete(&Favorite{}).Error
}
