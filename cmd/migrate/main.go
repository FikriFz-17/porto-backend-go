package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"

	"porto-backend-go/internal/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	_ = config.LoadConfig()

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		url.QueryEscape(os.Getenv("DB_USER")),
		url.QueryEscape(os.Getenv("DB_PASSWORD")),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	m, err := migrate.New("file://database/migrations", dbURL)
	if err != nil {
		log.Fatalf("Failed to init migrate: %v", err)
	}
	defer m.Close()

	flag.Parse()
	switch flag.Arg(0) {
	case "up":
		if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			log.Fatalf("Migrate up failed: %v", err)
		}
		log.Println("Migrate up success!")
	case "down":
		if err := m.Steps(-1); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			log.Fatalf("Migrate down failed: %v", err)
		}
		log.Println("Rollback DOWN (1 step) success!")
	default:
		log.Println("Use: go run cmd/migrate/main.go [up|down]")
	}
}
