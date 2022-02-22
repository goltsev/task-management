ifneq (${wildcard ./.env},)
	include .env
	export
endif

DOCKER_COMPOSE_FILE := './docker-compose.yml'

run:
	go run ./cmd/task-mgmt

check-utils:
	which swagx || go install github.com/swaggo/swag/cmd/swag@latest
	which mockgen || go install github.com/golang/mock/mockgen@v1.6.0

swagger: check-utils
	swag init -g ./cmd/task-mgmt/main.go

mocks: check-utils
	mockgen -source=./service/interface.go -destination=./service/mock_service/interface.go
	mockgen -source=./handlers/interface.go -destination=./handlers/mock_handlers/interface.go

dev-up:
	docker-compose -f ${DOCKER_COMPOSE_FILE} up -d

dev-down:
	docker-compose -f ${DOCKER_COMPOSE_FILE} down	

prod-up:
	docker-compose -f ${DOCKER_COMPOSE_FILE} --profile prod up -d

prod-down:
	docker-compose -f ${DOCKER_COMPOSE_FILE} --profile prod down

build:
	CGO_ENABLED="0" GOOS="linux" GOARCH="amd64" go build -o ./bin/task-mgmt ./cmd/task-mgmt

docker:
	docker build -t task-mgmt .
	docker run --rm --env-file ./.env task-mgmt

