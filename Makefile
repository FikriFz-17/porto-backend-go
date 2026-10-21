create-migration:
	migrate create -ext sql -dir database/migrations -seq $(name)
migrate-up:
	go run cmd/migrate/main.go up
migrate-down:
	go run cmd/migrate/main.go down