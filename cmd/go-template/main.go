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

	"go-template/internal/utils/logutil"
)

func main() {
	// 載入配置
	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// 初始化日誌
	logutil.InitLogger(&cfg.Logger)
	// 確保程式結束前刷新日誌
	defer func(Logger *zap.SugaredLogger) {
		err := Logger.Sync()
		if err != nil {
			logutil.Logger.Fatal("failed to sync logger", zap.Error(err))
		}
	}(logutil.Logger)

	// 可以開始使用 logger 記錄日誌
	logutil.Logger.Info("Configuration loaded successfully")

	// 初始化 server (使用 Wire 進行依賴注入)
	srv, cleanup, err := InitializeServer(cfg)
	if err != nil {
		logutil.Logger.Fatalf("failed to initialize server: %v", err)
	}
	defer cleanup()

	logutil.Logger.Info("Server is starting...")

	// 建立 channel 來接收中斷訊號
	done := make(chan bool, 1)
	go gracefulShutdown(srv, done)

	// 啟動 server
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logutil.Logger.Fatalf("listen: %s", err)
	}

	// 等待 gracefulShutdown 完成
	<-done
	logutil.Logger.Info("Graceful shutdown complete.")
}

// gracefulShutdown 優雅地關閉 server
func gracefulShutdown(server *http.Server, done chan bool) {
	// 建立 context 來監聽中斷訊號 (SIGINT, SIGTERM)
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// 等待中斷訊號
	<-ctx.Done()
	logutil.Logger.Info("shutting down gracefully, press Ctrl+C again to force")

	// 建立一個 5 秒的 timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 關閉 server
	if err := server.Shutdown(ctx); err != nil {
		logutil.Logger.Fatalf("server forced to shutdown with error: %v", err)
	}

	logutil.Logger.Info("Server exiting")

	// 通知 main goroutine server 已關閉
	done <- true
}
