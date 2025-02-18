include .env
POSTGRESQL_URL='postgres://$(POSTGRES_USER):${POSTGRES_PASSWORD}@localhost:5433/${POSTGRES_DB}?sslmode=disable'


PHONY: migr_cr
migr_cr:
	migrate create -verbose -ext sql -dir ./migrations -seq create_aux_tables

PHONY: migr_up
migr_up:
	migrate -verbose -database ${POSTGRESQL_URL} -path ./migrations up 
	


PHONY: migr_down
migr_down:
	migrate -verbose -database ${POSTGRESQL_URL} -path ./migrations down

# migrate -database postgres://postgres:crmpassword@localhost:5433/postgres?sslmode=disable -path ./migrations up

#   include .env

#   create_migration:
#     migrate create -ext=sql -dir=internal/database/migrations -seq init

#   migrate_up:
#     migrate -path=internal/database/migrations -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" -verbose up

#   migrate_down:
#     migrate -path=internal/database/migrations -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" -verbose down

#   .PHONY: create_migration migrate_up migrate_down