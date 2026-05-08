include .env
export

export PROJECT_ROOT=${shell pwd}

env-up:
	@docker compose up -d todoapp-postgres
env-down:
	@docker compose down todoapp-postgres

env-cleanup:
	@read -p "Delete all environment files? Data loss risk. [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down todoapp-postgres && \
		rm -rf out/pgdata && \
		echo "Postgres environment files cleaned"; \
	else \
		echo "Postgres environment cleanup cancelled"; \
	fi

migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "No seq entered"; \
		exit 1; \
	fi; \
	docker compose run --rm todoapp-postgres-migrate \
		create \
		-ext sql \
		-dir ./migrations \
		-seq "$(seq)" 

migrate-down:
	@make migrate-action action=down

migrate-up:
	@make migrate-action action=up

migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "No action entered"; \
		exit 1; \
	fi; \
	docker compose run --rm todoapp-postgres-migrate \
		-path /migrations \
		-database postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@todoapp-postgres:5432/$(POSTGRES_DB)&sslmode=disable \
		"$(action)"