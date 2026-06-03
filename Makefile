include .env
export


PROJECT_ROOT = $(pwd)
POSTGRES_URL := 'postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@todoapp-postgres:5432/$(POSTGRES_DB)?sslmode=disable'
MIGRATIONS_PATH := ./migrations


env-up:
	@docker compose up -d todoapp-postgres
env-down:
	@docker compose down todoapp-postgres

env-cleanup:
	@read -p "Delete all environment files? Data loss risk. [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down todoapp-postgres port-forwarder && \
		rm -rf ${PROJECT_ROOT}/out/pgdata && \
		echo "Postgres environment files cleaned"; \
	else \
		echo "Postgres environment cleanup cancelled"; \
	fi

env-port-forward:
	@docker compose up -d port-forwarder

env-port-close:
	@docker compose down port-forwarder

.PHONY: migrate-create
migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "No seq entered"; \
		exit 1; \
	fi; \
	docker compose run --rm todoapp-postgres-migrate \
		create \
		-ext sql \
		-dir $(MIGRATIONS_PATH) \
		-seq "$(seq)"

.PHONY: migrate-down
migrate-down:
	@make migrate-action action=down

.PHONY: migrate-up
migrate-up:
	@make migrate-action action=up

.PHONY: migrate-action		
migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "No action entered"; \
		exit 1; \
	fi; \
	docker compose run --rm todoapp-postgres-migrate \
		-path $(MIGRATIONS_PATH) \
		-database $(POSTGRES_URL) \
		"$(action)"

.PHONY: migrate-force
migrate-force:
	docker compose run --rm todoapp-postgres-migrate \
		force \
		-database $(POSTGRES_URL) \
		-path $(MIGRATIONS_PATH) \
		"$(VERSION)"

logs-cleanup:
	@read -p "Delete all logs files? Data loss risk. [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		rm -rf ${PROJECT_ROOT}/out/logs  && \
		echo "Logs files cleaned"; \
	else \
		echo "Logs cleanup cancelled"; \
	fi

todoapp-run:
	@export LOGGER_FOLDER=${PROJECT_ROOT}/out/logs/ && \
	export POSTGRES_HOST=LOCALHOST && \
	go mod tidy && \
	go run ${PROJECT_ROOT}/cmd/todoapp/main.go	

todoapp-deploy:
	docker compose up -d --build todoapp

ps:
	@docker compose ps