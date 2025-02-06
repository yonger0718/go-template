# config

這個目錄下應存放一些設定用的檔案，主要應該是存放那些設定的參數用的 `.go` 程式和 `.ini`、`confd.toml` 之類的設定檔 。

因為 `go` 程式可以直接讀取 `.env`，且已經透過 `godotenv/autoload` 去進行，所以不會有一個負責讀取 `env` 類型檔案的程式。

## 檔案

*   **`config.go`**:  用於載入和解析配置。

## 說明

*   `config.go` 使用 `github.com/joho/godotenv` 庫從 `.env` 檔案中載入環境變數。
*   `LoadConfig` 函數用於載入配置，並將配置儲存到 `Config` 結構體中。
*   `Config` 結構體中包含了應用程式的所有配置項。

## example

* [gogs/conf](https://github.com/gogs/gogs/tree/main/conf)
* [moby/runconfig](https://github.com/moby/moby/tree/master/runconfig)
* [hugo/config](https://github.com/gohugoio/hugo/tree/master/config)
* [micro/internal/config](https://github.com/micro/micro/tree/master/internal/config)
* [daytona/cmd/daytona/config](https://github.com/daytonaio/daytona/tree/main/cmd/daytona/config)