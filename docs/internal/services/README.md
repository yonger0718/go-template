# 服務層 (Services)

此目錄包含應用程式的業務邏輯。服務層負責協調不同領域的操作，並透過介面與其他層互動。

## 資料夾結構

主要通常會有兩種選擇

### 方案一 (適用於較大型專案)

當專案規模較大，服務數量眾多時，建議使用以下結構：

```
internal/services/
├── interfaces/   # 存放服務介面定義
│   ├── user.go
│   └── ...
└── implementations/  # 存放服務的具體實作
    ├── user.go
    └── ...
```

-   `interfaces/`: 存放所有服務的介面定義。每個介面定義一個檔案，例如 `user.go` 定義 `UserService` 介面。
-   `implementations/`: 存放服務介面的具體實作。每個實作可以有多個版本，例如 `user.go` 可以是預設實作，`user_cached.go` 可以是帶有快取的實作。

這種結構的好處是可以清楚地區分介面和實作，方便管理和擴展。

### 方案二 (目前採用的方案)

對於較小的專案，或者服務數量不多時，可以使用以下結構：

```
internal/services/
└── user/
    ├── user.go          # 介面定義
    └── user_default.go  # 預設實作
```

-   `user.go`: 存放 `UserService` 介面定義。
-   `user_default.go`: 存放 `UserService` 的預設實作。

這種結構比較簡單，適合快速開發和小型專案。

## 依賴注入

服務層的依賴應該透過建構子注入，並使用介面作為依賴類型。這樣可以方便測試和替換不同的實作。

```go
// 介面定義 (user/user.go)
type UserService interface {
    CreateUser(user *models.User) error
    // ... 其他方法
}

// 實作 (user/user_default.go)
type userService struct {
    userRepo repository.UserRepository // 使用介面類型
    // ... 其他依賴
}

func NewUserService(userRepo repository.UserRepository, ...) UserService {
    return &userService{userRepo: userRepo, ...}
}
