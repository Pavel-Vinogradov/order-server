package main

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"net/http"
	"order-server/infrastructure/db"
	authHandler "order-server/internal/auth/handler"
	authRepository "order-server/internal/auth/repository"
	authService "order-server/internal/auth/service"
	"order-server/internal/config"
	"order-server/internal/middleware"
	productHandler "order-server/internal/product/handler"
	productRepository "order-server/internal/product/repository"
	productService "order-server/internal/product/service"
	userRepository "order-server/internal/user/repository"
	userService "order-server/internal/user/service"
	"order-server/pkg/logger"
)

// @title Order API
// @version 1.0
// @description API для управления заказами и продуктами
// @host http://localhost:8080
// @BasePath /
func main() {
	cfg := config.Load()
	logger.Init()
	dbConn, err := db.Init(cfg)
	if err != nil {
		logrus.WithError(err).Fatal("Ошибка подключения к БД: %v", err)
	}
	defer func(dbConn *sql.DB) {
		err := dbConn.Close()
		if err != nil {
			return
		}
	}(dbConn)
	router := http.NewServeMux()

	// Product
	newProductRepository := productRepository.NewProductRepository(dbConn)
	newProductService := productService.NewProductService(newProductRepository)
	productHandler.NewRouterProduct(newProductService, router)

	// User
	newUserRepository := userRepository.NewUserRepository(dbConn)
	newUserService := userService.NewUserService(newUserRepository)
	//userHandler.NewRouterUser(newUserService, router)
	//Auth
	newAuthRepository := authRepository.NewAuthRepository(dbConn)
	newAuthService := authService.NewAuthService(newAuthRepository, newUserService)
	authHandler.NewRouteAuth(newAuthService, router)

	middlewareStack := middleware.Chain(
		middleware.Logging,
		middleware.CORS,
	)
	addr := ":" + cfg.App.Port
	server := http.Server{
		Addr:    addr,
		Handler: middlewareStack(router),
	}

	logrus.Info("Сервер запущен на %s\n", addr)
	err = server.ListenAndServe()
	if err != nil {
		logrus.WithError(err).Fatal("Ошибка запуска сервера: %v", err)
	}
}
