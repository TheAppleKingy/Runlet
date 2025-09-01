swagger:
	@swag init -g ./cmd/server/main.go

lint:
	@golangci-lint run

runlet.start:
	@docker compose -f build/dev/docker-compose.dev.yaml up

runlet.rebuild.start: swagger
	@docker compose -f build/dev/docker-compose.dev.yaml up --build

runlet.stop:
	@docker compose -f build/dev/docker-compose.dev.yaml down

runlet.teardown:
	@docker compose -f build/dev/docker-compose.dev.yaml down -v --remove-orphans

migrations.up:
	@docker compose -f build/dev/docker-compose.dev.yaml run --rm migrations sh -c 'migrate -path /migration_files -database "postgres://$$POSTGRES_USER:$$POSTGRES_PASSWORD@database:5432/$$POSTGRES_DB?sslmode=disable" up'

migrations.down:
	@docker compose -f build/dev/docker-compose.dev.yaml run --rm migrations sh -c 'migrate -path /migration_files -database "postgres://$$POSTGRES_USER:$$POSTGRES_PASSWORD@database:5432/$$POSTGRES_DB?sslmode=disable" down 1'

runlet.start.test_db:
	@docker compose -f build/test/docker-compose.test.yaml up -d test_database
	@until docker compose -f build/test/docker-compose.test.yaml exec test_database pg_isready -U test_user; do sleep 1; done

runlet.test.integration: runlet.start.test_db
	@docker compose -f build/test/docker-compose.test.yaml run --rm test_app go test ./tests/... -v -cover; docker compose -f build/test/docker-compose.test.yaml down test_database