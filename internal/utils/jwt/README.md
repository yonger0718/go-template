# internal/utils/jwt 目錄

此目錄包含 JWT 相關的函數。

## 檔案

*   **`jwt.go`**: JWT 的產生和驗證。

## 說明

*   `jwt.go` 使用 `github.com/golang-jwt/jwt/v5` 庫來產生和驗證 JWT。
*   `NewJWTService` 函數用於建立一個新的 `JWTService` 實例。
*   `GenerateToken` 函數用於產生 JWT token。
*   `ValidateToken` 函數用於驗證 JWT token。