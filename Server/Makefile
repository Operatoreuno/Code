# Sqlc commands
.PHONY: sqlc-generate
sqlc-generate:
	@echo "Generating sqlc code..."
	cd db/sqlc && ~/go/bin/sqlc generate

# Development
.PHONY: dev
dev:
	go run main.go

.PHONY: build
build:
	go build -o bin/server

# Test
.PHONY: test
test:
	go test ./...

# Help
.PHONY: help
help:
	@echo "Server Commands:"
	@echo ""
	@echo "  make dev              - Run development server"
	@echo "  make build            - Build production binary"
	@echo "  make test             - Run tests"
	@echo "  make sqlc-generate    - Generate sqlc code from queries"
	@echo ""
	@echo "Database migrations are managed separately:"
	@echo "  cd ../Database && make help"
