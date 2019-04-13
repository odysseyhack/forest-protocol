gen-api:
	swagger generate server -f ./api/swagger.yaml -t ./gen -C go-swagger.yaml --server-package api --model-package dto

start: mongo-start api-start

api-start:
	go run ./cmd/backend-api/main.go

mongo-start:
	docker-compose up -d mongo

mongo-stop:
	docker-compose stop mongo

build:
	env GOOS=linux GOARCH=amd64 go build -o bin/api cmd/backend-api/main.go

deploy:
	docker build --tag registry.unchain.io/forest/forest-api .
	docker push registry.unchain.io/forest/forest-api
