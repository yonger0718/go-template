# internal 目錄

此目錄包含應用程式的私有程式碼。這些程式碼不應該被外部專案導入。

## 子目錄

*   **`api/`**: API 相關的邏輯，包括 handlers、routes、errors 等。
*   **`models/`**: 資料模型定義。
*   **`repositories/`**: 資料庫操作的介面和實作。
*   **`services/`**: 業務邏輯的介面和實作。
*   **`server/`**:  HTTP 伺服器的設定和啟動。
*   **`middleware/`**: 中介軟體，例如身份驗證。
*   **`validators/`**: 請求參數驗證函數。
*   **`utils/`**: 通用的工具函數。
*   **`configs/`**: 配置管理相關的函數。

## 說明

*   `internal` 目錄的設計遵循 Go 語言的專案結構慣例，用於存放私有程式碼。
*   `internal` 目錄下的程式碼只能被同一個 `internal` 層級或其子層級的程式碼引用。
*   將程式碼放在 `internal` 目錄下可以避免程式碼被外部專案意外地導入和使用。