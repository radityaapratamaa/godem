export LINTER_VERSION ?= 1.52.2

.PHONY: run
run:
	@echo "\t\t========= RUNNING the apps =============="
	APPS_ENV=$(env) IS_DEBUG=$(debug) air -c .air.conf

.PHONY: db-up
db-up:
	@echo "\t\t ====== RUNNING POSTGRES DB ======" && \
    docker-compose up --build -d db

.PHONY: test
test:
	@echo "\t\t ====== RUNNING unit test ======" && \
    go test -v -cover -race `go list ./... | grep -v "vendor/" | grep -v "testfile/" | grep -v "vendor.orig/" | grep -v "docker/" | grep -v "mocks/"`

test-out:
	@echo "\t\t ====== RUNNING unit test ======" && \
    go test -v -coverprofile=cover.out -cover -race `go list ./... | grep -v "vendor/" | grep -v "testfile/" | grep -v "vendor.orig/" | grep -v "docker/" | grep -v "mocks/"` > test-result.out

check-pr-action: test lint

install-air:
	@echo "\t\t ====== INSTALLING air hot reload ======"
	@go install github.com/cosmtrek/air@latest && \
	air -v

install-goose:
	@echo "\t\t ====== INSTALLING goose DB Migrations ======"
	@go install github.com/pressly/goose/v3/cmd/goose@latest

migration-up:
	export GOOSE_DRIVER=postgres && export GOOSE_DBSTRING="postgres://postgres:postgres@localhost:32771/godem" && \
	cd database/migrations && \
	goose up && \
	cd ../..

migration-down:
	export GOOSE_DRIVER=postgres && export GOOSE_DBSTRING="postgres://postgres:postgres@localhost:32771/godem" && \
	cd database/migrations && \
	goose down && \
	cd ../..

migration-reset:
	export GOOSE_DRIVER=postgres && export GOOSE_DBSTRING="postgres://postgres:postgres@localhost:32771/godem" && \
	cd database/migrations && \
    goose reset && \
    cd ../..

bin:
	@mkdir -p bin

tool-lint: bin
	@test -e ./bin/golangci-lint || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ./bin v${LINTER_VERSION}

lint: tool-lint
	./bin/golangci-lint run -v --timeout 3m0s

dep:
	go mod download
