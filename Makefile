include .env
export

export PROJECT_ROOT=${shell pwd}

env-up:
	@docker compose up -d bookShop-postgres

env-down:
	@docker compose down bookShop-postgres

env-cleanup:
	@read -p "Очистить все volume файлы? [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down bookShop-postgres port-forwarder minio && \
		sudo rm -rf "${PROJECT_ROOT}/out/pgdata" && \
		echo "Очищено"; \
	else \
		echo "отмена"; \
	fi

env-port-forward:
	@docker compose up -d port-forwarder

env-port-close:
	@docker compose down port-forwarder

env-minio-up:
	@docker compose up -d minio

env-minio-down:
	@docker compose down minio

migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "Отсутсвует параметр seq"; \
		exit 1; \
	fi; \
	docker compose run --rm bookShop-postgres-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"

migrate-up:
	@make migrate-action action=up

migrate-down:
	@make migrate-action action=down

migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "Отсутсвует параметр action"; \
		exit 1; \
	fi; \
	docker compose run --rm bookShop-postgres-migrate \
		-path /migrations \
		-database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@bookShop-postgres:5432/${POSTGRES_DB}?sslmode=disable" \
		"$(action)"

logs-cleanup:
	@read -p "Очистить все log файлы? [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		sudo rm -rf "${PROJECT_ROOT}/out/logs" && \
		echo "Очищено"; \
	else \
		echo "Отмена"; \
	fi

bookshop-run:
	@export "LOGGER_FOLDER=${PROJECT_ROOT}/out/logs" && \
	export POSTGRES_HOST=$$(docker inspect bookShop-env-postgres --format '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}') && \
	echo "POSTGRES_HOST=$$POSTGRES_HOST" && \
	go mod tidy && \
	go run "${PROJECT_ROOT}/cmd/bookshop/main.go"