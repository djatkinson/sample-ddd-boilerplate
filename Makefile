include .env

migrateup:
	@migrate -path migrations/postgres -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose up
migratedown:
	@migrate -path migrations/postgres -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose down
migraterollbackversion:
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
