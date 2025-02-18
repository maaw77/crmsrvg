include .env

POSTGRESQL_URL='postgres://$(POSTGRES_USER):${POSTGRES_PASSWORD}@localhost:5433/${POSTGRES_DB}?sslmode=disable'

.DEFAULT_GOAL := help

.PHONY: help
help: ## display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

PHONY: migr_cr 
migr_cr: ### create new migration
	migrate -verbose create -ext sql -dir ./migrations -seq create_aux_tables

PHONY: migr_up  
migr_up: ### migration up
	migrate -verbose -database ${POSTGRESQL_URL} -path ./migrations up 
	
PHONY: migr_down  
migr_down: ### print current migration version
	migrate -verbose -database ${POSTGRESQL_URL} -path ./migrations down

PHONY: migr_ver 
migr_ver: ### migration down
	migrate -verbose -database ${POSTGRESQL_URL} -path ./migrations version 

# migrate -database postgres://postgres:crmpassword@localhost:5433/postgres?sslmode=disable -path ./migrations up

#   include .env

#   create_migration:
#     migrate create -ext=sql -dir=internal/database/migrations -seq init

#   migrate_up:
#     migrate -path=internal/database/migrations -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" -verbose up

#   migrate_down:
#     migrate -path=internal/database/migrations -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" -verbose down

#   .PHONY: create_migration migrate_up migrate_down