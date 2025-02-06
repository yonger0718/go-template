package validators

import (
	"errors"
	"go-template/internal/models"
	"net/mail"
)

// ValidateUser 驗證使用者資料
// @param user *models.User 要驗證的使用者資料
// @return error 錯誤訊息
func ValidateUser(user *models.User) error {
	// 檢查使用者名稱長度
	if len(user.Username) < 4 {
		return errors.New("username must be at least 4 characters long")
	}
	// 檢查密碼長度
	if len(user.Password) < 6 {
		return errors.New("password must be at least 6 characters long")
	}
	// 檢查 email 格式
	if _, err := mail.ParseAddress(user.Email); err != nil {
		return errors.New("invalid email format")
	}
	return nil
}

// ValidateNewUser 驗證新增使用者資料
// @param user *models.User 要驗證的新增使用者資料
// @return error 錯誤訊息
func ValidateNewUser(username string, email string, password string) error {
	// 檢查使用者名稱長度
	if len(username) < 4 {
		return errors.New("username must be at least 4 characters long")
	}
	// 檢查密碼長度
	if len(password) < 6 {
		return errors.New("password must be at least 6 characters long")
	}
	// 檢查 email 格式
	if _, err := mail.ParseAddress(email); err != nil {
		return errors.New("invalid email format")
	}
	return nil
}
