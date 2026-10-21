package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"porto-backend-go/internal/config"
	"porto-backend-go/internal/database"
)

func main() {
	_ = godotenv.Load()
	cfg := config.LoadConfig()
	db := database.InitDB(cfg.DBDSN)

	if len(os.Args) < 2 {
		log.Fatal("Usage: make seed name=<table_or_file_name> or make seed name=all")
	}

	target := strings.TrimSuffix(os.Args[1], ".json")

	if target == "all" {
		files, err := filepath.Glob("database/seed/*.json")
		if err != nil {
			log.Fatalf("Failed to scan seed folder: %v", err)
		}
		for _, file := range files {
			tableName := strings.TrimSuffix(filepath.Base(file), ".json")
			seedFromFile(db, tableName, file)
		}
		return
	}

	filePath := fmt.Sprintf("database/seed/%s.json", target)
	seedFromFile(db, target, filePath)
}

func seedFromFile(db *gorm.DB, tableName string, filePath string) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file %s: %v", filePath, err)
	}

	var records []map[string]interface{}
	if err := json.Unmarshal(data, &records); err != nil {
		log.Fatalf("Error parsing JSON in %s: %v", filePath, err)
	}

	if len(records) == 0 {
		log.Printf("[%s] No data to seed in %s", tableName, filePath)
		return
	}

	if err := db.Table(tableName).Clauses(clause.OnConflict{DoNothing: true}).Create(&records).Error; err != nil {
		log.Fatalf("[%s] Failed to seed: %v", tableName, err)
	}

	log.Printf("[%s] Successfully seeded %d rows from %s", tableName, len(records), filePath)
}
