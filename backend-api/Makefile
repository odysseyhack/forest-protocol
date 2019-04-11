start: mongo-start api-start

api-start:
	go run ./cmd/backend-api/main.go

mongo-start:
	docker-compose up -d mongo

mongo-stop:
	docker-compose stop mongo