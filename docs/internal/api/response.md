# internal/api/handlers/response 目錄

此目錄包含 HTTP 回應的結構定義。

## 檔案

*   **`response.go`**: 定義 API 回應的結構體。

## 說明

*   `response.go` 定義了 `Response` 結構體，用於統一 API 的回應格式。
*   `Response` 結構體包含 `Code`、`Message` 和 `Data` 欄位。
    *   `Code`: 回應的狀態碼。
    *   `Message`: 回應的訊息。
    *   `Data`: 回應的資料，可以是任何型別。
*   可以使用 `SuccessResponse` 和 `ErrorResponse` 函數來建立 `Response` 結構體的實例。
