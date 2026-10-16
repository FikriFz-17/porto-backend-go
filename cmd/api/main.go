package main

import (
	"porto-backend-go/internal/database"
	"porto-backend-go/internal/config"
	"porto-backend-go/internal/handler"
	"porto-backend-go/internal/repository"
	"porto-backend-go/internal/route"
	"porto-backend-go/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// config
	cfg:=config.LoadConfig()

	// Database connection
	db := database.InitDB(cfg.DBDSN)

	// Repository
	skillRepo := repository.NewSkillRepository(db)

	// Service
	skillService := service.NewSkillService(skillRepo)

	// Handler
	handler := handler.NewSkillHandler(skillService)

	// Router
	r := gin.Default()
	route.SetupRoutes(r, handler)

	r.Run(":8000")
}