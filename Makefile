include .env

build:
	go build -o ./bin/${APP_NAME}-http-server cmd/main.go
build-migration:
	go build -o ./bin/${APP_NAME}-migration cmd/migration/main.go
start-server:
	make build
	./bin/${APP_NAME}-http-server
run-migrate-up:
	make build-migration
	./bin/${APP_NAME}-migration migrate up
run-migrate-down:
	make build-migration
	./bin/${APP_NAME}-migration migrate down
migrateup:
	while read line; do export $line; done < .env
	@migrate -path migrations/postgres -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose up
migratedown:
	@while read line; do export $line; done < .env
	@migrate -path migrations/postgres -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose down
migraterollbackversion:
	while read line; do export $line; done < .env
	@migrate -path migrations/postgres -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose goto $(version);
generatemigrate:
	@if [ -z "$(file)" ]; then echo "define file name file={file-name}"; exit 1; fi
	migrate create -ext sql -dir ./migration -seq $(file)


test:
	go test ./...
test-coverage:
	if [ ! -d "test-coverage" ];then     \
			mkdir test-coverage;           \
	fi
	go test -coverprofile=test-coverage/coverage.out ./... ; go tool cover -func=test-coverage/coverage.out

.PHONY: migrateup migratedown migraterollbackversion test-coverage
