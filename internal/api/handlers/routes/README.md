# Route

這個路徑並非官方指定的路徑，而是我們自己解耦合所產生的路徑。

主要可以用類似 controller 的方式來處理路由。

## example

* [prometheus/common/route](https://github.com/prometheus/common/tree/main/route)

> 這個案例是另一種模式，整合了路由、控制器等模組直接放在web包底下，其實他的模式就是直接採用MVC的架構下去實作
* [nps/web/routers](https://github.com/ehang-io/nps/tree/master/web/routers)

* [gogs/internal/route](https://github.com/gogs/gogs/tree/main/internal/route)