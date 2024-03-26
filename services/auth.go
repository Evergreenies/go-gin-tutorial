package services

import (
	"errors"

	internal "github.com/evergreenies/go-gin-tutorial/internal/model"
	"gorm.io/gorm"
)

type AuthServices struct {
	db *gorm.DB
}

func InitAuthService(db *gorm.DB) *AuthServices {
	db.AutoMigrate(&internal.User{})
	return &AuthServices{
		db: db,
	}
}

func (a *AuthServices) Login(email *string, password *string) (*internal.User, error) {
	if email == nil {
		return nil, errors.New("email cant be null")
	}

	if password == nil {
		return nil, errors.New("you must provide password")
	}

	var user internal.User
	if err := a.db.Where("email = ?", email).Where("password = ?", password).Find(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (a *AuthServices) Register(email *string, password *string) (*internal.User, error) {
	if email == nil {
		return nil, errors.New("email cannot be null")
	}

	if password == nil {
		return nil, errors.New("you must have to provide password")
	}

	var user internal.User
	user.Email = *email
	user.Password = *password

	if err := a.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
