package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"
)

// 測試基礎日誌初始化，並確保 logger 正常運行
func TestInitLogger(t *testing.T) {
	t.Run("Default Config", func(t *testing.T) {
		tempDir := createTempLogDir(t)
		defer os.RemoveAll(tempDir)

		cfg := &Config{
			Level:       "info",
			Filename:    filepath.Join(tempDir, "test"),
			LocalTime:   true,
			Compress:    false,
			ConsoleOut:  true,
			ServiceName: "test-service",
			EnableFile:  true,
		}

		_ = Init(cfg)
		assert.NotNil(t, Logger)
		Logger.Info("test info message")
		Logger.Error("test error message")
		Close()
	})

	t.Run("Console Only", func(t *testing.T) {
		cfg := &Config{
			Level:       "info",
			Filename:    "",
			LocalTime:   true,
			Compress:    false,
			ConsoleOut:  true,
			ServiceName: "test-service",
			EnableFile:  false,
		}

		_ = Init(cfg)
		assert.NotNil(t, Logger)
		Logger.Info("this should appear on console only")
		Close()
	})

	t.Run("File Only", func(t *testing.T) {
		tempDir := createTempLogDir(t)
		defer os.RemoveAll(tempDir)

		cfg := &Config{
			Level:       "info",
			Filename:    filepath.Join(tempDir, "test"),
			LocalTime:   true,
			Compress:    false,
			ConsoleOut:  false,
			ServiceName: "test-service",
			EnableFile:  true,
		}

		_ = Init(cfg)
		assert.NotNil(t, Logger)
		Logger.Info("this should be in file only")
		Close()
	})
}

// 測試日誌輪轉機制
func TestLogRotation(t *testing.T) {
	tempDir := createTempLogDir(t)
	defer os.RemoveAll(tempDir)

	logFileBase := filepath.Join(tempDir, "test")
	cfg := &Config{
		Level:       "debug",
		Filename:    logFileBase,
		LocalTime:   true,
		Compress:    false,
		ConsoleOut:  false,
		ServiceName: "test-service",
		EnableFile:  true,
	}

	// 模擬午夜前時間
	mockedTime := time.Date(2024, time.February, 17, 23, 59, 59, 0, time.UTC)
	timeNow = func() time.Time { return mockedTime }
	lastRotate = mockedTime

	_ = Init(cfg)
	Logger.Info("Log before rotation")

	// 驗證第一天的日誌檔案是否存在
	expectedFile1 := fmt.Sprintf("%s-%s.log", logFileBase, "2024-02-17")
	_, err := os.Stat(expectedFile1)
	assert.NoError(t, err, "Expected log file does not exist before rotation")

	// 模擬跨天
	mockedTime = time.Date(2024, time.February, 18, 0, 1, 0, 0, time.UTC)
	timeNow = func() time.Time { return mockedTime }

	// 強制觸發檔案輪轉
	checkAndRotate(logFileBase)

	Logger.Info("Log after rotation")

	// 驗證新的日誌檔案是否建立
	expectedFile2 := fmt.Sprintf("%s-%s.log", logFileBase, "2024-02-18")
	_, err = os.Stat(expectedFile2)
	assert.NoError(t, err, "Expected rotated log file does not exist")

	// 確保舊日誌檔案仍然存在
	_, err = os.Stat(expectedFile1)
	assert.NoError(t, err, "Old log file should still exist")

	Close()
}

// 測試異常場景
func TestLoggerErrors(t *testing.T) {
	t.Run("Invalid Log Directory", func(t *testing.T) {
		cfg := &Config{
			Level:       "info",
			Filename:    "/root/forbidden/test",
			LocalTime:   true,
			Compress:    false,
			ConsoleOut:  false,
			ServiceName: "test-service",
			EnableFile:  true,
		}

		err := Init(cfg)
		assert.NoError(t, err, "Logger should fall back to console logging instead of failing")
	})

	t.Run("checkAndRotate when fileWriter is nil", func(t *testing.T) {
		fileWriter = nil
		assert.NotPanics(t, func() {
			checkAndRotate("dummy-filename")
		}, "checkAndRotate should not panic when fileWriter is nil")
	})

	t.Run("Close multiple times", func(t *testing.T) {
		Close()
		assert.NotPanics(t, func() {
			Close()
		}, "Close should be idempotent and not panic when called multiple times")
	})
}

// 測試 getZapLevel 的所有可能輸入值
func TestGetZapLevel(t *testing.T) {
	testCases := []struct {
		level    string
		expected zapcore.Level
	}{
		{"debug", zapcore.DebugLevel},
		{"info", zapcore.InfoLevel},
		{"warn", zapcore.WarnLevel},
		{"error", zapcore.ErrorLevel},
		{"dpanic", zapcore.DPanicLevel},
		{"panic", zapcore.PanicLevel},
		{"fatal", zapcore.FatalLevel},
		{"invalid", zapcore.InfoLevel}, // 預設值
	}

	for _, tc := range testCases {
		t.Run(tc.level, func(t *testing.T) {
			actual := getZapLevel(tc.level)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

// 工具函數：建立臨時目錄
func createTempLogDir(t *testing.T) string {
	tempDir, err := os.MkdirTemp("", "logger-test-")
	if err != nil {
		t.Fatal(err)
	}
	return tempDir
}
