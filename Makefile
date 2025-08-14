export $(shell sed 's/=.*//' .env)


docker.up.rebuild: swagger
	@docker compose up --build

docker.up:
	@docker compose up

docker.down:
	@docker compose down

docker.teardown:
	@docker compose down -v --remove-orphans

migrations.up:
	@docker compose run --rm migrations sh -c 'migrate -path /migrations -database "$$DATABASE_URL" up'

migrations.down:
	@docker compose run --rm migrations sh -c 'migrate -path /migrations -database "$$DATABASE_URL" down 1'

swagger:
	@swag init -g ./cmd/server/main.go