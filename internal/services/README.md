# internal/services 目錄

此目錄包含業務邏輯的介面和實作。

## 子目錄

*   **`impl/`**: 業務邏輯的具體實作。

## 檔案

*   **`user_service.go`**: 使用者服務的介面定義。

## 說明

*   `services` 層定義了應用程式的業務邏輯，它不應該依賴於 HTTP 協議或其他外部的細節。
*   `services` 層的介面定義了業務邏輯的操作，而具體實作則在 `impl` 子目錄中。