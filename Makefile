.PHONY: up down

up: ## docker compose up with default option
	docker-compose up -d --build
down: ## docker compose down
	docker-compose down
