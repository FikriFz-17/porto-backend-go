package config

import (
	"os"
	"github.com/joho/godotenv"
)

type Config struct{
	Port string
	DBDSN string
	JWTSecret string
}

func LoadConfig() *Config {
	_ = godotenv.Load()
	dsn := "host=" + os.Getenv("DB_HOST") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" port=" + os.Getenv("DB_PORT") +
		" sslmode=" + os.Getenv("DB_SSLMODE")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return &Config{Port: port, DBDSN: dsn, JWTSecret: os.Getenv("JWT_SECRET")}
}
