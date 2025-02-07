package routes

import (
	"github.com/gin-gonic/gin"
	"go-template/internal/api/handlers/user"
	"go-template/internal/middleware"
	"go-template/internal/utils/jwt"
)

// UserRoutes 結構體，用於管理使用者相關的路由
type UserRoutes struct {
	handler    *user.Handler
	jwtService *jwt.Service
}

// NewUserRoutes 建立一個新的 UserRoutes 實例
func NewUserRoutes(handler *user.Handler, jwtService *jwt.Service) *UserRoutes {
	return &UserRoutes{handler: handler, jwtService: jwtService}
}

// RegisterUserRoutes 註冊使用者相關的路由
func (r *UserRoutes) RegisterUserRoutes(router *gin.Engine) {
	// 建立一個 /users 的路由群組
	userGroup := router.Group("/api/user")
	{
		// 公開路由 (不需要身份驗證)
		userGroup.POST("/register", r.handler.RegisterUser)
		userGroup.POST("/login", r.handler.LoginUser)

		// 受保護的路由 (需要身份驗證)
		// 將 AuthMiddleware 應用到 protectedGroup
		protectedGroup := userGroup.Group("/")
		protectedGroup.Use(middleware.AuthMiddleware(r.jwtService))
		{
			protectedGroup.GET("/:id", r.handler.GetUser)
			protectedGroup.PUT("/:id", r.handler.UpdateUser)
			protectedGroup.DELETE("/:id", r.handler.DeleteUser)
		}
	}
}
