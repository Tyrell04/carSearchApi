db:
	docker compose -f dev.docker-compose.yml up -d
dev:
	go run cmd/main.go