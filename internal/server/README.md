# server

這個路徑並非官方指定的路徑，而是我們自己解耦合所產生的路徑。

server 路徑下通常也只會放一個 server.go，用意主要就是簡化在 main.go 中建立 server 的程式碼，
~~當然這邊的變化也很多，可以實作類似 DI 架構的模式。(為甚麼我要立flag)~~

## 檔案

*   **`server.go`**:  HTTP 伺服器的設定和啟動。

## 說明

*   `server.go` 使用 Gin 框架來建立 HTTP 伺服器。
*   `server.go` 使用 `middleware` 中介軟體來處理身份驗證等通用邏輯。
*   `server.go` 使用 `routes` 來註冊 API 路由。

## example

* [beego/server/web](https://github.com/beego/beego/tree/master/server/web)

* [etcd/server](https://github.com/etcd-io/etcd/tree/main/server)