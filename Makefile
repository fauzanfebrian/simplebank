ifneq (,$(wildcard ./.env))
    include .env
    export
endif

migratecreate:
	migrate create -ext sql -dir db/migration -seq $(name)
migrateforce:
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose force 1
migrateup:
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose up
migratedown:
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose down
migrateup1:
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose up 1
migratedown1:
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose down 1

sqlc:
	sqlc generate && echo "sqlc generated successfully"
test:
	go clean -testcache
	GIN_MODE=test go test -v -cover ./...
server:
	rm -f ./bin/simplebank
	go build -o bin/simplebank .
	./bin/simplebank
mock:
	mockery -d

db_docs:
	dbdocs build doc/db.dbml

db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

setupdeps:
	(command -v sqlc >/dev/null 2>&1 || go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.21.0) && echo "sqlc installed successfully"
	(command -v migrate >/dev/null 2>&1 || go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest) && echo "migrate installed successfully"
	(command -v mockery >/dev/null 2>&1 || go install github.com/vektra/mockery/v2@v2.33.3) && echo "mockery installed successfully"
	(command -v dbdocs >/dev/null 2>&1 || npm install -g dbdocs) && echo "dbdocs installed successfully"
	(command -v dbml2sql >/dev/null 2>&1 || npm install -g @dbml/cli) && echo "dbml installed successfully"

.SILENT:
.PHONY: migrateup migratedown sqlc test server mock setupdeps migratecreate migrateup1 migratedown1 migrateforce migratecreate db_docs db_schema
.DEFAULT_GOAL := server