# server

這個路徑並非官方指定的路徑，而是我們自己解耦合所產生的路徑。

server 路徑下通常也只會放一個 server.go，用意主要就是簡化在 main.go 中建立 server 的程式碼，

## 檔案

- **`server.go`**:  HTTP 伺服器的設定和啟動。

## 說明

- `server.go` 定義了 `Start` 函數，用於建立和設定 HTTP 伺服器。
- `Start` 函數接收一個 `Config` 結構體作為參數，其中包含了資料庫連線、JWT 服務、使用者路由和通用配置。
- 使用 Gin 框架建立一個新的路由器。
- 註冊 Swagger 路由。
- 註冊使用者相關的路由。
- 建立 `http.Server` 實例，並設定位址、處理器、逾時等。

## 範例

- [beego/server/web](https://github.com/beego/beego/tree/master/server/web)

- [etcd/server](https://github.com/etcd-io/etcd/tree/main/server)
