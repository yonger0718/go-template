package server

import (
	"fmt"
	"go-template/internal/configs"
	"gorm.io/gorm"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go-template/internal/api/handlers/routes"
	"go-template/internal/utils/jwt"
)

// Config Struct，用於設定 server
type Config struct {
	DB          *gorm.DB
	JwtService  *jwt.Service
	UserService *routes.UserRoutes
	Config      *configs.Config // 這是通用的配置，例如 AppPort
}

// NewServer 建立一個新的 HTTP server 實例
func NewServer(cfg Config) *http.Server {
	router := gin.Default()

	// 註冊使用者相關的路由
	cfg.UserService.RegisterUserRoutes(router)

	// 建立 HTTP server 實例
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Config.AppPort), // 使用配置中的 AppPort
		Handler:      router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
