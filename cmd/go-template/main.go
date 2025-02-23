package main

import (
	"context"
	"errors"
	"go-template/internal/configs"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	_ "go-template/assets/swagger" // 導入 docs package，這個 package 由 swag init 產生
	"go-template/internal/utils/logger"
)

// @title Go Template API
// @version 1.0
// @description This is a sample server Go Template server.
// @termsOfService http://swagger.io/terms/

// @host localhost:8080
// @BasePath /api
// @schemes http https

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	// 載入配置
	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// 初始化日誌
	_ = logger.Init(&cfg.Logger)
	// 確保程式結束前刷新日誌
	defer func(Logger *zap.SugaredLogger) {
		err := Logger.Sync()
		if err != nil {
			logger.Logger.Fatal("failed to sync logger", zap.Error(err))
		}
		logger.Close()
	}(logger.Logger)

	// 可以開始使用 logger 記錄日誌
	logger.Logger.Info("Configuration loaded successfully")

	// 初始化 server (使用 Wire 進行依賴注入)
	srv, cleanup, err := InitializeServer(cfg)
	if err != nil {
		logger.Logger.Fatalf("failed to initialize server: %v", err)
	}
	defer cleanup()

	logger.Logger.Info("Server is starting...")

	// 建立 channel 來接收中斷訊號
	done := make(chan bool, 1)
	go gracefulShutdown(srv, done)

	// 啟動 server
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Logger.Fatalf("listen: %s", err)
	}

	// 等待 gracefulShutdown 完成
	<-done
	logger.Logger.Info("Graceful shutdown complete.")
}

// gracefulShutdown 優雅地關閉 server
func gracefulShutdown(server *http.Server, done chan bool) {
	// 建立 context 來監聽中斷訊號 (SIGINT, SIGTERM)
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// 等待中斷訊號
	<-ctx.Done()
	logger.Logger.Info("shutting down gracefully, press Ctrl+C again to force")

	// 建立一個 5 秒的 timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 關閉 server
	if err := server.Shutdown(ctx); err != nil {
		logger.Logger.Fatalf("server forced to shutdown with error: %v", err)
	}

	logger.Logger.Info("Server exiting")

	// 通知 main goroutine server 已關閉
	done <- true
}
