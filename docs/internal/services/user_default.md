# user_default.go

`user_default.go` 實現了 `UserService` 介面，提供了使用者相關的業務邏輯的具體實作。

## 說明

*   `NewUserService` 函數用於建立 `userService` 結構體的實例，並注入 `repository.UserRepository` 和 `jwt.Service` 的依賴。
*   `userService` 結構體包含了 `repository.UserRepository` 和 `jwt.Service` 的實例。
