.PHONY: run
run:
	@echo "\t\t========= RUNNING the apps =============="
	@go run cmd/*.go

.PHONY: db-up
db-up:
	@echo "\t\t ====== RUNNING MYSQL DB ======" && \
    docker-compose up --build -d db