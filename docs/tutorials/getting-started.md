## 快速開始教學

1.  **取得程式碼**：

    ```bash
    git clone https://github.com/xxx/go-template.git # 替換成你的專案網址
    cd go-template
    ```

2.  **安裝依賴**：

    ```bash
    go mod download
    ```

3.  **設定環境變數**：

    複製 `.env.example` 檔案為 `.env`，並根據您的環境修改其中的設定。

4.  **執行應用程式**：

    ```bash
    make run
    ```

5.  **產生 Swagger 文件** (可選)：

    ```bash
    make swag
    ```

    然後在瀏覽器中開啟 `http://localhost:8080/swagger/index.html` 查看 API 文件。

## 前置要求

*   **Go:**  版本 1.18 或更高版本。
*   **PostgreSQL:**  用於資料庫。
*   **Git:**  用於版本控制。

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

## 安裝及更新專案依賴

你可以使用以下 Makefile 指令來安裝或更新專案所需的依賴：

```bash
make deps  # 安裝專案所需的工具 (golangci-lint, gosec, govulncheck, wire, swag, editorconfig-checker)
make tidy  # 整理 go.mod 和 go.sum 檔案
```

`make deps` 會安裝以下工具：

*   **golangci-lint:** Go 程式碼檢查工具。
*   **gosec:** Go 程式碼安全性掃描工具。
*   **govulncheck:** Go 漏洞檢查工具。
*   **wire:** Go 依賴注入工具。
*   **swag:** Swagger 文件產生工具。
*   **editorconfig-checker:** EditorConfig 檔案檢查工具。

`make tidy` 會根據 `go.mod` 檔案中的定義，更新 `go.sum` 檔案，並移除未使用的依賴。

## 環境變數設定

請根據你的環境，在 `.env` 檔案中設定以下環境變數：
> 如果有需要，可以參考 `.env.example` 檔案作為參考。

```
DB_HOST=
DB_PORT=
DB_USERNAME=
DB_PASSWORD=
DB_DATABASE=
JWT_SECRET=                # 請務必修改成高強度密鑰，並定期更新
JWT_OLD_SECRETS=           # 舊的 JWT 密鑰，用於支援密鑰輪換 (可選, 多組密鑰使用逗號分隔)
TOKEN_EXPIRES_IN=24        # 單位: 小時
PORT=8080
LOG_LEVEL=info             # 日誌等級: debug, info, warn, error, dpanic, panic, fatal
LOG_FILE=logs/app          # 日誌檔案路徑以及前綴
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

## 資料庫Migrate

專案目前使用 GORM 的 `AutoMigrate` 功能進行資料庫遷移。**當你需要更新資料庫結構時 (例如新增表格或欄位)**，請執行以下步驟：

1.  修改 `internal/models` 中的資料模型。
2.  執行 `go run migrations/migrate.go`。

**警告：** `AutoMigrate` 在生產環境中有一定的風險，建議在生產環境中使用專業的資料庫遷移工具，例如 `golang-migrate/migrate`。

## 執行 Wire

每次修改 `cmd/go-template/wire.go` 後，都需要執行以下指令重新產生 `cmd/go-template/wire_gen.go`：
先 `cd cmd/go-template` 再執行下面指令

```bash
# 雖然這邊都使用 ./... 但也可以針對特定的目標，即 ./wire.go 去生成目標
# 直接使用wire 或可以使用 go wire
wire gen ./cmd/go-template/...

# 或是使用 go generate
go generate ./cmd/go-template/...

# 簡易版 使用make
make generate
```

## 啟動伺服器

1.  **設定環境變數：** 根據你的資料庫設定和 JWT 配置，在 `.env` 檔案中設定對應的環境變數。可以參考 `.env.example`。
2.  **建立 `logs` 目錄:**  在專案根目錄下建立 `logs` 目錄，用於儲存日誌檔案, 如果沒有此目錄會導致程式無法正常執行。
3.  **安裝依賴：** 在專案根目錄下執行 `go mod tidy`。
4.  **執行 Wire：** `cd cmd/go-template`，執行 `go generate ./...` 重新產生 `wire_gen.go` 檔案。
5.  **執行資料庫遷移：** 執行 `go run migrations/migrate.go`。
6.  **啟動伺服器：** 執行 `go run cmd/go-template/main.go`。

或使用 Makefile：

- 建置： make build
- 執行： make run
- 執行 Wire: make wire
- 執行資料庫遷移： make migrate
- 清除： make clean
