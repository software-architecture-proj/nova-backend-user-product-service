package pocket

import (
    "gorm.io/gorm"
)

type Repository interface {
    CreatePocket(pocket *Pocket) error
    GetPocketsByUserID(user_id string) ([]Pocket, error)
    UpdatePocketByID(id string, new_pocket *Pocket) error
    DeletePocketByID(id string) error
}

type repo struct {
    db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repo{db}
}

func (r *repo) CreatePocket(pocket *Pocket) error {
	return r.db.Create(pocket).Error
}

func (r *repo) GetPocketsByUserID(user_id string) ([]Pocket, error) {
	var pockets []Pocket
	if err := r.db.Where("User_ID = ?", user_id).Find(&pockets).Error; err != nil {
		return nil, err
	}
	return pockets, nil
}

func (r *repo) UpdatePocketByID(id string, new_pocket *Pocket) error {
    return r.db.Model(&Pocket{}).Where("ID = ?", id).Updates(new_pocket).Error
}

func (r *repo) DeletePocketByID(id string) error {
    return r.db.Where("ID = ?", id).Delete(&Pocket{}).Error
}
