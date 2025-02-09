package services

import (
	"errors"
	"go-template/internal/models"
)

// 可能的錯誤代碼
var (
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

// UserService 介面，定義使用者服務的方法
type UserService interface {
	CreateUser(user *models.User) error
	GetUserByID(id uint) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id uint) error
	Login(username, password string) (string, error)
}
