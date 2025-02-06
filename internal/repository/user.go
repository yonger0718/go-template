package repository

import (
	"go-template/internal/models"
	"go-template/internal/utils/logutil"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository 建立一個新的 UserRepository 實例
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Create 新增一個使用者
// @param user 新增的使用者
// @return error 錯誤訊息
func (r *UserRepository) Create(user *models.User) error {
	result := r.db.Create(user)
	if result.Error != nil {
		logutil.Logger.Errorf("Error creating user in database: %v", result.Error) // 記錄資料庫錯誤
		return result.Error
	}
	logutil.Logger.Debugf("User created in database: %s", user.Username) // 記錄使用者已建立
	return nil
}

// GetByID 根據 ID 取得使用者
// @param id 使用者 ID
// @return user 使用者
// @return error 錯誤訊息
func (r *UserRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	result := r.db.First(&user, id)
	if result.Error != nil {
		logutil.Logger.Errorf("Error getting user by ID from database: %v", result.Error) // 記錄資料庫錯誤
		return nil, result.Error
	}
	logutil.Logger.Debugf("User found by ID in database: %d", id) // 記錄使用者已找到
	return &user, nil
}

// GetByUsername 根據使用者名稱取得使用者
// @param username 使用者名稱
// @return user 使用者
// @return error 錯誤訊息
func (r *UserRepository) GetByUsername(username string) (*models.User, error) {
	var user models.User
	result := r.db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		logutil.Logger.Errorf("Error getting user by username from database: %v", result.Error) // 記錄資料庫錯誤
		return nil, result.Error
	}
	logutil.Logger.Debugf("User found by username in database: %s", username) // 記錄使用者已找到
	return &user, nil
}

// Update 更新使用者資訊
// @param user 要更新的使用者
// @return error 錯誤訊息
func (r *UserRepository) Update(user *models.User) error {
	result := r.db.Save(user)
	if result.Error != nil {
		logutil.Logger.Errorf("Error updating user in database: %v", result.Error) // 記錄資料庫錯誤
		return result.Error
	}
	logutil.Logger.Debugf("User updated in database: %s", user.Username) // 記錄使用者已更新
	return nil
}

// Delete 根據 ID 刪除使用者
// @param id 使用者 ID
// @return error 錯誤訊息
func (r *UserRepository) Delete(id uint) error {
	result := r.db.Delete(&models.User{}, id)
	if result.Error != nil {
		logutil.Logger.Errorf("Error deleting user from database: %v", result.Error) // 記錄資料庫錯誤
		return result.Error
	}
	logutil.Logger.Debugf("User deleted from database with ID: %d", id) // 記錄使用者已刪除
	return nil
}
