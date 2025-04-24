postgres:
	docker run --name postgres16 -p 5432:5432  -e POSTGRES_USER=root -e POSTGRES_PASSWORD=hiring_platform -d postgres:16.0-alpine3.18

createdb:
	docker exec -it postgres16 createdb --user=root --owner=root hiring_platform

dropdb:
	docker exec -it postgres16 dropdb --user=root hiring_platform

migrateup:
	migrate -path db/migration -database "postgresql://root:hiring_platform@localhost:5432/hiring_platform?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:hiring_platform@localhost:5432/hiring_platform?sslmode=disable" -verbose down

sqlc:
	sqlc generate

httpserver:
	go run main.go --server=http

grpcserver:
	go run main.go --server=grpc

setupdb:
	docker run --name postgres16 -p 5432:5432  -e POSTGRES_USER=root -e POSTGRES_PASSWORD=hiring_platform -d postgres:16.0-alpine3.18
	docker exec -it postgres16 createdb --user=root --owner=root hiring_platform
	migrate -path db/migration -database "postgresql://root:hiring_platform@localhost:5432/hiring_platform?sslmode=disable" -verbose up

cleandb:
	migrate -path db/migration -database "postgresql://root:hiring_platform@localhost:5432/hiring_platform?sslmode=disable" -verbose down
	docker exec -it postgres16 dropdb --user=root hiring_platform
	
proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    proto/*.proto

.PHONY: postgres createdb dropdb migrateup migratedown sqlc server
