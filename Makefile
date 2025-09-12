DEV_COMPOSE_FILE = build/dev/docker-compose.dev.yaml
PROD_COMPOSE_FILE = build/dev/docker-compose.prod.yaml

TEST_COMPOSE_FILE = build/test/docker-compose.test.yaml
TESTS_PATH = ./tests/...
UNIT_TESTS_PATH = ./tests/unit/...
INTEGRATION_TESTS_PATH = ./tests/integration/...


# Develop environment
swagger:
	@swag init -g ./cmd/server/main.go

lint:
	@golangci-lint run

runlet.start:
	@docker compose -f ${DEV_COMPOSE_FILE} up

runlet.rebuild.start: swagger
	@docker compose -f ${DEV_COMPOSE_FILE} up --build

runlet.stop:
	@docker compose -f ${DEV_COMPOSE_FILE} down

runlet.teardown:
	@docker compose -f ${DEV_COMPOSE_FILE} down -v --remove-orphans

migrations.up:
	@docker compose -f ${DEV_COMPOSE_FILE} run --rm migrations sh -c 'migrate -path /migration_files -database "postgres://$$POSTGRES_USER:$$POSTGRES_PASSWORD@database:5432/$$POSTGRES_DB?sslmode=disable" up'

migrations.down:
	@docker compose -f ${DEV_COMPOSE_FILE} run --rm migrations sh -c 'migrate -path /migration_files -database "postgres://$$POSTGRES_USER:$$POSTGRES_PASSWORD@database:5432/$$POSTGRES_DB?sslmode=disable" down 1'
#----------------------------------------------------------------------------------------------------------------------

# Testing
runlet.start.test_db:
	@docker compose -f ${TEST_COMPOSE_FILE} up -d test_database
	@until docker compose -f ${TEST_COMPOSE_FILE} exec test_database pg_isready -U test_user; do sleep 1; done

runlet.test.integration: runlet.start.test_db
	@docker compose -f ${TEST_COMPOSE_FILE} run --rm test_app go test ${INTEGRATION_TESTS_PATH} -v -coverpkg=./internal/... -cover; docker compose -f ${TEST_COMPOSE_FILE} down test_database

runlet.test.unit:
	@docker compose -f ${TEST_COMPOSE_FILE} run --rm test_app go test ${UNIT_TESTS_PATH} -v -coverpkg=./internal/... -cover

runlet.test.full: runlet.start.test_db
	@docker compose -f ${TEST_COMPOSE_FILE} run --rm test_app go test ${TESTS_PATH} -v -coverpkg=./internal/... -cover; docker compose -f ${TEST_COMPOSE_FILE} down test_database
#----------------------------------------------------------------------------------------------------------------------