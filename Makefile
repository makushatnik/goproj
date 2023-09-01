GOOSE_DRIVER := "postgres"
GOOSE_DB_STRING := "host=localhost port=5432 user=postgres password=postgres dbname=webapp"

run:
	./start.sh

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