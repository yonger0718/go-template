# util

> 非官方指定路徑。

用意是解決自己內部重複的程式碼，提取成公用函式，方便重複使用。

> 或者是可以使用 `internal` 路徑，這個是官方比較推崇的用法，但實戰上使用 `util` 路徑也沒有太大問題。

## 子目錄

*   **`database/`**: 資料庫連線相關的函數。
*   **`jwt/`**: JWT 產生和驗證相關的函數。
*   **`logutil/`**: 日誌相關的函數。
*   **`response/`**: HTTP 回應的輔助函數。

## 說明

*   `utils` 目錄中的程式碼應該儘可能地通用和可重用。
*   `utils` 目錄中的程式碼不應該依賴於應用程式的特定業務邏輯。

## example

* [prometheus/util](https://github.com/prometheus/prometheus/tree/main/util)

> 以底下的案例來說，他的 internal 路徑下的 folder 基本上就是各種的工具函式集 `util` 
* [gogs/internal](https://github.com/gogs/gogs/tree/main/internal)

* [moby/internal](https://github.com/moby/moby/tree/master/internal)

待確認: 
https://github.com/daytonaio/daytona/blob/main/pkg/db/target_store.go
https://github.com/880831ian/go-restful-api-repository-messageboard/blob/master/message_board/repository/Repository.go
https://github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/blob/master/internal/helpers/auth.go
https://github.com/dhanibaharzah/golang-auth-jwt/blob/main/app/controllers/auth/authController.go
https://www.youtube.com/watch?v=lf_kiH_NPvM