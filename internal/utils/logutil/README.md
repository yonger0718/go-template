# internal/utils/logutil 目錄

此目錄包含日誌相關的函數。
> 這個目錄其實意外的常見

## 檔案

*   **`logutil.go`**: 日誌的初始化和配置。

## 說明

*   `logutil.go` 使用 `go.uber.org/zap` 和 `github.com/natefinch/lumberjack` 庫來記錄日誌。
*   `InitLogger` 函數用於初始化日誌。
*   `Logger` 變數是用於記錄日誌的全域 `zap.SugaredLogger` 實例。
*   支援的日誌等級：`debug`、`info`、`warn`、`error`、`dpanic`、`panic`、`fatal`。
*   日誌會輸出到控制台和檔案 (如果 `LOG_CONSOLE_OUT` 設定為 `true`)。
*   日誌檔案位於 `logs/app.log` (預設路徑)，可以通過環境變數 `LOG_FILE` 修改。
*   可以透過環境變數配置日誌的各項參數 (例如日誌等級、檔案大小、保留時間等)。