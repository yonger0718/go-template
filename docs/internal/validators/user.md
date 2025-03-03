# Validator

如果有客製化驗證的需求，可以在這邊進行實作。

驗證方法又分為幾類(主要是用下面兩種居多):

- struct級別的驗證: 可以 binding 在 struct 屬性上，用來驗證輸入的資料是否符合預期的格式。
- api級別的驗證: 可以透過 middleware 來驗證 api 參數是否符合預期的格式。

## 檔案

- **`user.go`**: 使用者資料驗證函數。
- **`common.go`**: 通用的驗證函數 (目前為空)。

## 說明

- `validators` 目錄包含用於驗證資料的函數。
- `user.go` 定義了 `ValidateUser` 和 `ValidateNewUser` 函數，用於驗證使用者資料。
- `common.go` 目前是空的，但可以用於存放通用的驗證函數。

## 參考資料

- [Go-playground/validator](https://github.com/go-playground/validator)

> 可以從裡面的 `_examples` 找到範例
