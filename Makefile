SHELL = sh
.EXPORT_ALL_VARIABLES:
.ONESHELL:


default: help

help: ## Show this help
	@printf "\033[33m%s:\033[0m\n" 'Available commands'
	@awk 'BEGIN {FS = ":.*?## "} /^[\a-zA-Z0-9_-]+:.*?## / {printf "  \033[32m%-18s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

db-migrate-create: ## создать миграцию make db-migrate-create name=<name>
	migrate create -ext sql -dir migrations $(name)

db-migrate: ## Накатить миграции
	migrate -path migrations -database "postgres://postgres:password@localhost:5432/go_shorty_db?sslmode=disable" up

db-migrate-down: ## Откатить последнюю миграцию
	migrate -path migrations -database "postgres://postgres:password@localhost:5432/go_shorty_db?sslmode=disable" down 1

db-status:
	migrate -path migrations -database "postgres://postgres:password@localhost:5432/go_shorty_db?sslmode=disable" version
