postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root piket

dropdb:
	docker exec -it postgres12 dropdb piket

migrateup:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/piket?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/piket?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test: 
	go test -v -cover ./...

.PHONY:
	createdb
	dropdb
	postgres
	migrateup
	migratedown
	sqlc 
	test 