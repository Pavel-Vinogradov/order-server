package main

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"net/http"
	"order-server/infrastructure/db"
	"order-server/internal/config"
	"order-server/internal/middleware"
	"order-server/internal/product/handler"
	"order-server/internal/product/repository"
	"order-server/internal/product/service"
	"order-server/pkg/logger"
)

// @title Order API
// @version 1.0
// @description API для управления заказами и продуктами
// @host localhost:8080
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

	repo := repository.NewProductRepository(dbConn)
	productService := service.NewProductService(repo)

	router := handler.NewRouterProduct(productService)

	middlewareStack := middleware.Chain(
		middleware.Logging,
	)
	finalHandler := middlewareStack(router)
	addr := ":" + cfg.App.Port
	logrus.Info("Сервер запущен на %s\n", addr)
	err = http.ListenAndServe(addr, finalHandler)
	if err != nil {
		logrus.WithError(err).Fatal("Ошибка запуска сервера: %v", err)
	}
}
