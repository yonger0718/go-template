# Go Template 專案文件

這份文件提供了 Go Template 專案的相關資訊，包括專案目標、功能、架構、如何建置和執行專案，以及各個模組的詳細說明。

## 專案概述

Go Template 是一個基於 Go 語言的 Web 應用程式模板，旨在幫助開發者快速啟動新的 Web 專案。它提供了一個基本的專案結構、常用的功能模組和清晰的文件，讓開發者可以專注於業務邏輯的開發。

## 功能

*   使用者註冊和登入
*   JWT 身份驗證
*   資料庫操作 (使用 GORM)
*   日誌記錄 (使用 Zap)
*   設定檔管理
*   API 路由和處理
*   Swagger API 文件

## 專案架構

專案的程式碼主要位於 `internal` 目錄下，按照功能模組劃分為不同的子目錄：

*   [api](./internal/api): API 相關的程式碼，包括路由、處理函數、回應格式等。
    *   [handlers](./internal/api/handlers.md): HTTP 請求處理函數
    *   [routes](./internal/api/routes.md): 路由定義
    *   [response](./internal/api/response.md): 回應格式處理
*   [configs](./internal/configs): 設定檔相關的程式碼。
*   [middleware](./internal/middleware): 中介軟體相關的程式碼，例如身份驗證。
*   [models](./internal/models): 資料模型定義。
*   [repository](./internal/repository): 資料庫操作相關的程式碼。
*   [server](./internal/server): HTTP 伺服器相關的程式碼。
*   [services](./internal/services): 業務邏輯相關的程式碼。
    *   [user](./internal/services/user/user.md): 使用者服務介面
    *   [user_default](./internal/services/user/user_default.md): 使用者服務實作
*   [utils](./internal/utils): 通用工具函數。
    *   [database](./internal/utils/database): 資料庫連線相關的函數。
    *   [jwt](./internal/utils/jwt): JWT 產生和驗證相關的函數。
    *   [logger](./internal/utils/logger): 日誌相關的函數。
*   [validators](./internal/validators): 資料驗證相關的程式碼。

## 快速開始

請參閱 [快速開始教學](./tutorials/getting-started.md) 以取得詳細的專案設定和執行說明。

## 模組文件

更詳細的模組說明，以及開發流程及規範，請參閱以下文件：

*   [cmd/go-template](./cmd/go-template.md): 主要應用程式
*   [internal/api](./internal/api): API 相關文件
    *   [handlers](./internal/api/handlers.md): 請求處理函數文件
    *   [routes](./internal/api/routes.md): 路由定義文件
    *   [response](./internal/api/response.md): 回應格式文件
*   [internal/configs](./internal/configs/config.md): 設定檔文件
*   [internal/middleware](./internal/middleware/auth.md): 中介軟體文件
*   [internal/models](./internal/models/user.md): 資料模型文件
*   [internal/repository](./internal/repository/user.md): 資料庫操作文件
*   [internal/server](./internal/server/server.md): 伺服器文件
*   [internal/services](./internal/services/README.md): 服務層文件
    *   [user](./internal/services/user/user.md): 使用者服務介面文件
    *   [user_default](./internal/services/user/user_default.md): 使用者服務實作文件
*   [internal/utils](./internal/utils): 工具函數文件
    *   [database](./internal/utils/database/database.md): 資料庫連線文件
    *   [jwt](./internal/utils/jwt/jwt.md): JWT 相關文件
    *   [logger](./internal/utils/logger/logger.md): 日誌相關文件
*   [internal/validators](./internal/validators/user.md): 資料驗證文件
