migrationPath=./cmd/migrate/migrations
dsn=postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable

build: clean
	@mkdir -p ./bin;
	@go build -o ./bin/api cmd/api/main.go;

clean:
	@rm -f ./bin/api;

run:
	@./bin/api;

dev:
	@go run cmd/api/main.go;

db-status:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(dsn) goose --dir $(migrationPath) status;

.PHONY: build clean run;
