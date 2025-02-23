# Go Template

這是一個 Go 語言的專案範本，基於以下的服務搭建而成：

*   **Gin:** 高效能的 HTTP Web 框架
*   **GORM:** 強大的 ORM 庫
*   **JWT:** JSON Web Token 身份驗證
*   **Zap:** 高效能日誌庫
*   **Wire:** 編譯時依賴注入
*   **Lumberjack:** 日誌滾動和歸檔

## 專案結構

```
project-root/
├── assets/                    # 靜態資源目錄
│   └── swagger/              # Swagger API 文件
│       ├── docs.go           # Swagger 自動生成的文件
│       ├── swagger.json      # Swagger API 定義檔
│       └── swagger.yaml      # Swagger API YAML 格式定義檔
├── cmd/                      # 應用程式進入點
│   └── go-template/         # 主要應用程式
│       ├── main.go          # 程式進入點
│       ├── wire.go          # 依賴注入配置
│       └── wire_gen.go      # 自動生成的依賴注入程式碼
├── docs/                     # 文件目錄
│   ├── internal/            # 內部模組文件
│   ├── assets/              # 文件相關資源
│   └── cmd/                 # 命令列工具文件
├── internal/                 # 私有程式碼
│   ├── api/                 # API 相關邏輯
│   │   ├── handlers/       # HTTP 請求處理
│   │   │   ├── user/      # 使用者相關處理
│   │   │   ├── response/  # 回應格式處理
│   │   │   └── errors.go  # 錯誤碼定義
│   │   │   └── routes/        # 路由定義
│   ├── models/             # 資料模型
│   ├── repositories/       # 資料庫操作
│   ├── services/          # 業務邏輯
│   │   └── user/          # 使用者相關服務
│   ├── server/            # HTTP 伺服器
│   ├── middleware/        # 中介軟體
│   ├── validators/        # 參數驗證
│   ├── utils/             # 工具函數
│   │   ├── database/     # 資料庫連線
│   │   ├── jwt/         # JWT 相關
│   │   ├── logger/      # 日誌相關
│   │   └── response/    # HTTP 回應輔助
│   └── configs/          # 配置管理
├── pkg/                    # 可被外部導入的程式碼
├── tests/                  # 測試程式碼
├── scripts/                # 腳本
├── migrations/             # 資料庫遷移
│   └── migrate.go         # 遷移程式
├── .env                    # 環境變數設定檔
└── .env.example           # 環境變數範例檔
```

## 目錄說明

*   **`assets/`**: 靜態資源目錄。
    *   **`swagger/`**: Swagger API 文件相關檔案。
        *   **`docs.go`**: Swagger 自動生成的文件。
        *   **`swagger.json`**: Swagger API 定義檔。
        *   **`swagger.yaml`**: Swagger API YAML 格式定義檔。
*   **`cmd/`**: 應用程式的進入點。
    *   **`server/`**: 伺服器相關的程式碼。
        *   **`main.go`**: 伺服器啟動的程式碼。
        *   **`wire.go`**: Wire 依賴注入的配置檔案。
        *   **`wire_gen.go`**: Wire 自動產生的檔案，**請勿手動修改**。
*   **`internal/`**: 存放私有程式碼，不應該被外部專案導入。
    *   **`api/`**: API 相關的邏輯。
        *   **`handlers/`**: HTTP 請求的處理函數。
            *   **`user/`**: 使用者相關的 handlers。
            *   **`response/`**: 回應格式相關的處理。
            *   **`errors.go`**: API 錯誤碼定義。
        *   **`routes/`**: 路由定義。
    *   **`models/`**: 定義資料模型。
    *   **`repositories/`**: 定義資料庫操作的介面和實作。
    *   **`services/`**: 定義業務邏輯的介面和實作。
        *   **`user/`**: 使用者相關服務的介面和實作。
    *   **`server/`**: HTTP 伺服器的設定和啟動。
    *   **`middleware/`**: 中介軟體，例如身份驗證。
    *   **`validators/`**: 請求參數驗證函數。
    *   **`utils/`**: 通用的工具函數。
        *   **`database/`**: 資料庫連線相關。
        *   **`jwt/`**: JWT 相關。
        *   **`logger/`**: 日誌相關。
        *   **`response/`**: HTTP 回應輔助函數。
    *   **`configs/`**: 配置管理相關。
*   **`pkg/`**: 存放可以被外部專案導入的程式碼。
*   **`tests/`**: 存放測試程式碼。
*   **`scripts/`**: 存放開發、部署等腳本。
*   **`migrations/`**: 存放資料庫遷移腳本。
    *   **`migrate.go`**: 執行資料庫遷移的程式。
*   **`docs/`**: 專案文件目錄。
    *   **`internal/`**: 內部模組的文件。
    *   **`assets/`**: 文件相關的資源。
    *   **`cmd/`**: 命令列工具的文件。
*   **`.env`**: 環境變數設定檔，**不應該提交到版本控制**。
*   **`.env.example`**: 環境變數範例檔。

## 開發指南

請參閱 [快速開始教學](./docs/tutorials/getting-started.md) 以取得詳細的專案設定和執行說明。

詳細開發流程及規範，請參考 [開發指南](./docs/tutorials/development.md)。

## FAQ

1.  **如何修改資料庫設定？**
    - 修改 `.env` 檔案中的 `DB_HOST`、`DB_PORT`、`DB_USERNAME`、`DB_PASSWORD` 和 `DB_DATABASE` 環境變數。
2.  **如何修改 JWT 相關設定？**
    - 修改 `.env` 檔案中的 `JWT_SECRET`、`JWT_OLD_SECRETS` 和 `TOKEN_EXPIRES_IN` 環境變數。
3.  **如何修改日誌設定？**
    - 修改 `.env` 檔案中的 `LOG_LEVEL`、`LOG_FILE`、`LOG_MAX_SIZE`、`LOG_MAX_AGE`、`LOG_MAX_BACKUPS`、`LOG_LOCAL_TIME`、`LOG_COMPRESS` 和 `LOG_CONSOLE_OUT` 環境變數。
4.  **如何新增新的 API 路由？**
    - 在 `internal/api/handlers` 中新增 controller。
    - 在 `internal/api/routes` 中新增路由。
    - 在 `internal/services` 中新增對應的 service。
    - 在 `cmd/go-template/wire.go` 中加入新的依賴項。
    - 執行 `make generate` 或在 `cmd/go-template/` 目錄下執行 `./...` 重新產生 `cmd/go-template/wire_gen.go`。
5.  **如何執行測試？**
    - 執行 `go test ./...`。

## 參考資料

* [社群版 Golang 結構](https://github.com/golang-standards/project-layout/blob/master/README_zh-TW.md)
* [官方提供 Golang 結構](https://go.dev/doc/modules/layout)
> 其實社群上有一派的說法是認為因為在 `go` 中的包並不是以路徑導向的，而是以名稱導向，所以相較起來放在哪裡比較不重要。
> 
> 故我們應該更注重於程式碼的品質，而非程式的 layout

實際參考專案: 
* [prometheus/prometheus](https://github.com/prometheus/prometheus)
* [prometheus/common](https://github.com/prometheus/common)
* [gothinkster/golang-gin-realworld-example-app](https://github.com/gothinkster/golang-gin-realworld-example-app)

> 底下這篇 reddit 討論了為甚麼 DI 永遠都應該要實作
> 
> <s>所以我就花了更多的時間為了搞定DI</s>
* [Is DI in Go a thing?](https://www.reddit.com/r/golang/comments/wbawx5/is_dependency_injection_in_go_a_thing/?rdt=46475)
