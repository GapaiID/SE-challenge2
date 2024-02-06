## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'


# runserver: run application
.PHONY: runserver
runserver:
	go run . runserver


## makemigrations name=$1: create a new database migration
.PHONY: makemigrations
db/migration/new:
	@echo "Create migration files for ${name}"
	go run . makemigrations -f ${name}


## migrate: migrate untracked migration files
.PHONY: migrate
migrate:
	go run . migrate -e up


## audit: check quality of codes
.PHONY: audit
audit:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Formatting code...'
	go fmt ./...
	@echo 'Vetting code...'
	go vet ./...
	staticcheck ./...
	@echo 'Running tests...'
	go test -race -vet=off ./...
