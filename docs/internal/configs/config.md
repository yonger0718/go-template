# config

這個目錄下應存放一些設定用的檔案，主要應該是存放那些設定的參數用的 `.go` 程式和 `.ini`、`confd.toml` 之類的設定檔 。

因為 `go` 程式可以直接讀取 `.env`，且已經透過 `godotenv/autoload` 去進行，所以不會有一個負責讀取 `env` 類型檔案的程式。

## 檔案

*   **`config.go`**:  用於載入和解析配置。

## 說明

*   `config.go` 使用 `github.com/joho/godotenv` 庫從 `.env` 檔案中載入環境變數。
*   `LoadConfig` 函數用於載入配置，並回傳一個 `Config` 結構體的指標。
*   檔案中定義了 `getEnv` 和 `getBoolEnv` 輔助函數，用於取得環境變數並提供預設值。

## 範例

* [gogs/conf](https://github.com/gogs/gogs/tree/main/conf)
* [micro/internal/config](https://github.com/micro/micro/tree/master/internal/config)
