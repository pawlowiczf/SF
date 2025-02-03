sqlc: 
	sqlc generate

test:
	go test -v -cover ./...

server: 
	go run main.go 

migrateup:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/swift_db?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/swift_db?sslmode=disable" -verbose down
	

# migrate create -ext sql -dir ./db/migration -seq create_db