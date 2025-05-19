package repos

import (
    "gorm.io/gorm"
    "github.com/software-architecture-proj/nova-backend-user-product-service/internal/models"
)

type PocketRepository interface {
    CreatePocket(pocket *Pocket) error
    GetPocketsByUserID(user_id string) ([]Pocket, error)
    UpdatePocketByID(id string, new_pocket *Pocket) error
    DeletePocketByID(id string) error
}

type pocketRepo struct {
    db *gorm.DB
}

func NewRepository(db *gorm.DB) PocketRepository {
	return &pocketRepo{db}
}

func (r *pocketRepo) CreatePocket(pocket *Pocket) error {
	return r.db.Create(pocket).Error
}

func (r *pocketRepo) GetPocketsByUserID(user_id string) ([]Pocket, error) {
	var pockets []Pocket
	if err := r.db.Where("User_ID = ?", user_id).Find(&pockets).Error; err != nil {
		return nil, err
	}
	return pockets, nil
}

func (r *pocketRepo) UpdatePocketByID(id string, new_pocket *Pocket) error {
    return r.db.Model(&Pocket{}).Where("ID = ?", id).Updates(new_pocket).Error
}

func (r *pocketRepo) DeletePocketByID(id string) error {
    return r.db.Where("ID = ?", id).Delete(&Pocket{}).Error
}
