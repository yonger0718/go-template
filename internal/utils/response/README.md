# internal/utils/response 目錄

此目錄包含 HTTP 回應的輔助函數。

## 檔案

*   **`response.go`**:  用於建立 JSON 格式的回應。

## 說明

*   `response.go` 提供了 `SuccessResponse` 和 `ErrorResponse` 兩個函數，用於快速建立 JSON 格式的回應。
*   `SuccessResponse` 用於回應成功的請求。
*   `ErrorResponse` 用於回應錯誤的請求，並使用 `internal/api/handlers/errors.go` 中定義的錯誤碼。