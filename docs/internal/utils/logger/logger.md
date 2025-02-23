# internal/utils/logger 目錄

此目錄包含日誌相關的函數。

## 檔案

*   **`logger.go`**: 日誌的初始化和配置。

## 說明

*   `logger.go` 定義了 `Logger` 全域變數，用於記錄日誌。
*   `Init` 函數用於初始化日誌，接收 `Config` 結構體作為參數。
*   使用 `go.uber.org/zap` 和 `gopkg.in/natefinch/lumberjack.v2` 庫來記錄日誌。
*   支援的日誌等級：`debug`、`info`、`warn`、`error`、`dpanic`、`panic`、`fatal`。
*   日誌可以輸出到控制台和檔案。
*   日誌檔案位於 `logs/app` (預設路徑)，可以通過環境變數 `LOG_FILENAME` 修改。
*   日誌檔案會根據日期自動輪換。
*   `Close` 函數用於關閉日誌。
