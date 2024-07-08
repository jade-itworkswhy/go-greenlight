include .env 

.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]


## run/api: run the cmd/api application, default - port 8080, env development
.PHONY: run/api
run/api:
	go run ./cmd/api

## run/api/no-limit: run the cmd/api application, without rate limit
.PHONY: run/api/no-limit
run/api/no-limit:
	go run ./cmd/api -limiter-enabled=false

## run/api/cors: run the cmd/api application, with cors trusted origins
.PHONY: run/api/cors
run/api/cors:
	go run ./cmd/api -cors-trusted-origins=${ORIGINS}

## run/example-cors: run the cors example application(frontend)
.PHONY: run/web/example-cors
run/web/example-cors:
	go run ./cmd/examples/cors/simple


## kill/api: kill the cmd/api application using defined port
.PHONY: kill/api
kill/api: confirm
	@PID=$$(lsof -ti:${PORT}); \
	if [ ! -z "$$PID" ]; then \
		kill -9 $$PID; \
	else \
		echo "No process is using port 8080"; \
	fi


## db/psql: connect to the database using psql
.PHONY: db/psql
db/psql: 
	psql ${DB_DSN}

## db/migrations/new name=$1: create a new database migration
.PHONY: db/migrations/new
db/migrations/new:
	@echo 'Creating migration files for ${name}...'
	migrate create -seq -ext=.sql -dir=./migrations ${name}

## db/migrations/up: apply all up database migrations
.PHONY: db/migrations/up
db/migrations/up: confirm
	@echo 'Running up migrations...'
	migrate -path ./migrations -database ${DB_DSN} up
