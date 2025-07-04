include .env

up:
	@echo "Starting containers..."
	docker-compose up --build -d --remove-orphans

down:
	@echo "Stopping containers..."
	docker-compose down

build:
	@go build -o bin/chat-backend-app cmd/main.go
	@echo "Build completed"

start:
	./bin/chat-backend-app

restart: build start
