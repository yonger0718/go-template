# Middleware

如果快速以 `Java Spring` 的概念來講，其實就是一個 `aop` 的概念，可以透過 interceptor 來進行攔截、修改、增強 request、response 的流程。

比較多實際用途會是用於驗證、記錄、追蹤、統計、監控等等。

## 檔案

*   **`auth_middleware.go`**: 身份驗證中介軟體。

## 說明

*   `auth_middleware.go` 驗證 JWT token，並將使用者 ID 儲存到 `gin.Context` 中。
*   中介軟體可以用於在處理 HTTP 請求之前或之後執行一些通用邏輯，例如身份驗證、日誌記錄、錯誤處理等。

## example

* [fiber/middleware](https://github.com/gofiber/fiber/tree/master/middleware)

* [iris/middleware](https://github.com/kataras/iris/tree/main/middleware)