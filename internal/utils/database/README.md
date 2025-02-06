# internal/utils/database 目錄

此目錄包含資料庫連線相關的函數。

## 檔案

*   **`database.go`**: 資料庫連線的初始化和管理。

## 說明

*   `database.go` 使用 GORM 庫來連線到資料庫。
*   `Start` 函數用於初始化資料庫連線。
*   `GetDB` 函數用於取得資料庫連線。
*   `Close` 函數用於關閉資料庫連線。