# 最佳實踐指南

本指南提供了一些在 Go Template 專案中開發的最佳實踐建議。

## 程式碼風格

-   遵循 Go 語言的官方程式碼風格規範：[Effective Go](https://go.dev/doc/effective_go) 以及 [Google Go Style](https://google.github.io/styleguide/go/)
-   使用 `go fmt` 格式化程式碼。
-   使用 `golint` 或 `golangci-lint` 檢查程式碼風格。
-   撰寫清晰、簡潔、易於理解的程式碼。
-   避免過度設計和過度抽象。

## 錯誤處理

-   始終檢查錯誤。
-   使用 `errors.New` 或 `fmt.Errorf` 建立自定義錯誤。
-   在 API 回應中使用一致的錯誤格式 (參考 `internal/api/handlers/response/response.go`)。
-   不要忽略錯誤，除非你有充分的理由。
-   在日誌中記錄錯誤的詳細資訊。

## 日誌記錄

-   使用 `internal/utils/logger` 進行日誌記錄。
-   使用不同的日誌等級 (debug, info, warn, error) 來記錄不同嚴重程度的事件。
-   在日誌訊息中包含足夠的上下文資訊，以便於除錯。
-   不要在日誌中記錄敏感資訊 (例如密碼、API 金鑰)。

## 依賴注入

-   使用 Wire 進行依賴注入。
-   將依賴項定義在 `cmd/go-template/wire.go` 中。
-   每次修改 `wire.go` 後，執行 `make generate` 或在 `cmd/go-template/` 目錄下執行 `go generate ./...` 或 `wire gen ./...` 重新產生 `wire_gen.go`。

## 資料庫

-   使用 GORM 進行資料庫操作。
-   在 `internal/models` 中定義資料模型。
-   在 `internal/repositories` 中定義資料庫操作介面和實作。
-   使用 `AutoMigrate` 進行資料庫遷移 (僅限開發環境)。
-   在生產環境中使用專業的資料庫遷移工具 (例如 `golang-migrate/migrate`)。

## API 設計

-   遵循 RESTful API 設計原則。
-   使用清晰、一致的 URL 路徑。
-   使用標準的 HTTP 方法 (GET, POST, PUT, DELETE)。
-   使用 JSON 作為請求和回應的格式。
-   使用 HTTP 狀態碼表示請求的結果。
-   使用 Swagger 生成 API 文件。

## 測試

-   編寫單元測試和整合測試。
-   使用 Go 內建的 `testing` 套件。
-   使用測試框架 (例如 `testify`) 來簡化測試程式碼。
-   確保測試覆蓋率足夠高。

## 安全性

-   驗證使用者輸入。
-   使用 JWT 進行身份驗證。
-   保護敏感資訊 (例如密碼、API 金鑰)。
-   定期更新依賴項，以修復安全漏洞。
-   避免常見的 Web 安全漏洞 (例如 XSS、CSRF、SQL 注入)。

## 效能

-   避免不必要的資料庫查詢。
-   使用快取來減少資料庫負載。
-   優化程式碼以減少 CPU 和記憶體使用量。
-   使用效能分析工具 (例如 `pprof`) 來識別效能瓶頸。

## 其他

-   保持程式碼模組化。
-   使用有意義的變數和函數名稱。
-   撰寫清晰的註解。
-   定期重構程式碼，以提高可讀性和可維護性。
-   使用 Git 進行版本控制。
-   使用 CI/CD 進行自動化建置、測試和部署。
