## 開發流程

1.  **定義資料模型 (`internal/models`)：**  根據你的業務需求定義資料模型。
2.  **定義資料庫操作 (`internal/repositories`)：**  定義資料庫操作的介面和實作。
3.  **定義業務邏輯 (`internal/services`)：**  定義業務邏輯的介面和實作。
4.  **編寫處理函數 (`internal/api/handlers`)：**  編寫 HTTP 請求的處理函數。
5.  **定義路由 (`internal/api/routes`)：**  將 HTTP 請求路由到對應的處理函數。
6.  **加入日誌：**  在程式碼中加入適當的日誌記錄，方便除錯和監控。
7.  **編寫測試：**  編寫單元測試和整合測試，確保程式碼的正確性。

## 新增功能

1.  **在 `internal/api/handlers` 中新增 controller (e.g., `product_handler.go`)。**
2.  **在 `internal/api/routes` 中新增路由 (e.g., `product_routes.go`)。**
3.  **在 `internal/models` 中新增 model (e.g., `product.go`)。**
4.  **在 `internal/repositories` 中新增 repository (e.g., `product_repository.go`)。**
5.  **在 `internal/services` 中新增 service (e.g., `product_service.go` 和 `impl/product_service_impl.go`)。**
6.  **在 `cmd/go-template/wire.go` 中加入新的依賴項。**
7.  **執行 `make genegare` 或在 `cmd/go-template/` 目錄下執行 `./...` 重新產生 `cmd/go-template/wire_gen.go`。**
8.  **在 `internal/server/server.go` 中的 `NewServer` 函數加入相關依賴 (如果需要)。**

## 日誌記錄

*   使用 `logutil.Logger` 記錄日誌。
*   支援的日誌等級：`debug`、`info`、`warn`、`error`、`dpanic`、`panic`、`fatal`。
*   預設日誌等級為 `info`，可以通過環境變數 `LOG_LEVEL` 修改。
*   日誌會同時輸出到控制台和檔案 (如果 `LOG_CONSOLE_OUT` 設定為 `true`)。
*   日誌檔案位於 `logs/app.log` (預設路徑)，可以通過環境變數 `LOG_FILE` 修改。

## 錯誤處理

*   錯誤碼和錯誤訊息定義在 `internal/api/handlers/exception/errors.go` 中。
*   使用 `response.ErrorResponse` 函數回覆錯誤訊息。

## 身份驗證

*   使用 JWT 進行身份驗證。
*   `AuthMiddleware` 中介軟體會驗證 `Authorization` header 中的 JWT token。
*   受保護的路由 (例如 `/users/:id` 的 GET, PUT, DELETE) 需要使用者提供有效的 JWT token 才能訪問。
*   **無 token 或 token 無效會返回 401 Unauthorized 錯誤。**

## 路由

*   `/users/register`: 使用者註冊 (POST)
*   `/users/login`: 使用者登入 (POST)
*   `/users/:id`: 取得、更新、刪除使用者資訊 (GET, PUT, DELETE) - 需要身份驗證

## JWT 密鑰輪換

*   可以透過設定 `JWT_OLD_SECRETS` 環境變數 (逗號分隔的多個舊密鑰) 來支援 JWT 密鑰輪換。
*   更新 `JWT_SECRET` 後，舊的 JWT 在過期之前仍然有效。
*   當所有舊的 JWT 都過期後，可以從 `JWT_OLD_SECRETS` 中移除舊的密鑰。

## 程式碼風格

*   請遵循 Go 語言的程式碼風格規範，主要follow [Google版規範](https://google.github.io/styleguide/go/)。

## 常見使用案例

1.  **新增使用者註冊和登入功能：**
    - 參考 `internal/api/handlers/user/user.go` 中的 `Register` 和 `Login` 函數。
    - 使用 `internal/services/user.go` 中的 `CreateUser` 和 `Login` 方法。
2.  **新增受保護的路由：**
    - 參考 `internal/middleware/auth.go` 中的 `AuthMiddleware`。
    - 在需要身份驗證的路由上加入 `AuthMiddleware`。
3.  **新增資料庫操作：**
    - 參考 `internal/repository/user.go` 中的 CRUD 操作。
    - 在 `internal/models` 中定義新的資料模型。
    - 在 `internal/repository` 中新增對應的 repository。
4.  **新增自定義錯誤處理：**
    - 參考 `internal/api/handlers/exception/errors.go` 中的錯誤碼定義。
    - 在 `internal/api/handlers/response/response.go` 中使用 `ErrorResponse` 函數。
