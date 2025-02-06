# Validator

如果有客製化驗證的需求，可以在這邊進行實作。

驗證方法又分為幾類(主要是用下面兩種居多):
* struct級別的驗證: 可以 binding 在 struct 屬性上，用來驗證輸入的資料是否符合預期的格式。
* api級別的驗證: 可以透過 middleware 來驗證 api 參數是否符合預期的格式。
## 檔案

*   **`user_validator.go`**: 使用者資料驗證函數。

## 說明

*   `user_validator.go` 中的 `ValidateUser` 函數用於驗證使用者註冊和更新時提供的資料是否合法。

## 參考資料

* [Go-playground/validator](https://github.com/go-playground/validator)
> 可以從裡面的 `_examples` 找到範例