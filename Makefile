GOFILES := $(shell find . -name "*.go" ! -path "./.go/*")
PACKAGES ?= $(shell go list ./... | grep -v /vendor/)
GO_MOD_DIR ?= $(shell find . -name "go.mod" ! -path "./.go/*" -exec dirname {} \;)

# Make settings
.DEFAULT_GOAL := help
.PHONY: deps \
		fmt \
		db-run \
		vet \
		misspell \
		golang_ci_lint \
		lint \
		race \
		unittest-coverage \
		unittest \
		ci-check \
		download \
		help

# Make goals
frontend-run:  ## Runs frontend
	cd frontend && yarn serve

backend-run:  ## Runs compose
	docker-compose -f docker-compose.yml down --volumes --remove-orphans && docker-compose -f docker-compose.yml up --build

api-run-local:  ## Runs container with all needed dependencies
	cd ${GO_MOD_DIR} && go run .

db-run:  ## Runs container with database
	docker-compose -f docker-compose.yml down --volumes --remove-orphans db && docker-compose -f docker-compose.yml up --build db

api-run:  ## Runs container with database
	docker-compose -f docker-compose.yml down --volumes --remove-orphans api && docker-compose -f docker-compose.yml up --build api

db-container-exec:  ## Enters database container
	docker exec -it calories-scanner_db_1 /bin/bash

db-login:  ## Login psql console
	PGPASSWORD=postgres psql -U postgres -h localhost -p 5433

deps: ## Synchronises all dependencies
	cd ${GO_MOD_DIR} && go mod tidy

fmt: ## Formats all Golang files
	gofmt -l -s -w $(GOFILES)

vet: ## Runs `go vet` for all packages
	cd ${GO_MOD_DIR} && go vet $(PACKAGES)

misspell: ## Runs misspell
	@hash misspell > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		echo 'Please install "misspell" tool: https://github.com/client9/misspell'; \
		exit 1; \
	fi
	misspell -w $(GOFILES)

golang_ci_lint: ## Runs golangci-lint
	@hash golangci-lint > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		echo 'Please install "golangci-lint" tool: https://github.com/golangci/golangci-lint'; \
		exit 1; \
	fi
	cd ${GO_MOD_DIR} && golangci-lint run -v

lint: golang_ci_lint misspell ## Runs golang lint

race: ## Runs all unit tests with -race flag
	cd ${GO_MOD_DIR} && go test -race -coverprofile coverage.out -v ./... && go tool cover -html=coverage.out -o ../../coverage_report.html

unittest-coverage: ## Runs all unit tests with coverage
	cd ${GO_MOD_DIR} && go test -coverprofile coverage.out -v ./... && go tool cover -html=coverage.out -o ../../coverage_report.html

unittest: ## Runs all unit tests
	cd ${GO_MOD_DIR} && go test

ci-check: deps fmt vet lint race ## Combines `deps` `fmt` `vet` `misspell` `lint` `race` `build` commands

download: ## Downloads all dependencies
	cd ${GO_MOD_DIR} && go mod download

help: ## Displays this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ \
	{ printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)		
