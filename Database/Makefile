# Database Configuration
DB_USER ?= dmnc
DB_NAME ?= db
DB_HOST ?= localhost
DB_PORT ?= 5432
DB_URL = postgres://$(DB_USER)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

# Schema and Migration commands
.PHONY: schema
schema:
	@echo "Applying schema..."
	psql -U $(DB_USER) -d $(DB_NAME) -f schema.sql

.PHONY: migrate-up
migrate-up:
	@echo "Running migrations..."
	migrate -path migrations -database "$(DB_URL)" up

.PHONY: migrate-down
migrate-down:
	@echo "Rolling back last migration..."
	migrate -path migrations -database "$(DB_URL)" down 1

.PHONY: migrate-down-all
migrate-down-all:
	@echo "Rolling back all migrations..."
	migrate -path migrations -database "$(DB_URL)" down -all

.PHONY: migrate-force
migrate-force:
	@echo "Forcing migration version..."
	@read -p "Enter version: " version; \
	migrate -path migrations -database "$(DB_URL)" force $$version

.PHONY: migrate-version
migrate-version:
	@echo "Current migration version:"
	migrate -path migrations -database "$(DB_URL)" version

.PHONY: migrate-create
migrate-create:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir migrations -seq $$name

# Help
.PHONY: help
help:
	@echo "Database Management Commands:"
	@echo ""
	@echo "  make schema              - Apply schema.sql to database"
	@echo "  make migrate-up          - Apply all pending migrations"
	@echo "  make migrate-down        - Rollback last migration"
	@echo "  make migrate-down-all    - Rollback all migrations"
	@echo "  make migrate-version     - Show current migration version"
	@echo "  make migrate-create      - Create new migration files"
	@echo "  make migrate-force       - Force migration to specific version"