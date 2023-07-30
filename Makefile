db:
	docker-compose -f dev.docker-compose.yml up -d
dev:
	go run cmd/main.go
db-down:
	docker-compose -f dev.docker-compose.yml down
db-rm:
	docker-compose -f dev.docker-compose.yml down
	docker volume rm carsearchapi_postgres_data