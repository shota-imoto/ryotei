.PHONY: up down

up: ## docker compose up with default option
	docker-compose up -d --build
down: ## docker compose down
	docker-compose down
migrate: ## db migrate
	migrate -source 'file://server/db/migrations' -database 'postgresql://postgres:postgres@localhost:5432/trvl?sslmode=disable' up 1
	\ && 	migrate -source 'file://server/db/migrations' -database 'postgresql://postgres:postgres@localhost:5432/trvl?sslmode=disable' up 1
