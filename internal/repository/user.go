package repository

import (
	"go-template/internal/models"
	"go-template/internal/utils/logger"
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
// @Param user body models.User true "新增的使用者資料"
// @return error "錯誤訊息"
func (repo *UserRepository) Create(user *models.User) error {
	result := repo.db.Create(user)
	if result.Error != nil {
		logger.Logger.Errorf("Error creating user in database: %v", result.Error) // 記錄資料庫錯誤
		return result.Error
	}
	logger.Logger.Debugf("User created in database: %s", user.Username) // 記錄使用者已建立
	return nil
}

// GetByID 根據 ID 取得使用者
// @param id path uint true "使用者 ID"
// @return models.User "使用者"
// @return error "錯誤訊息"
func (repo *UserRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	result := repo.db.First(&user, id)
	if result.Error != nil {
		logger.Logger.Errorf("Error getting user by ID from database: %v", result.Error) // 記錄資料庫錯誤
		return nil, result.Error
	}
	logger.Logger.Debugf("User found by ID in database: %d", id) // 記錄使用者已找到
	return &user, nil
}

// GetByUsername 根據使用者名稱取得使用者
// @param username path string true "使用者名稱"
// @return models.User "使用者"
// @return error "錯誤訊息"
func (repo *UserRepository) GetByUsername(username string) (*models.User, error) {
	var user models.User
	result := repo.db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		logger.Logger.Errorf("Error getting user by username from database: %v", result.Error) // 記錄資料庫錯誤
		return nil, result.Error
	}
	logger.Logger.Debugf("User found by username in database: %s", username) // 記錄使用者已找到
	return &user, nil
}

// Update 更新使用者資訊
// @Param user body models.User true "修改的使用者資料"
// @return error "錯誤訊息"
func (repo *UserRepository) Update(user *models.User) error {
	result := repo.db.Save(user)
	if result.Error != nil {
		logger.Logger.Errorf("Error updating user in database: %v", result.Error) // 記錄資料庫錯誤
		return result.Error
	}
	logger.Logger.Debugf("User updated in database: %s", user.Username) // 記錄使用者已更新
	return nil
}

// Delete 根據 ID 刪除使用者
// @param id path uint true "使用者 ID"
// @return error "錯誤訊息"
func (repo *UserRepository) Delete(id uint) error {
	result := repo.db.Delete(&models.User{}, id)
	if result.Error != nil {
		logger.Logger.Errorf("Error deleting user from database: %v", result.Error) // 記錄資料庫錯誤
		return result.Error
	}
	logger.Logger.Debugf("User deleted from database with ID: %d", id) // 記錄使用者已刪除
	return nil
}
