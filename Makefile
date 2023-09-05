GOOSE_DRIVER := "postgres"
GOOSE_DB_STRING := "host=localhost port=5432 user=postgres password=postgres dbname=webapp"

TAG := webapp:1.0.0

run:
	./start.sh

build:
	@go build ./cmd/app/main.go

docker-build:
	@docker build -t $(TAG) deployment/

generate-swagger:
	@mkdir -p ./internal/generated
	@swagger generate server -f ./api/swagger/spec.yml -t ./internal/generated --exclude-main

generate-proto:
	@protoc --go_out=./pkg --go-grpc_out=./pkg ./api/proto/*.proto

check-connections:
	@docker-compose exec postgres psql -U postgres -d webapp -c 'SELECT pid, usename, state, query FROM pg_stat_activity WHERE state IS NOT NULL;'

create-user-table:
	@docker-compose exec postgres psql -U postgres -d webapp -c 'CREATE TABLE users (id SERIAL NOT NULL, created_at DATETIME);'

migrate-up:
	@goose -dir ./migrations ${GOOSE_DRIVER} ${GOOSE_DB_STRING} up

migrate-down:
	@goose -dir ./migrations ${GOOSE_DRIVER} ${GOOSE_DB_STRING} down

show-user-table:
	@docker-compose exec postgres psql -U postgres -d webapp -c '\d+ user;'

show-order-table:
	@docker-compose exec postgres psql -U postgres -d webapp -c '\d+ order;'