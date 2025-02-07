# docs

這個目錄應該放置各種說明文件，如 API 文件、使用手冊、操作說明等。

## gin-swagger

目前採用 auto-gen 的模式來進行，可以使用 `make swag` 直接快速生成 API 文件。

或是透過
```bash
# 如果尚未安裝 swag
go install github.com/swaggo/swag/cmd/swag@latest

swag init -g cmd/go-template/main.go -o ./docs --parseDependency --parseInternal
```

以目前的架構來說，即可以在 `http://localhost:8080/swagger/index.html` 路徑找到 API 文件。

已知issue:
according to [this issue](https://github.com/swaggo/gin-swagger/issues/90),
目前需要手動在驗證的 Authorization header 加上 `Bearer ` 才能正常執行驗證的流程
![Bearer header](https://i.imgur.com/TnUg371.png)


## example

> 其實這邊只要大專案基本都會有，可以直接去找大專案

* [nps/docs](https://github.com/ehang-io/nps/tree/master/docs)

* [fiber/docs](https://github.com/gofiber/fiber/tree/main/docs)

* [minio/docs](https://github.com/minio/minio/tree/master/docs)