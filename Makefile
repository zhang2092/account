DB_URL=postgresql://root:secret@localhost:5432/account?sslmode=disable

network:
	docker network create account-network

postgres:
	docker run --name postgres --network account-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it postgres createdb --username=root --owner=root account

dropdb:
	docker exec -it postgres dropdb account

migrateinit:
	migrate create -ext sql -dir db/migration -seq init_schema

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: network postgres createdb dropdb migrateup migratedown sqlc test server
