package response

import (
	"github.com/gin-gonic/gin"
	"go-template/internal/api/handlers"
)

// SuccessResponse 回應成功的 JSON 數據
func SuccessResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, gin.H{
		"success": true,
		"message": message,
		"data":    data,
	})
}

// ErrorResponse 回應錯誤的 JSON 數據
func ErrorResponse(c *gin.Context, statusCode int, errCode int) {
	c.JSON(statusCode, gin.H{
		"success": false,
		"message": handlers.GetErrorMessage(errCode), // 使用 errors.go 中的 GetErrorMessage 函數取得錯誤訊息
	})
}
