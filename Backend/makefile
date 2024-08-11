postgres:
	docker run --name postgres16 -p 5432:5432  -e POSTGRES_USER=root -e POSTGRES_PASSWORD=passwd -d postgres:16.0-alpine3.18

createdb:
	docker exec -it postgres16 createdb --user=root --owner=root hirirng_platform

dropdb:
	docker exec -it postgres16 dropdb --user=root hirirng_platform

migrateup:
	migrate -path db/migration -database "postgresql://root:passwd@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:passwd@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown
