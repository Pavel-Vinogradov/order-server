package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"order-server/internal/config"
)

func Init(config *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DBConfig.Host,
		config.DBConfig.Port,
		config.DBConfig.User,
		config.DBConfig.Password,
		config.DBConfig.Name,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}
