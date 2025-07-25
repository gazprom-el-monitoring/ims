.PHONY:
.SILENT:
.DEFAULT_GOAL := run
#include .env
#export $(shell sed 's/=.*//' .env)

test:
	go test --short -coverprofile=cover.out -v ./...
	make test.coverage
test.coverage:
	go tool cover -func=cover.out | grep "total"
run:
	go run ./cmd/rest/main.go
migrate:
	migrate -path db/migrations -database "postgresql://$$DB_USER:$$DB_PASSWORD@$$DB_SERVER:$$DB_PORT/$$DB_NAME?sslmode=disable" -verbose up
migrate_force:
	migrate -path db/migrations -database "postgresql://$$DB_USER:$$DB_PASSWORD@$$DB_SERVER:$$DB_PORT/$$DB_NAME?sslmode=disable" -verbose force $(id)