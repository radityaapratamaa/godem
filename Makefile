.PHONY: run
run:
	@echo "\t\t========= RUNNING the apps =============="
	@go run cmd/*.go

.PHONY: db-up
db-up:
	@echo "\t\t ====== RUNNING MYSQL DB ======" && \
    docker-compose up --build -d db

.PHONY: test
test:
	@echo "\t\t ====== RUNNING unit test ======" && \
    go test -v -coverprofile=cover.out -cover -race `go list ./... | grep -v "vendor/" | grep -v "testfile/" | grep -v "vendor.orig/" | grep -v "docker/" | grep -v "mocks/"` > test-result.out