# export $(shell sed 's/=.*//' .env)


docker.up.rebuild: swagger
	@docker compose -f build/dev/docker-compose.dev.yaml up --build

docker.up:
	@docker compose -f build/dev/docker-compose.dev.yaml up

docker.down:
	@docker compose -f build/dev/docker-compose.dev.yaml down

docker.teardown:
	@docker compose -f build/dev/docker-compose.dev.yaml down -v --remove-orphans

migrations.up:
	@docker compose -f build/dev/docker-compose.dev.yaml run --rm migrations sh -c 'migrate -path /migration_files -database "postgres://$$POSTGRES_USER:$$POSTGRES_PASSWORD@database:5432/$$POSTGRES_DB?sslmode=disable" up'

migrations.down:
	@docker compose -f build/dev/docker-compose.dev.yaml run --rm migrations sh -c 'migrate -path /migration_files -database "postgres://$$POSTGRES_USER:$$POSTGRES_PASSWORD@database:5432/$$POSTGRES_DB?sslmode=disable" down 1'

swagger:
	@swag init -g ./cmd/server/main.go

lint:
	@golangci-lint run

docker.up.test_db:
	@docker compose -f build/test/docker-compose.test.yaml up -d test_database
	@until docker compose -f build/test/docker-compose.test.yaml exec test_database pg_isready -U test_user; do sleep 1; done

test.integration: docker.up.test_db
	@docker compose -f build/test/docker-compose.test.yaml run --rm test_app go test ./tests/... -v -cover; docker compose -f build/test/docker-compose.test.yaml down test_database