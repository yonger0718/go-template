//go:build wireinject
// +build wireinject

package main

import (
	userHandler "go-template/internal/api/handlers/user"
	"go-template/internal/repository"
	"net/http"

	"github.com/google/wire"
	"go-template/internal/api/handlers/routes"
	"go-template/internal/configs"
	"go-template/internal/server"
	userSvc "go-template/internal/services/user"
	"go-template/internal/utils/database"
	"go-template/internal/utils/jwt"
)

// InitializeServer 使用 Wire 進行依賴注入，初始化 HTTP server
func InitializeServer(cfg *configs.Config) (*http.Server, func(), error) {
	wire.Build(
		// 依序綁定各個依賴項
		database.Start,
		repository.NewUserRepository,
		jwt.NewService,
		userSvc.NewUserService,
		userHandler.NewHandler,
		routes.NewUser,
		server.Start,
		// 將多個依賴項組合成 ServerConfig 結構體
		wire.Struct(new(server.Config), "*"),
	)
	return &http.Server{}, func() {
		// 保持空即可
	}, nil
}
