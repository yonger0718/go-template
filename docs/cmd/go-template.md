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
    *   `gracefulShutdown` 函數處理伺服器的優雅關閉，確保在收到中斷訊號 (SIGINT, SIGTERM) 時，伺服器能夠完成當前請求並關閉。
*   `main.go` 檔案開頭的註解是 Swagger 註解，用於產生 API 文件。這些註解描述了 API 的標題、版本、描述、服務條款、主機、基本路徑、協定和安全性定義。
*   在 `main.go` 中使用了 `defer` 關鍵字確保在程式結束前會執行 logger.Sync() 以及 logger.Close()，確保所有紀錄都會被寫入並且關閉 logger。
*   `wire.go` 檔案定義了依賴項之間的綁定關係。
*   `wire_gen.go` 檔案由 Wire 自動產生，包含了依賴注入的程式碼。
