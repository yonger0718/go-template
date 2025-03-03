# Repository

此目錄包含資料庫操作的介面和實作。

## 檔案

- **`user.go`**: 使用者資料的 CRUD 操作。

## 說明

- `user.go` 實現了使用者資料的 CRUD 操作。
- `NewUserRepository` 函數用於建立 `UserRepository` 結構體的實例。
- `UserRepository` 結構體包含了與資料庫互動的 `db` 欄位。
- 提供了 `Create`、`GetByID`、`GetByUsername`、`Update` 和 `Delete` 等方法。
- `user.go` 使用 GORM 來與資料庫互動。
