.PHONY: help build run test clean deps ui-deps ui-dev worker client vault-setup

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

deps: ## Install Go dependencies
	go mod download
	go mod tidy

ui-deps: ## Install UI dependencies
	cd ui && npm install

build: deps ## Build Go binaries
	go build -o bin/worker ./worker
	go build -o bin/client ./client

run: build ## Run the worker
	./bin/worker

test: ## Run tests
	go test ./...

clean: ## Clean build artifacts
	rm -rf bin/
	rm -rf ui/dist/
	rm -rf ui/node_modules/

ui-dev: ui-deps ## Start UI development server
	cd ui && npm run dev

worker: ## Run worker in development
	go run worker/main.go

client: ## Run client in development
	go run client/main.go --tenant=default --lead=test123

docker-up: ## Start all services with Docker Compose
	docker-compose up -d

docker-down: ## Stop all services
	docker-compose down

docker-logs: ## View logs from all services
	docker-compose logs -f

vault-setup: ## Setup Vault with initial secrets
	@echo "Setting up Vault secrets..."
	@curl -X POST -H "X-Vault-Token: dev-token" \
		-H "Content-Type: application/json" \
		-d '{"data":{"google_ads_key":"your-google-ads-key"}}' \
		http://localhost:8200/v1/secret/data/adflowhq/google-ads
	@curl -X POST -H "X-Vault-Token: dev-token" \
		-H "Content-Type: application/json" \
		-d '{"data":{"facebook_ads_key":"your-facebook-ads-key"}}' \
		http://localhost:8200/v1/secret/data/adflowhq/facebook-ads
	@curl -X POST -H "X-Vault-Token: dev-token" \
		-H "Content-Type: application/json" \
		-d '{"data":{"api_key":"your-callrail-key"}}' \
		http://localhost:8200/v1/secret/data/adflowhq/callrail
	@echo "Vault setup complete!"

dev: ## Start development environment
	@echo "Starting development environment..."
	@make docker-up
	@sleep 10
	@make vault-setup
	@echo "Development environment ready!"
	@echo "Temporal UI: http://localhost:8080"
	@echo "Vault UI: http://localhost:8200"
	@echo "React UI: http://localhost:3000" 