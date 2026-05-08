include .env
export

SHELL := C:/Program Files/Git/bin/bash.exe

export PROJECT_ROOT=${CURDIR}

env-up:
	docker compose up -d todoapp-postgres
env-down:
	docker compose down todoapp-postgres

env-cleanup:
	@read -p "Delete all environment files? Data loss risk. [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down todoapp-postgres && \
		rm -rf out/pgdata && \
		echo "Postgres environment files cleaned"; \
	else \
		echo "Postgres environment cleanup cancelled"; \
	fi