# internal/utils/database 目錄

此目錄包含資料庫連線相關的函數。

## 檔案

- **`database.go`**: 資料庫連線的初始化和管理。

## 說明

- `database.go` 定義了 `Start` 函數，用於初始化資料庫連線。
- `Start` 函數接收一個 `configs.Config` 結構體作為參數，其中包含了資料庫連線所需的資訊。
- 使用 GORM 庫和 Postgres 驅動程式來連線到資料庫。
- 如果連線失敗，則記錄錯誤並返回 `nil`。
- 如果連線成功，則將 `gorm.DB` 實例儲存在 `dbInstance` 變數中，並返回該實例。
- 新增了 `Service` 介面，定義資料庫操作方法。
