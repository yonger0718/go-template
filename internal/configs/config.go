package configs

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"go-template/internal/utils/logutil"
)

// Config struct，定義了應用程式的配置
type Config struct {
	DBHost         string         // 資料庫主機
	DBPort         int            // 資料庫埠號
	DBUser         string         // 資料庫使用者名稱
	DBPassword     string         // 資料庫密碼
	DBName         string         // 資料庫名稱
	JWTSecret      string         // JWT 密鑰
	JWTOldSecrets  []string       // 舊的 JWT 密鑰，用於支援密鑰輪換
	TokenExpiresIn time.Duration  // Token 過期時間
	AppPort        int            // 應用程式埠號
	Logger         logutil.Config // 日誌配置
}

// LoadConfig 載入配置
func LoadConfig() (*Config, error) {
	// 預設先讀取專案跟目錄的 .env 檔案
	// 使用 godotenv.Load() 就不會有 autoload 的問題產生
	if err := godotenv.Load(); err != nil {
		// 這邊單純判斷找不到檔案，如果是其他類型的錯誤依舊會報錯
		if !errors.Is(err, os.ErrNotExist) {
			return nil, fmt.Errorf("error loading .env file: %w", err)
		}
	}

	// 讀取 DB_PORT 環境變數，如果不存在則預設為 5432
	dbPort, err := strconv.Atoi(getEnv("DB_PORT", "5432"))
	if err != nil {
		return nil, fmt.Errorf("invalid DB_PORT: %w", err)
	}

	// 讀取 PORT 環境變數，如果不存在則預設為 8080
	appPort, err := strconv.Atoi(getEnv("PORT", "8080"))
	if err != nil {
		return nil, fmt.Errorf("invalid PORT: %w", err)
	}

	// 讀取 TOKEN_EXPIRES_IN 環境變數，如果不存在則預設為 24 (小時)
	tokenExpiresIn, err := strconv.ParseInt(getEnv("TOKEN_EXPIRES_IN", "24"), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid TOKEN_EXPIRES_IN: %w", err)
	}

	// 讀取 JWT_SECRET
	jwtSecret := getEnv("JWT_SECRET", "")

	// 讀取 JWT_OLD_SECRETS (用逗號分隔的多個舊密鑰)
	var jwtOldSecrets []string
	if oldSecretsStr := getEnv("JWT_OLD_SECRETS", ""); oldSecretsStr != "" {
		jwtOldSecrets = strings.Split(oldSecretsStr, ",")
	}

	// 建立 Config 結構體並返回
	return &Config{
		DBHost:         getEnv("DB_HOST", "localhost"), // 預設為 localhost
		DBPort:         dbPort,
		DBUser:         getEnv("DB_USERNAME", "postgres"), // 預設為 postgres
		DBPassword:     getEnv("DB_PASSWORD", ""),         // 預設為空
		DBName:         getEnv("DB_DATABASE", "mydb"),     // 預設為 mydb
		JWTSecret:      jwtSecret,
		JWTOldSecrets:  jwtOldSecrets,
		TokenExpiresIn: time.Duration(tokenExpiresIn) * time.Hour, // 將小時轉換成 time.Duration
		AppPort:        appPort,
		Logger: logutil.Config{
			Level:       getEnv("LOG_LEVEL", "info"),
			Filename:    getEnv("LOG_FILENAME", "logs/app"),
			LocalTime:   getBoolEnv("LOG_LOCAL_TIME", true),
			Compress:    getBoolEnv("LOG_COMPRESS", true),
			ConsoleOut:  getBoolEnv("LOG_CONSOLE_OUT", true),   // 開發環境建議開啟
			ServiceName: getEnv("SERVICE_NAME", "go-template"), // 新增服務名稱
			EnableFile:  getBoolEnv("LOG_ENABLE_FILE", true),
		},
	}, nil
}

// getEnv 是一個輔助函數，用於取得環境變數，並在環境變數不存在時提供預設值
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

//// getIntEnv 是一個輔助函數，用於取得環境變數 (整數)，並在環境變數不存在時提供預設值
//func getIntEnv(key string, defaultValue int) int {
//	valueStr := getEnv(key, "")
//	if value, err := strconv.Atoi(valueStr); err == nil {
//		return value
//	}
//	return defaultValue
//}

// getBoolEnv 是一個輔助函數，用於取得環境變數 (布林值)，並在環境變數不存在時提供預設值
func getBoolEnv(key string, defaultValue bool) bool {
	valueStr := getEnv(key, "")
	if value, err := strconv.ParseBool(valueStr); err == nil {
		return value
	}
	return defaultValue
}
