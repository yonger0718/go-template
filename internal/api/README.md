# API

# internal/api 目錄

此目錄包含與 API 相關的程式碼。

> 有的專案也會拿來存放部分api相關的文件，例如說 Swagger 定義、OpenAPI 定義等等。

## 子目錄

*   **`handlers/`**: HTTP 請求的處理函數 (控制器)。
*   **`routes/`**: 路由定義。
*   **`errors.go`**: 定義了 API 相關的錯誤碼和錯誤訊息。

## 說明

*   `handlers` 目錄中的每個子目錄都應該對應一個資源 (例如 `user`)。
*   `routes` 目錄中定義了 URL 路徑和 HTTP 方法到 `handlers` 中處理函數的映射。
*   `errors.go` 檔案中定義的錯誤碼應該遵循一定的命名規範，例如 `ErrCode<Resource><Error>` (例如 `ErrCodeUserNotFound`)。

## example

* [kubernetes/api](https://github.com/kubernetes/kubernetes/tree/master/api/api-rules)

* [moby/api](https://github.com/moby/moby/tree/master/api)

* [prometheus/web/api/v1](https://github.com/prometheus/prometheus/tree/main/web/api/v1)