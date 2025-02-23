# Route

這個路徑並非官方指定的路徑，而是我們自己解耦合所產生的路徑。

主要可以用類似 controller 的方式來處理路由。

## 路由列表

### 使用者相關路由 (/api/user)

| 方法   | 路徑       | 說明         | 身份驗證 |
| ------ | -------- | ------------ | -------- |
| POST   | /register | 註冊使用者     | 否       |
| POST   | /login    | 使用者登入     | 否       |
| GET    | /:id      | 取得使用者資訊 | 是       |
| PUT    | /:id      | 更新使用者資訊 | 是       |
| DELETE | /:id      | 刪除使用者     | 是       |

## 範例

* [prometheus/common/route](https://github.com/prometheus/common/tree/main/route)

> 這個案例是另一種模式，整合了路由、控制器等模組直接放在web包底下，其實他的模式就是直接採用MVC的架構下去實作
* [nps/web/routers](https://github.com/ehang-io/nps/tree/master/web/routers)

* [gogs/internal/route](https://github.com/gogs/gogs/tree/main/internal/route)
