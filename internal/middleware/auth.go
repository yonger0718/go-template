package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-template/internal/api/handlers/exception"
	"go-template/internal/api/handlers/response"
	"go-template/internal/constants"
	"go-template/internal/utils/jwt"
	"go-template/internal/utils/logger"
)

// Auth 驗證 JWT token 的中介軟體
func Auth(jwtService *jwt.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		if jwtService == nil {
			logger.Logger.Error("jwtService is nil in Auth") // 新增日誌
			panic("jwtService is nil")                       // 或者返回錯誤，避免 panic
		}
		// 從 Authorization header 中取得 token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			logger.Logger.Debugf("Authorization header is missing")
			response.Error(c, http.StatusUnauthorized, exception.ErrCodeInvalidRequest)
			c.Abort() // 中止後續的處理函數
			return
		}

		// 解析 Bearer token
		tokenString := authHeader[len("Bearer "):]
		userID, err := jwtService.ValidateToken(tokenString)
		if err != nil {
			logger.Logger.Debugf("Invalid token: %v", err)
			response.Error(c, http.StatusUnauthorized, exception.ErrCodeInvalidCredentials)
			c.Abort() // 中止後續的處理函數
			return
		}

		// 將 userID 儲存到 gin.Context 中，方便後續的處理函數使用
		c.Set(constants.CtxUserIDKey, userID)

		// 呼叫下一個處理函數
		c.Next()
	}
}
