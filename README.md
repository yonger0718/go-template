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
├── cmd/               # 主應用程序入口點
│   └── go-template/
│       └── main.go      # 程式進入點
│       └── wire.go      # Wire 依賴注入配置
│       └── wire_gen.go  # Wire 自動生成的檔案，請勿修改
│
├── docs/              # 文件
│   └── docs.go        # swagger API 文件，透過 gin-swagger 自動產生
│   └── swagger.json   # swagger API 文件，透過 gin-swagger 自動產生
│   └── swaager.yaml   # swagger API 文件，透過 gin-swagger 自動產生
│ 
├── internal/          # 私有代碼，不能被其他專案導入
│   ├── api/           # API 相關邏輯
│   │   ├── handlers/  # 控制器（handlers）
│   │   │   ├── user/
│   │   │   │   └── user_handler.go    # 使用者相關的處理函數
│   │   │   └──routes/                 # 路由定義
│   │   │       └── user_routes.go     # 使用者相關的路由
│   │   └── errors.go                  # 錯誤定義
│   │
│   ├── models/        # 數據模型
│   │   └── user.go    # 使用者模型
│   │
│   ├── repository/    # 數據訪問層
│   │   └── user.go    # 使用者資料的 CRUD 操作
│   │
│   ├── services/      # 業務邏輯層
│   │   └── impl
│   │       └── userImpl.go  # 使用者服務的具體實現
│   │   └── user.go          # 使用者服務的介面定義
│   │
│   ├── server/              # 後端伺服器
│   │   └── server.go        # 伺服器設定與啟動
│   │
│   ├── middleware/    # 中間件
│   │   └── auth_middleware.go # 身份驗證中間件
│   │
│   ├── validators/    # 請求驗證邏輯
│   │   └── user_validator.go # 使用者資料驗證
│   │
│   ├── utils/         # 通用工具函數
│   │   ├── database/
│   │   │   └── database.go # 資料庫連線
│   │   ├── jwt/
│   │   │   └── jwt.go      # JWT 相關函數
│   │   ├── logutil/        # 日誌工具
│   │   │   └── logutil.go  # 日誌函數
│   │   └── response/
│   │       └── response.go # HTTP 回應的輔助函數
│   │
│   └── configs/       # 配置管理
│       └── config.go  # 載入和解析配置
│
├── pkg/               # 可以被外部專案導入的庫
│
├── tests/             # 測試相關
│   ├── unit/          # 單元測試
│   ├── integration/   # 集成測試
│   └── mocks/         # 模擬對象
│
├── scripts/           # 開發、部署腳本
├── migrations/        # 數據庫遷移腳本
│   └── migrate.go     # 資料庫遷移程式
│
├── .env               # 環境變數設定檔 (請不要將其加入到版本控制中)
├── .env.example       # 環境變數範例檔
├── go.mod             # Go 模組檔案
└── go.sum             # Go 模組鎖定檔
```

## 目錄說明

*   **`cmd/`**: 應用程式的進入點。
    *   **`server/`**:  伺服器相關的程式碼。
        *   **`main.go`**: 伺服器啟動的程式碼。
        *   **`wire.go`**:  Wire 依賴注入的配置檔案。
        *   **`wire_gen.go`**:  Wire 自動產生的檔案，**請勿手動修改**。
*   **`internal/`**: 存放私有程式碼，不應該被外部專案導入。
    *   **`api/`**:  API 相關的邏輯，包括 handlers、routes 等。
        *   **`handlers/`**:  HTTP 請求的處理函數 (控制器)。
            *   **`user`**: 使用者相關的 handlers。
            *   **`errors.go`**: 定義了 API 相關的錯誤碼和錯誤訊息。
        *   **`routes/`**:  路由定義。
    *   **`models/`**:  定義資料模型。
    *   **`repositories/`**:  定義資料庫操作的介面和實作。
    *   **`services/`**:  定義業務邏輯的介面和實作。
    *   **`server/`**:  HTTP 伺服器的設定和啟動。
    *   **`middleware/`**:  中介軟體，例如身份驗證。
    *   **`validators/`**:  請求參數驗證函數。
    *   **`utils/`**:  通用的工具函數。
        *   **`database/`**:  資料庫連線相關的函數。
        *   **`jwt/`**:  JWT 產生和驗證相關的函數。
        *   **`logutil/`**:  日誌相關的函數。
        *   **`response/`**:  HTTP 回應的輔助函數。
    *   **`configs/`**:  配置管理相關的函數。
*   **`pkg/`**: 存放可以被外部專案導入的程式碼。
*   **`tests/`**: 存放測試程式碼。
*   **`scripts/`**: 存放開發、部署等腳本。
*   **`migrations/`**: 存放資料庫遷移腳本。
    *   **`migrate.go`**:  執行資料庫遷移的程式。
*   **`.env`**: 環境變數設定檔，**不應該提交到版本控制**。
*   **`.env.example`**: 環境變數範例檔，可以提交到版本控制，方便開發者了解需要設定哪些環境變數。

## 開發指南

### 前置要求

*   **Go:**  版本 1.18 或更高版本。
*   **PostgreSQL:**  用於資料庫。
*   **Git:**  用於版本控制。

### 安裝依賴
> 如果有需要參考完整安裝步驟的話，可以拉到本文的最底下。

```bash
go mod tidy
```

### 環境變數設定

請根據你的環境，在 `.env` 檔案中設定以下環境變數：
> 如果有需要，可以參考 `.env.example` 檔案作為參考。

```
DB_HOST=localhost
DB_PORT=5432
DB_USERNAME=your_db_user
DB_PASSWORD=your_db_password
DB_DATABASE=your_db_name
JWT_SECRET=your-secret-key # 請務必修改成高強度密鑰，並定期更新
JWT_OLD_SECRETS=          # 舊的 JWT 密鑰，用於支援密鑰輪換 (可選, 多組密鑰使用逗號分隔)
TOKEN_EXPIRES_IN=24        # 單位: 小時
PORT=8080
LOG_LEVEL=info             # 日誌等級: debug, info, warn, error, dpanic, panic, fatal
LOG_FILE=logs/app.log      # 日誌檔案路徑
LOG_MAX_SIZE=100           # 每個日誌檔案的最大大小 (MB)
LOG_MAX_AGE=30             # 日誌檔案保留的最大天數
LOG_MAX_BACKUPS=5          # 最多保留的日誌檔案數量
LOG_LOCAL_TIME=true        # 是否使用本地時間
LOG_COMPRESS=true          # 是否壓縮日誌檔案
LOG_CONSOLE_OUT=true       # 開發環境建議開啟
SERVICE_NAME=go-template   # 服務名稱
```

**重要：**

*   請勿將 `.env` 檔案提交到版本控制系統中。
*   請務必將 `JWT_SECRET` 修改成一個高強度的密鑰，並定期更新。
*   根據你的實際環境修改資料庫和其他配置。

### 資料庫Migrate

專案目前使用 GORM 的 `AutoMigrate` 功能進行資料庫遷移。**當你需要更新資料庫結構時 (例如新增表格或欄位)**，請執行以下步驟：

1. 修改 `internal/models` 中的資料模型。
2. 執行 `go run migrations/migrate.go`。

**警告：** `AutoMigrate` 在生產環境中有一定的風險，建議在生產環境中使用專業的資料庫遷移工具，例如 `golang-migrate/migrate`。

### 執行 Wire

每次修改 `cmd/go-template/wire.go` 後，都需要執行以下指令重新產生 `cmd/go-template/wire_gen.go`：
先 `cd cmd/go-template` 再執行下面指令

```bash
wire gen
```

### 啟動伺服器

1. **設定環境變數：** 根據你的資料庫設定和 JWT 配置，在 `.env` 檔案中設定對應的環境變數。可以參考 `.env.example`。
2. **建立 `logs` 目錄:**  在專案根目錄下建立 `logs` 目錄，用於儲存日誌檔案, 如果沒有此目錄會導致程式無法正常執行。
3. **安裝依賴：** 在專案根目錄下執行 `go mod tidy`。
4. **執行 Wire：** `cd cmd/go-template`，執行 `go generate ./...` 重新產生 `wire_gen.go` 檔案。
5. **執行資料庫遷移：** 執行 `go run migrations/migrate.go`。
6. **啟動伺服器：** 執行 `go run cmd/go-template/main.go`。

或使用 Makefile：
- 建置： make build
- 執行： make run
- 執行 Wire: make wire
- 執行資料庫遷移： make migrate
- 清除： make clean

### 開發流程

1. **定義資料模型 (`internal/models`)：**  根據你的業務需求定義資料模型。
2. **定義資料庫操作 (`internal/repositories`)：**  定義資料庫操作的介面和實作。
3. **定義業務邏輯 (`internal/services`)：**  定義業務邏輯的介面和實作。
4. **編寫處理函數 (`internal/api/handlers`)：**  編寫 HTTP 請求的處理函數。
5. **定義路由 (`internal/api/routes`)：**  將 HTTP 請求路由到對應的處理函數。
6. **加入日誌：**  在程式碼中加入適當的日誌記錄，方便除錯和監控。
7. **編寫測試：**  編寫單元測試和整合測試，確保程式碼的正確性。

### 新增功能

1. **在 `internal/api/handlers` 中新增 controller (e.g., `product_handler.go`)。**
2. **在 `internal/api/routes` 中新增路由 (e.g., `product_routes.go`)。**
3. **在 `internal/models` 中新增 model (e.g., `product.go`)。**
4. **在 `internal/repositories` 中新增 repository (e.g., `product_repository.go`)。**
5. **在 `internal/services` 中新增 service (e.g., `product_service.go` 和 `impl/product_service_impl.go`)。**
6. **在 `cmd/go-template/wire.go` 中加入新的依賴項。**
7. **執行 `make genegare` 或在 `cmd/go-template/` 目錄下執行 `./...` 重新產生 `cmd/go-template/wire_gen.go`。**
8. **在 `internal/server/server.go` 中的 `NewServer` 函數加入相關依賴 (如果需要)。**

### 日誌記錄

*   使用 `logutil.Logger` 記錄日誌。
*   支援的日誌等級：`debug`、`info`、`warn`、`error`、`dpanic`、`panic`、`fatal`。
*   預設日誌等級為 `info`，可以通過環境變數 `LOG_LEVEL` 修改。
*   日誌會同時輸出到控制台和檔案 (如果 `LOG_CONSOLE_OUT` 設定為 `true`)。
*   日誌檔案位於 `logs/app.log` (預設路徑)，可以通過環境變數 `LOG_FILE` 修改。

### 錯誤處理

*   錯誤碼和錯誤訊息定義在 `internal/api/handlers/errors.go` 中。
*   使用 `response.ErrorResponse` 函數回覆錯誤訊息。

### 身份驗證

*   使用 JWT 進行身份驗證。
*   `AuthMiddleware` 中介軟體會驗證 `Authorization` header 中的 JWT token。
*   受保護的路由 (例如 `/users/:id` 的 GET, PUT, DELETE) 需要使用者提供有效的 JWT token 才能訪問。
*   **無 token 或 token 無效會返回 401 Unauthorized 錯誤。**

### 路由

*   `/users/register`: 使用者註冊 (POST)
*   `/users/login`: 使用者登入 (POST)
*   `/users/:id`: 取得、更新、刪除使用者資訊 (GET, PUT, DELETE) - 需要身份驗證

### JWT 密鑰輪換

*   可以透過設定 `JWT_OLD_SECRETS` 環境變數 (逗號分隔的多個舊密鑰) 來支援 JWT 密鑰輪換。
*   更新 `JWT_SECRET` 後，舊的 JWT 在過期之前仍然有效。
*   當所有舊的 JWT 都過期後，可以從 `JWT_OLD_SECRETS` 中移除舊的密鑰。

### 程式碼風格

*   請遵循 Go 語言的程式碼風格規範。

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

## 安裝 Go 環境

### GVM (Go Version Manager)
> 這邊先打預防針，如果是使用 `windows` 的朋友們，在使用請做好 `$PATH` 會被汙染的準備。
> 主要會影響的範圍以 `wsl` auto import windows path 的部分為主，其餘透過 `export` 去進行設定的環境變數比較不會被影響到。

僅有 `linux` 和 `macos` 的使用者可以參考以下步驟安裝 `gvm`：
```bash
gvm install go1.4
gvm use go1.4
export GOROOT_BOOTSTRAP=$GOROOT
#go1.17.13
gvm use go1.17.13
export GOROOT_BOOTSTRAP=$GOROOT
#go1.20.6
gvm use go1.20.6
export GOROOT_BOOTSTRAP=$GOROOT
#go1.23.4 不一定要這個版本，只是選用了當下的最新版
gvm use go1.23.4 --default # 不然可能會常態性的報錯
```

**注意**:
使用 `wsl` 進行開發的話可能會有錯誤:
![image.png](https://i.imgur.com/6oc1aN4.png)
`gvm` 的安裝檔/script貌似有問題，需要手動在 `.bashrc` / `.zshrc` 中進行配置:
`export PATH=$(echo $PATH | sed 's/:\([^/]\)/ \1/g')`
~~請加在最後面，不然可能會遇到未知錯誤?~~

known issue: [GVM issue #445](https://github.com/moovweb/gvm/issues/445)

如果要使用 `wsl` 的配置的話，建議參考 [wsl config](https://learn.microsoft.com/zh-tw/windows/wsl/wsl-config) 的官方文件把 `appendWindowsPath` 給關閉比較實在，並在 `shell` 的設定檔中去添加自己的環境變數:

```sh
export PATH="你所需要的路徑:$PATH"
```

