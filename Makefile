include .env
export

# Run project
up:
	docker-compose up -d && go build -tags dev -o app ./cmd && ./app

up-prod:
	docker-compose up -d && go build -o app ./cmd && ./app

# Example: make migration-create create_orders_table
migration-create:
	migrate create -ext sql -dir migrations -seq $(filter-out $@,$(MAKECMDGOALS))

# Run migrations
migrate:
	migrate -path migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@localhost:5432/$(DB_NAME)?sslmode=disable" up
