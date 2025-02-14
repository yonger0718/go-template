package main

import (
	"go-template/internal/configs"
	"go.uber.org/zap"
	"log"
	"reflect"

	"go-template/internal/models"
	"go-template/internal/utils/database"
	"go-template/internal/utils/logutil"
)

func main() {
	// 載入配置
	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// 初始化日誌
	_ = logutil.InitLogger(&cfg.Logger)
	// 確保程式結束前刷新日誌
	defer func(Logger *zap.SugaredLogger) {
		err := Logger.Sync()
		if err != nil {
			logutil.Logger.Error("Failed to flush logger: %v", err)
		}
	}(logutil.Logger)

	// 取得資料庫連線
	db := database.Start(cfg)
	if db == nil {
		logutil.Logger.Fatal("Failed to connect to database.")
	}

	// 定義需要遷移的 Model
	autoMigrateLists := []interface{}{
		&models.User{}, // 使用 models.User
	}

	// 執行 AutoMigrate
	logutil.Logger.Info("Starting auto migration...")
	for _, table := range autoMigrateLists {
		// 輸出當前遷移的 Model 名稱
		logutil.Logger.Infof("Migrating table: %s", reflect.TypeOf(table).Elem().Name())
		if err := db.AutoMigrate(table); err != nil {
			logutil.Logger.Fatalf("auto migration error: %v", err)
		}
	}

	logutil.Logger.Info("Auto migration successfully completed.")
}
