# migrations 目錄

此目錄包含資料庫遷移相關的程式碼。

## 檔案

*   **`migrate.go`**: 執行資料庫遷移的程式。

## 說明

*   `migrate.go` 使用 GORM 的 `AutoMigrate` 功能來自動遷移資料庫結構。
*   執行 `go run migrations/migrate.go` 來執行資料庫遷移。

**警告：** `AutoMigrate` 在生產環境中有一定的風險，建議在生產環境中使用專業的資料庫遷移工具，例如 `golang-migrate/migrate`。