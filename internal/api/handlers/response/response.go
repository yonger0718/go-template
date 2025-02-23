package response

import (
	"github.com/gin-gonic/gin"
	"go-template/internal/api/handlers/exception"
)

// SuccessData 回應成功的 JSON Struct
type SuccessData struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ErrorData 回應錯誤的 JSON Struct
type ErrorData struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// Success 回應成功的 JSON 數據
func Success(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, SuccessData{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// Error 回應錯誤的 JSON 數據
func Error(c *gin.Context, statusCode int, errCode int) {
	c.JSON(statusCode, ErrorData{
		Success: false,
		Message: exception.GetErrorMessage(errCode), // 使用 errors.go 中的 GetErrorMessage 函數取得錯誤訊息
	})
}
