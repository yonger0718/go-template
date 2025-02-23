# UserService 介面

## 介紹
`UserService` 介面定義了使用者相關的核心業務邏輯，包括使用者註冊、登入、查詢、更新和刪除等操作。

## 方法

### CreateUser
建立新的使用者。

**參數：**

- `user *models.User`: 使用者資訊
  - `Username`: 使用者名稱 (必填)
  - `Password`: 使用者密碼 (必填，將自動進行 bcrypt 加密)
  - `Email`: 電子郵件 (必填)

**返回值：**

- `error`: 可能的錯誤
  - `nil`: 成功
  - 其他: 建立失敗

**使用範例：**

```go
user := &models.User{
    Username: "john_doe",
    Password: "secure_password",
    Email:    "john@example.com",
}
err := userService.CreateUser(user)
if err != nil {
    // 處理錯誤
    log.Printf("建立使用者失敗: %v", err)
}
```

### GetUserByID
根據 ID 取得使用者資訊。

**參數：**

- `id uint`: 使用者 ID

**返回值：**

- `*models.User`: 使用者資訊
- `error`: 可能的錯誤
  - `nil`: 成功
  - `services.ErrUserNotFound`: 使用者不存在

**使用範例：**

```go
user, err := userService.GetUserByID(1)
if err == services.ErrUserNotFound {
    // 處理使用者不存在的情況
} else if err != nil {
    // 處理其他錯誤
}
```

### GetUserByUsername
根據使用者名稱取得使用者資訊。

**參數：**

- `username string`: 使用者名稱

**返回值：**

- `*models.User`: 使用者資訊
- `error`: 可能的錯誤
  - `nil`: 成功
  - `services.ErrUserNotFound`: 使用者不存在

**使用範例：**

```go
user, err := userService.GetUserByUsername("john_doe")
if err == services.ErrUserNotFound {
    // 處理使用者不存在的情況
}
```

### UpdateUser
更新使用者資訊。

**參數：**

- `user *models.User`: 使用者資訊 (包含要更新的欄位)

**返回值：**

- `error`: 可能的錯誤
  - `nil`: 成功
  - 其他: 更新失敗

**使用範例：**

```go
user := &models.User{
    ID:       1,
    Username: "john_updated",
    Email:    "john_updated@example.com",
}
err := userService.UpdateUser(user)
```

### DeleteUser
刪除使用者。

**參數：**

- `id uint`: 使用者 ID

**返回值：**

- `error`: 可能的錯誤
  - `nil`: 成功
  - 其他: 刪除失敗

**使用範例：**

```go
err := userService.DeleteUser(1)
```

### Login
使用者登入並取得 JWT token。

**參數：**

- `username string`: 使用者名稱
- `password string`: 密碼

**返回值：**

- `authToken string`: JWT token
- `err error`: 可能的錯誤
  - `nil`: 登入成功
  - `services.ErrUserNotFound`: 使用者不存在
  - `services.ErrInvalidCredentials`: 密碼錯誤

**使用範例：**

```go
authToken, err := userService.Login("john_doe", "secure_password")
if err == services.ErrInvalidCredentials {
    // 處理密碼錯誤
} else if err != nil {
    // 處理其他錯誤
}
```

## 錯誤

- `ErrUserNotFound`: 使用者不存在。
- `ErrInvalidCredentials`: 無效的憑證 (例如密碼錯誤)。
