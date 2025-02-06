# test

> 其實 `go` 專案的test寫法也很多種，有一類的寫法是直接放在跟原始程式碼同一層，透過命名為 `*_test.go` 的檔案來寫測試程式。統一管理最大的好處是在寫 `ci` 的時候管理起來方便程度會大大的提升。

官方文件: 
>額外的外部測試應用程式和測試資料。你可以自在的調整你在 /test 目錄中的結構。對於較大的專案來說，通常會有一個 data 資料夾也是蠻正常的。例如：如果你需要 Go 忽略這些目錄下的檔案，你可以使用 /test/data 或 /test/testdata 當作你的目錄名稱。請注意：Go 還會忽略以 . 或 _ 開頭的目錄或檔案，所以你在測試資料的目錄命名上，將擁有更大的彈性。

## 說明

*   使用 Go 語言內建的 `testing` 套件編寫測試程式碼。
*   使用 `testify` 或 `go-sqlmock` 等第三方庫可以讓測試程式碼更簡潔易懂。
*   `mocks` 目錄可以用於存放模擬物件，例如使用 `testify/mock` 或 `sqlmock` 產生的模擬物件。

## example

* [minikube/test](https://github.com/kubernetes/minikube/tree/master/test)

* [gitea/test](https://github.com/go-gitea/gitea/tree/main/tests)

* [gogs/internal/testutil](https://github.com/gogs/gogs/tree/main/internal/testutil)