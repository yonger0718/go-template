package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// TODO 增加env的格式來選擇要開啟的模式單純為std或file

// Logger 全局日誌變數
var Logger *zap.SugaredLogger

// Config 日誌配置
type Config struct {
	Level       string // 日誌等級: debug, info, warn, error, dpanic, panic, fatal
	Filename    string // 日誌檔案路徑 (基本名稱，不考慮日期)
	LocalTime   bool   // 是否使用本地時間
	Compress    bool   // 是否壓縮日誌檔案
	ConsoleOut  bool   // 是否輸出到控制台
	ServiceName string // 服務名稱
	EnableFile  bool
}

// defaultConfig 預設的日誌配置
var defaultConfig = Config{
	Level:       "info",
	Filename:    "logs/app", // 預設日誌檔案路徑，會放在專案的 logs 資料夾下
	LocalTime:   true,
	Compress:    true,
	ConsoleOut:  true, // 開發環境下建議開啟
	ServiceName: "go-template",
	EnableFile:  true,
}

var (
	fileWriter *lumberjack.Logger // 全局變數，用於檔案輸出
	mu         sync.Mutex         // 互斥鎖，用於同步日誌檔案的切換
	timeNow    = time.Now         // 新增 timeNow 變數，預設值為 time.Now
	lastRotate time.Time          // 記錄上次旋轉的日期 * 提取到全域變數
	checkEvery = time.Minute      // 設定檢查時間
)

// Init 初始化日誌
func Init(cfg *Config) error {
	if cfg == nil {
		cfg = &defaultConfig
	}
	// 根據設定的日誌等級建立 zapcore.Level
	level := getZapLevel(cfg.Level)
	// 建立 encoder 配置
	encoderConfig := zap.NewProductionEncoderConfig()
	if cfg.LocalTime {
		encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339) // 使用 RFC3339 時間格式
	} else {
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // 使用 ISO8601 時間格式
	}
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // 使用大寫字母記錄日誌等級
	encoder := zapcore.NewJSONEncoder(encoderConfig)        // 使用 JSON 格式編碼

	var core zapcore.Core

	// 檔案輸出
	if cfg.EnableFile {
		// 取得日誌檔案的目錄和基本名稱
		logDir := filepath.Dir(cfg.Filename)

		// 確保日誌目錄存在
		if _, err := os.Stat(logDir); os.IsNotExist(err) {
			if err := os.MkdirAll(logDir, 0750); err != nil {
				fmt.Fprintf(os.Stderr, "Error creating log directory: %v\n", err)
				fmt.Fprintln(os.Stderr, "Falling back to console logging")
				cfg.EnableFile = false // 直接切換為 Console 模式
			}
		}

		if cfg.EnableFile {
			// 賦值給全域變數 fileWriter
			fileWriter = &lumberjack.Logger{
				Filename:  getLogFileName(cfg.Filename),
				LocalTime: cfg.LocalTime,
				Compress:  cfg.Compress,
			}
			fileCore := zapcore.NewCore(encoder, zapcore.AddSync(fileWriter), level)
			if cfg.ConsoleOut {
				consoleCore := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), level)
				core = zapcore.NewTee(fileCore, consoleCore)
			} else {
				core = fileCore
			}

			lastRotate = timeNow() // 初始化 lastRotate
			// 啟動一個 goroutine 來定期檢查日期是否改變
			go rotateByDate(cfg.Filename)
		}
	} else {
		core = zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), level)
	}

	// 建立 logger
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	// 添加服務名稱欄位
	logger = logger.With(zap.String("service", cfg.ServiceName))

	// 設定全域 logger
	Logger = logger.Sugar()

	// 紀錄日誌初始化完成
	Logger.Infof("Logger initialized with level: %s", cfg.Level)
	Logger.Infof("Logger output to file is enabled: %t", cfg.EnableFile)

	return nil
}

func getLogFileName(filename string) string {
	return fmt.Sprintf("%s-%s.log", filename, timeNow().Format("2006-01-02"))
}

// rotateByDate 定期檢查日期是否改變，如果改變則觸發日誌檔案滾動
func rotateByDate(baseFilename string) {
	ticker := time.NewTicker(checkEvery)
	defer ticker.Stop()

	for range ticker.C {
		checkAndRotate(baseFilename)
	}
}

// checkAndRotate 檢查日期並執行日誌輪換 (提取成獨立函數)
func checkAndRotate(baseFilename string) {
	now := timeNow()
	mu.Lock()
	defer mu.Unlock()

	if fileWriter == nil {
		return
	}

	if now.Format("2006-01-02") != lastRotate.Format("2006-01-02") {
		fileWriter.Filename = getLogFileName(baseFilename)
		if err := fileWriter.Rotate(); err != nil {
			fmt.Fprintf(os.Stderr, "Error rotating log file: %v\n", err)
		}
		Logger.Infof("Log file rotated to: %s", fileWriter.Filename)
		lastRotate = now // 更新 lastRotate 的值
	}
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

// Close 關閉 Logger
func Close() {
	if fileWriter != nil {
		_ = fileWriter.Close()
	}
}
