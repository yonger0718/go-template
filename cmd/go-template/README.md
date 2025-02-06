# cmd/go-template 目錄

此目錄包含 HTTP 伺服器的程式碼。

## 檔案

*   **`main.go`**: 伺服器啟動的程式碼。
*   **`wire.go`**: Wire 依賴注入的配置檔案。
*   **`wire_gen.go`**: Wire 自動產生的檔案，**請勿手動修改**。

## 說明

*   `main.go` 檔案負責：
    *   載入配置
    *   初始化日誌
    *   使用 Wire 進行依賴注入
    *   啟動 HTTP 伺服器
*   `wire.go` 檔案定義了依賴項之間的綁定關係。
*   `wire_gen.go` 檔案由 Wire 自動產生，包含了依賴注入的程式碼。