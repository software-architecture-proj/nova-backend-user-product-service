package favorite

import (
    "gorm.io/gorm"
)

type Repository interface {
    CreateFavorite(favorite *Favorite) error
    GetFavoritesByUserID(user_id string) ([]Favorite, error)
    UpdateFavoriteByID(id string, new_favorite *Favorite) error
    DeleteFavoriteByID(id string) error
}

type repo struct {
    db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repo{db}
}

func (r *repo) CreateFavorite(favorite *Favorite) error {
	return r.db.Create(favorite).Error
}

func (r *repo) GetFavoritesByUserID(user_id string) ([]Favorite, error) {
	var favorites []Favorite
	if err := r.db.Where("User_ID = ?", user_id).Find(&favorites).Error; err != nil {
		return nil, err
	}
	return favorites, nil
}

func (r *repo) UpdateFavoriteByID(id string, new_favorite *Favorite) error {
    return r.db.Model(&Favorite{}).Where("ID = ?", id).Updates(new_favorite).Error
}

func (r *repo) DeleteFavoriteByID(id string) error {
    return r.db.Where("ID = ?", id).Delete(&Favorite{}).Error
}
