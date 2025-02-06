# cmd

此目錄包含應用程式的主程式入口點 (main packages)。

每個應用程式的目錄名應該與你的執行檔案名稱一致 (例如：/cmd/myapp, /cmd/go-template)。

這個目錄下不應該有過多的程式碼，若有這類的需求可以考慮新增一個 `pkg` 目錄來存放這些共用的程式碼。

## 說明

*   `cmd` 目錄下的每個子目錄都應該對應一個獨立的可執行檔案。
*   `main.go` 檔案應該儘可能簡潔，主要負責：
    *   載入配置
    *   初始化日誌
    *   使用 Wire 進行依賴注入
    *   啟動伺服器
*   避免在 `main.go` 中編寫過多的業務邏輯。
*   每個 `main.go` 檔案都屬於 `main` package。

## example

* [kubernetes/cmd](https://github.com/kubernetes/kubernetes/tree/master/cmd)

> 要參考感覺這個會比較實在，k8s的老實說偏大
* [caddy/cmd](https://github.com/caddyserver/caddy/tree/master/cmd)

> 這個專案的模式也算是一種方案，他的layout比較偏向是以類似 `src` 單一路徑的方式在管理
* [minio/cmd](https://github.com/minio/minio/tree/master/cmd)