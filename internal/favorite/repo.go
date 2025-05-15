package favorite

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	SaveFavorite(currentUser, otherUser uuid.UUID, alias string) (*Favorite, error)
	GetFavorites(currentUser uuid.UUID) ([]FavoriteInfo, error)
	DeleteFavorite(currentUser, otherUser uuid.UUID) error
}
type (
	FavoriteInfo struct {
		ID       string
		Username string
		Alias    string
	}
	repo struct {
		db *gorm.DB
	}
)

func NewRepository(db *gorm.DB) Repository {
	return &repo{db}
}

func (r *repo) SaveFavorite(currentUser, otherUser uuid.UUID, alias string) (*Favorite, error) {
	var fav Favorite
	if err := r.db.Where("user = ? AND favorite_user = ?", currentUser, otherUser).First(&fav).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			fav = Favorite{
				ID:             uuid.New(),
				UserID:         currentUser,
				FavoriteUserID: otherUser,
				Alias:          alias,
			}
			if err := r.db.Create(&fav).Error; err != nil {
				return nil, err
			}
			return &fav, nil
		}
		return nil, err
	}
	fav.Alias = alias
	if err := r.db.Save(&fav).Error; err != nil {
		return nil, err
	}
	return &fav, nil

}
func (r *repo) DeleteFavorite(currentUser, otherUser uuid.UUID) error {
	err := r.db.Where("user = ? AND favorite_user = ?", currentUser, otherUser).Delete(&Favorite{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) GetFavorites(currentUser uuid.UUID) ([]FavoriteInfo, error) {
	var favorites []Favorite
	err := r.db.
		Where("user = ?", currentUser).
		Preload("FavoriteUser").
		Find(&favorites).Error
	if err != nil {
		return nil, err
	}

	result := make([]FavoriteInfo, 0, len(favorites))
	for _, fav := range favorites {
		result = append(result, FavoriteInfo{
			ID:       fav.FavoriteUserID.String(),
			Username: fav.FavoriteUser.Username,
			Alias:    fav.Alias,
		})
	}

	return result, nil
}
