package database

import (
	"fmt"
	"go-template/internal/configs"

	"go-template/internal/utils/logutil"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Service 介面，定義資料庫操作方法
type Service interface {
	AutoMigrate(models ...interface{}) error
	GetDB() *gorm.DB
	Close() error
}

var dbInstance *gorm.DB

// Start 初始化資料庫連線
func Start(cfg *configs.Config) *gorm.DB {
	// 如果 dbInstance 已經存在，則直接返回
	if dbInstance != nil {
		return dbInstance
	}

	// 組成 Postgres 連線字串
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	// 連線到 Postgres 資料庫
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error), // 設定 logger 等級為 Error
	})

	// 如果連線失敗，則 panic
	if err != nil {
		logutil.Logger.Fatalf("failed to connect database, err: %v", err)
		return nil
	}

	// 將 dbService 實例賦值給 dbInstance
	dbInstance = db

	return dbInstance
}
