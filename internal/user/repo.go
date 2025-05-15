package user

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(email, username, code_id string, phone int64, firstName, lastName string, birthdate time.Time) (*User, error)
	UpdateUser(current_user, firstName, lastName string) (*User, error)
	DeleteUser(user *User) error

	GetUserByUsername(username string) (*User, error)
	GetUsernameByID(id string) (string, error)

	UsernameIsValid(username string) error
	PhoneIsValid(phone int64) error
}
type repo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repo{db}
}

func (r *repo) CreateUser(email, username, code_id string, phone int64, firstName, lastName string, birthdate time.Time) (*User, error) {
	var g errgroup.Group
	ctx := context.Background()

	codeUUID, err := uuid.Parse(code_id)
	if err != nil {
		return nil, fmt.Errorf("invalid phone code_id: %w", err)
	}
	var user User = User{
		ID:        uuid.New(),
		Email:     email,
		Username:  username,
		CodeID:    codeUUID,
		Phone:     phone,
		FirstName: firstName,
		LastName:  lastName,
		Birthdate: birthdate,
	}

	g.Go(func() error {
		err := r.UsernameIsValid(user.Username)
		return err
	})
	if user.Phone != 0 {
		g.Go(func() error {
			err := r.PhoneIsValid(user.Phone)
			return err
		})
	}

	// Wait for all validations to complete
	if err := g.Wait(); err != nil {
		return nil, err // return the first error encountered
	}

	// All validations passed, now create the user
	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return &user, nil
}
func (r *repo) UpdateUser(current_user, firstName, lastName string) (*User, error) {
	var existingUser User
	userID, err := uuid.Parse(current_user)
	if err != nil {
		return nil, err
	}
	if err := r.db.First(&existingUser, "id = ?", userID).Error; err != nil {
		return nil, err
	}

	if firstName != "" {
		existingUser.FirstName = firstName
	}
	if lastName != "" {
		existingUser.LastName = lastName
	}

	if err := r.db.Updates(&existingUser).Error; err != nil {
		return nil, err
	}
	return &existingUser, nil
}
func (r *repo) DeleteUser(user *User) error {
	return r.db.Delete(user).Error
}
func (r *repo) GetUsernameByID(id string) (string, error) {
	var user User
	if err := r.db.Select("username").First(&user, "id = ?", id).Error; err != nil {
		return fmt.Sprintf("No User was found with id: '%s'", id), err
	}
	return user.Username, nil
}
func (r *repo) GetUserByUsername(username string) (*User, error) {
	var user User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *repo) UsernameIsValid(username string) error {
	var user User
	if err := r.db.Where("username = ?", username).First(&user).Error; err == nil {
		return fmt.Errorf("username '%s' is already in use", username)
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	} else {
		return err
	}
}
func (r *repo) PhoneIsValid(phone int64) error {
	var user User
	if err := r.db.Where("phone = ?", phone).First(&user).Error; err == nil {
		return fmt.Errorf("phone '%d' is already in use", phone)
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	} else {
		return err
	}
}
