package repos

import (
    "gorm.io/gorm"
    "github.com/software-architecture-proj/nova-backend-user-product-service/internal/models"
    "github.com/google/uuid"
)

type PocketRepository interface {
    CreatePocket(pocket *Pocket) error
    GetPocketsByUserID(user_id uuid.UUID) ([]Pocket, error)
    UpdatePocket(pocket *Pocket) error
    DeletePocketByID(id uuid.UUID) error
}

type pocketRepo struct {
    db *gorm.DB
}

func NewPocketRepository(db *gorm.DB) PocketRepository {
	return &pocketRepo{db}
}

func (r *pocketRepo) CreatePocket(pocket *Pocket) error {
	return r.db.Create(pocket).Error
}

func (r *pocketRepo) GetPocketsByUserID(user_id uuid.UUID) ([]Pocket, error) {
	var pockets []Pocket
	if err := r.db.Where("user_id = ?", user_id).Find(&pockets).Error; err != nil {
		return nil, err
	}
	return pockets, nil
}

func (r *pocketRepo) UpdatePocket(pocket *Pocket) error {
    return r.db.Model(&Pocket{}).Where("id = ?", Pocket.ID).Updates(pocket).Error
}

func (r *pocketRepo) DeletePocketByID(id uuid.UUID) error {
    return r.db.Where("id = ?", id).Delete(&Pocket{}).Error
}
