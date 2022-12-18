DIR := ${CURDIR}

# sql to go code
sqlc:
	docker run --rm -v "${DIR}/internal/repository/sqlc:/src" -w /src kjconroy/sqlc generate

# run project
go-dev:
	docker compose -f ./compose.yaml -f ./docker-compose.dev.yaml up --build 

# generate protoc files
protoc: 
	protoc --go_out=. --go-grpc_out=. ./proto/RoutesAnalysis.proto