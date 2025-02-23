# internal/utils/jwt 目錄

此目錄包含 JWT 相關的函數。

## 檔案

*   **`jwt.go`**: JWT 的產生和驗證。

## 說明

*   `jwt.go` 定義了 `Service` 結構體，用於產生和驗證 JWT token。
*   `NewService` 函數用於建立 `Service` 實例，並接收 `configs.Config` 作為參數。
*   `GenerateToken` 函數用於產生 JWT token，其中包含了使用者 ID、發行者、過期時間等資訊。
*   `ValidateToken` 函數用於驗證 JWT token，會先嘗試使用當前密鑰驗證，如果失敗則嘗試使用舊密鑰驗證。
*   `validateTokenWithSecret` 函數使用指定的 secret 驗證 token。
*   `jwt.go` 使用 `github.com/golang-jwt/jwt/v5` 庫來產生和驗證 JWT。
