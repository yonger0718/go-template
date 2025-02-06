package handlers

// 定義使用者相關的錯誤碼
const (
	_ = iota // 忽略第一個值，從 1 開始
	ErrCodeUserNotFound
	ErrCodeInvalidCredentials
	ErrCodeInvalidRequest
	ErrCodeUnknown
	ErrCodeUserIDNotInContext
	ErrCodeUserIDFormatInvalid
)

// 定義通用的錯誤訊息常數
const (
	ErrMsgInvalidRequestBody  = "Invalid request body: %v"
	ErrMsgUserIDNotInContext  = "User ID not found in context"
	ErrMsgUserIDFormatInvalid = "Invalid user ID format"
)

// 定義錯誤碼和錯誤訊息的對應關係
var errorMessages = map[int]string{
	ErrCodeUserNotFound:        "user not found",
	ErrCodeInvalidCredentials:  "invalid credentials",
	ErrCodeInvalidRequest:      "invalid request",
	ErrCodeUnknown:             "unknown error",
	ErrCodeUserIDNotInContext:  "user ID not found in context",
	ErrCodeUserIDFormatInvalid: "user ID format invalid",
}

// GetErrorMessage 根據錯誤碼取得對應的錯誤訊息
func GetErrorMessage(errCode int) string {
	message, ok := errorMessages[errCode]
	if !ok {
		return errorMessages[ErrCodeUnknown]
	}
	return message
}
