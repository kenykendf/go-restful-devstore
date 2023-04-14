help: ## You are here! showing all command documenentation.
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

#===================#
#== Env Variables ==#
#===================#
DOCKER_COMPOSE_FILE ?= docker-compose.yml


#========================#
#== DATABASE MIGRATION ==#
#========================#

migrate-up: ## Run migrations UP
	docker compose -f ${DOCKER_COMPOSE_FILE} --profile tools run --rm migrate up

migrate-down: ## Rollback migrations, latest migration (1)
	docker compose -f ${DOCKER_COMPOSE_FILE} --profile tools run --rm migrate down 1

migrate-down-all: ## Rollback migrations, all migrations
	docker compose -f ${DOCKER_COMPOSE_FILE} --profile tools run --rm migrate down 1

migrate-create: ## Create a DB migration files e.g `make migrate-create name=migration-name`
	docker compose -f ${DOCKER_COMPOSE_FILE} --profile tools run --rm migrate create -ext sql -dir /migrations $(name)

shell-db: ## Enter to database console
	docker compose -f ${DOCKER_COMPOSE_FILE} exec db psql -U postgres -d postgres

environment: ## Setup environment.
environment:
	docker compose -f ${DOCKER_COMPOSE_FILE} up -d
	
docker-restart: ## Run docker container
	docker compose down && docker compose up -d
docker-logs: ## Logging docker container
	docker logs -f studydso_db
run-app: ## run main.go
	go run cmd/main.go

lint: ## Running golangci-lint
	docker compose -f ${DOCKER_COMPOSE_FILE} --profile tools run --rm lint golangci-lint run -v