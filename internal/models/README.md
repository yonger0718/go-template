# model

主要存放 ORM 的模型定義，主要是用來對應資料庫的表格，主要也是對應本專案所選用的 `GORM` 框架。

## 檔案

*   **`user.go`**: 使用者資料模型。

## 說明

*   資料模型使用 GORM 的標籤 (tag) 來定義資料庫表格的結構。
*   `json` 標籤用於控制 JSON 的序列化和反序列化。

## example

> 這個不大算是標準版，不過這個看起來內容大多數是對的，但我認為裡面的內容應該把db相關的連線建立給獨立到 `util` 或是 `server` 層底下
* [gogs/internal/database](https://github.com/gogs/gogs/tree/main/internal/database)

* [prometheus/model](https://github.com/prometheus/prometheus/tree/main/model)

* [prometheus/common/model](https://github.com/prometheus/common/tree/main/model)