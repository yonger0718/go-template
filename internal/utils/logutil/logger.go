package logutil

import (
	"os"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Logger 全局日誌變數
var Logger *zap.SugaredLogger

// Config 日誌配置
type Config struct {
	Level       string // 日誌等級: debug, info, warn, error, dpanic, panic, fatal
	Filename    string // 日誌檔案路徑
	MaxSize     int    // 每個日誌檔案的最大大小 (MB)
	MaxAge      int    // 日誌檔案保留的最大天數
	MaxBackups  int    // 最多保留的日誌檔案數量
	LocalTime   bool   // 是否使用本地時間
	Compress    bool   // 是否壓縮日誌檔案
	ConsoleOut  bool   // 是否輸出到控制台
	ServiceName string // 服務名稱
}

// defaultConfig 預設的日誌配置
var defaultConfig = Config{
	Level:       "info",
	Filename:    "logs/app.log", // 預設日誌檔案路徑，會放在專案的 logs 資料夾下
	MaxSize:     100,            // 預設每個檔案 100 MB
	MaxAge:      30,             // 預設保留 30 天
	MaxBackups:  5,              // 預設保留 5 個檔案
	LocalTime:   true,
	Compress:    true,
	ConsoleOut:  true, // 開發環境下建議開啟
	ServiceName: "go-template",
}

// InitLogger 初始化日誌
func InitLogger(cfg *Config) {
	if cfg == nil {
		cfg = &defaultConfig
	}
	// 根據設定的日誌等級建立 zapcore.Level
	level := getZapLevel(cfg.Level)
	// 建立 encoder 配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // 使用 ISO8601 時間格式
	if cfg.LocalTime {
		encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339) // 使用 RFC3339 時間格式
	}
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // 使用大寫字母記錄日誌等級
	encoder := zapcore.NewJSONEncoder(encoderConfig)        // 使用 JSON 格式編碼

	// 建立日誌檔案的 writer
	fileWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   cfg.Filename,
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
		LocalTime:  cfg.LocalTime,
		Compress:   cfg.Compress,
	})

	// 建立同時輸出到控制台和檔案的 core
	var core zapcore.Core
	if cfg.ConsoleOut {
		// 開啟控制台輸出
		consoleWriter := zapcore.Lock(os.Stdout)
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, fileWriter, level),
			zapcore.NewCore(encoder, consoleWriter, level),
		)
	} else {
		core = zapcore.NewCore(encoder, fileWriter, level)
	}

	// 建立 logger
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	// 添加服務名稱欄位
	logger = logger.With(zap.String("service", cfg.ServiceName))

	// 設定全域 logger
	Logger = logger.Sugar()

	// 紀錄日誌初始化完成
	Logger.Infof("Logger initialized with level: %s", cfg.Level)
}

// getZapLevel 將字串表示的日誌等級轉換成 zapcore.Level
func getZapLevel(level string) zapcore.Level {
	switch strings.ToLower(level) {
	case "debug":
		return zapcore.DebugLevel // 輸出 debug, info, warn, error, dpanic, panic, fatal
	case "info":
		return zapcore.InfoLevel // 輸出 info, warn, error, dpanic, panic, fatal
	case "warn":
		return zapcore.WarnLevel // 輸出 warn, error, dpanic, panic, fatal
	case "error":
		return zapcore.ErrorLevel // 輸出 error, dpanic, panic, fatal
	case "dpanic":
		return zapcore.DPanicLevel // 輸出 dpanic, panic, fatal
	case "panic":
		return zapcore.PanicLevel // 輸出 panic, fatal
	case "fatal":
		return zapcore.FatalLevel // 輸出 fatal
	default:
		return zapcore.InfoLevel // 預設輸出 info, warn, error, dpanic, panic, fatal
	}
}
