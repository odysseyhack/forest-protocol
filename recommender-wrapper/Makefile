api-start:
	go run ./cmd/main.go

build:
	env GOOS=linux GOARCH=amd64 go build -o bin/api cmd/main.go

deploy:
	docker build --tag registry.unchain.io/forest/recommender-api .
	docker push registry.unchain.io/forest/recommender-api
	kubectl apply -f deployment/deployment.yaml
	kubectl apply -f deployment/ingress.yaml
	kubectl apply -f deployment/service.yaml

