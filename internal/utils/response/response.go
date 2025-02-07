package response

import (
	"github.com/gin-gonic/gin"
	"go-template/internal/api/handlers"
)

// SuccessResponseData 回應成功的 JSON Struct
type SuccessResponseData struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ErrorResponseData 回應錯誤的 JSON Struct
type ErrorResponseData struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// SuccessResponse 回應成功的 JSON 數據
func SuccessResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, SuccessResponseData{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// ErrorResponse 回應錯誤的 JSON 數據
func ErrorResponse(c *gin.Context, statusCode int, errCode int) {
	c.JSON(statusCode, ErrorResponseData{
		Success: false,
		Message: handlers.GetErrorMessage(errCode), // 使用 errors.go 中的 GetErrorMessage 函數取得錯誤訊息
	})
}
