package repos

import (
    "gorm.io/gorm"
    "github.com/software-architecture-proj/nova-backend-user-product-service/internal/models"
)

type FavoriteRepository interface {
    CreateFavorite(favorite *Favorite) error
    GetFavoritesByUserID(user_id string) ([]Favorite, error)
    UpdateFavoriteByID(id string, new_favorite *Favorite) error
    DeleteFavoriteByID(id string) error
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

func (r *favoriteRepo) GetFavoritesByUserID(user_id string) ([]Favorite, error) {
	var favorites []Favorite
	if err := r.db.Where("User_ID = ?", user_id).Find(&favorites).Error; err != nil {
		return nil, err
	}
	return favorites, nil
}

func (r *favoriteRepo) UpdateFavoriteByID(id string, new_favorite *Favorite) error {
    return r.db.Model(&Favorite{}).Where("ID = ?", id).Updates(new_favorite).Error
}

func (r *favoriteRepo) DeleteFavoriteByID(id string) error {
    return r.db.Where("ID = ?", id).Delete(&Favorite{}).Error
}
