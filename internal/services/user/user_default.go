package user

import (
	"time"

	"go-template/internal/models"
	"go-template/internal/repository"
	"go-template/internal/utils/jwt"
	"go-template/internal/utils/logger"

	"golang.org/x/crypto/bcrypt"
)

// ServiceDefault Struct，實作 UserService 介面
type ServiceDefault struct {
	userRepo   *repository.UserRepository
	jwtService *jwt.Service
}

// NewUserService 建立一個新的 user 實例
func NewUserService(userRepo *repository.UserRepository, jwtService *jwt.Service) Service {
	return &ServiceDefault{userRepo: userRepo, jwtService: jwtService}
}

// CreateUser 建立一個新的使用者
// @param user body models.User true "使用者資訊"
// @return error 錯誤訊息
func (svc *ServiceDefault) CreateUser(user *models.User) error {
	// 將密碼加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Logger.Errorf("Error hashing password: %v", err) // 記錄密碼加密錯誤
		return err
	}
	user.Password = string(hashedPassword)

	// 將使用者資料存入資料庫
	err = svc.userRepo.Create(user)
	if err != nil {
		logger.Logger.Errorf("Error creating user in repository: %v", err) // 記錄資料庫錯誤
		return err
	}

	logger.Logger.Infof("User created successfully: %s", user.Username) // 記錄使用者建立成功
	return nil
}

// GetUserByID 根據 ID 取得使用者資訊
// @param id path uint true "使用者 ID"
// @return user 使用者資訊
// @return error 錯誤訊息
func (svc *ServiceDefault) GetUserByID(id uint) (*models.User, error) {
	user, err := svc.userRepo.GetByID(id)
	if err != nil {
		logger.Logger.Debugf("Error getting user by ID: %v", err) // 記錄錯誤
		return nil, ErrUserNotFound
	}
	logger.Logger.Debugf("User found by ID: %d", id) // 記錄找到使用者
	return user, nil
}

// GetUserByUsername 根據使用者名稱取得使用者資訊
// @param username path string true "使用者名稱"
// @return user 使用者資訊
// @return error 錯誤訊息
func (svc *ServiceDefault) GetUserByUsername(username string) (*models.User, error) {
	user, err := svc.userRepo.GetByUsername(username)
	if err != nil {
		logger.Logger.Debugf("Error getting user by username: %v", err) // 記錄錯誤
		return nil, ErrUserNotFound
	}
	logger.Logger.Debugf("User found by username: %s", username) // 記錄找到使用者
	return user, nil
}

// UpdateUser 更新使用者資訊
// @param user body models.User true "使用者資訊"
// @return error 錯誤訊息
func (svc *ServiceDefault) UpdateUser(user *models.User) error {
	err := svc.userRepo.Update(user)
	if err != nil {
		logger.Logger.Errorf("Error updating user in repository: %v", err) // 記錄錯誤
		return err
	}
	logger.Logger.Debugf("User updated: %s", user.Username) // 記錄使用者已更新
	return nil
}

// DeleteUser 刪除使用者
// @param id path uint true "使用者 ID"
// @return error 錯誤訊息
func (svc *ServiceDefault) DeleteUser(id uint) error {
	err := svc.userRepo.Delete(id)
	if err != nil {
		logger.Logger.Errorf("Error deleting user in repository: %v", err) // 記錄錯誤
		return err
	}
	logger.Logger.Debugf("User deleted: %d", id) // 記錄使用者已刪除
	return nil
}

// Login 使用者登入
// @param username body string true "使用者名稱"
// @param password body string true "密碼"
// @return token JWT token
// @return error 錯誤訊息
func (svc *ServiceDefault) Login(username, password string) (string, error) {
	// 根據使用者名稱取得使用者資訊
	user, err := svc.userRepo.GetByUsername(username)
	if err != nil {
		logger.Logger.Debugf("Error getting user by username: %v", err) // 記錄錯誤
		return "", ErrUserNotFound
	}

	// 驗證密碼是否正確
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		logger.Logger.Debugf("Invalid credentials for user: %s", username) // 記錄錯誤
		return "", ErrInvalidCredentials
	}

	// 更新最後登入時間
	user.LastLogin = time.Now()
	updateErr := svc.userRepo.Update(user)
	if updateErr != nil {
		logger.Logger.Warnf("Error updating last login time: %v", updateErr) // 記錄錯誤
	}

	// 產生 JWT token
	token, err := svc.jwtService.GenerateToken(user.ID)
	if err != nil {
		logger.Logger.Errorf("Error generating token: %v", err) // 記錄錯誤
		return "", err
	}

	logger.Logger.Infof("User logged in: %s", username) // 記錄使用者登入
	return token, nil
}
