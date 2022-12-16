DIR := ${CURDIR}

# sql to go code
sqlc:
	docker run --rm -v "${DIR}/internal/repository/sqlc:/src" -w /src kjconroy/sqlc generate

go-dev:
	docker compose -f ./compose.yaml -f ./docker-compose.dev.yaml up --build 