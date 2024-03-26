package services

import (
	"errors"
	"fmt"

	internal "github.com/evergreenies/go-gin-tutorial/internal/model"
	"github.com/evergreenies/go-gin-tutorial/internal/utils"
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
	if err := a.db.Where("email = ?", email).Find(&user).Error; err != nil {
		return nil, err
	}

	if user.Email == "" {
		return nil, errors.New(fmt.Sprintf("no user found with email=%s", *email))
	}

	if !utils.CheckPasswordHash(*password, user.Password) {
		return nil, errors.New("incorrect password")
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

	if a.IsUserExist(email) {
		return nil, errors.New(fmt.Sprintf("user already exists with this email=%s", *email))
	}

	hashedPassword, err := utils.HashPassword(*password)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error hashing password, %v", err))
	}

	var user internal.User
	user.Email = *email
	user.Password = hashedPassword

	if err := a.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (a *AuthServices) IsUserExist(email *string) bool {
	var user *internal.User
	if err := a.db.Where("email = ?", email).Find(&user).Error; err != nil {
		return false
	}

	if user.Email == "" {
		return false
	}

	return true
}
