APP_NAME=task-manager
DOCKER_COMPOSE=docker-compose

up:
	$(DOCKER_COMPOSE) up --build

down:
	$(DOCKER_COMPOSE) down -v

logs:
	$(DOCKER_COMPOSE) logs -f api

ps:
	$(DOCKER_COMPOSE) ps

curl-list:
	curl http://localhost:8080/tasks | jq

curl-add:
	curl -X POST http://localhost:8080/tasks \
	  -H "Content-Type: application/json" \
	  -d '{"type":"sleep","payload":"{}"}' | jq

grpc-generate:
	protoc --go_out=. --go-grpc_out=. proto/task.proto