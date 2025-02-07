package models

import (
	"gorm.io/gorm"
	"time"
)

// User 定義使用者資料 Struct
type User struct {
	gorm.Model           // gorm.Model 包含了ID、CreatedAt、UpdatedAt字段
	Username   string    `json:"username"    validate:"required"       gorm:"unique;not null"` // 帳號名稱
	Email      string    `json:"email"       validate:"required,email" gorm:"unique;not null"` // 電子郵件
	Password   string    `json:"-"           validate:"required"       gorm:"not null"`        // 密碼
	LastLogin  time.Time `json:"last_login"`                                                   // 最後登入時間
	Status     int       `json:"status"`                                                       // 帳號狀態
}

// TableName 表名可以自定義
func (User) TableName() string {
	return "users"
}
