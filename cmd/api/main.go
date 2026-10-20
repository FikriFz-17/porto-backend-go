package main

import (
	"porto-backend-go/internal/config"
	"porto-backend-go/internal/database"
	"porto-backend-go/internal/handler"
	"porto-backend-go/internal/repository"
	"porto-backend-go/internal/route"
	"porto-backend-go/internal/service"
	"porto-backend-go/pkg/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load Config & Database
	cfg := config.LoadConfig()
	db := database.InitDB(cfg.DBDSN)

	storageService, err := storage.NewCloudinaryService()
	if err != nil{
		panic("Failed to Init Storage: " + err.Error())
	}

	// Wiring Modul Skill (Repo -> Service -> Handler)
	skillRepo := repository.NewSkillRepository(db)
	skillService := service.NewSkillService(skillRepo, storageService)
	skillHandler := handler.NewSkillHandler(skillService)

	// Wiring Modul Project (Repo -> Service -> Handler)
	projectRepo := repository.NewProjectRepository(db)
	projectService := service.NewProjectService(projectRepo, storageService)
	projectHandler := handler.NewProjectHandler(projectService)

	userRepo := repository.NewUserRepository(db)
	userService := service.NewAuthService(userRepo)
	userHandler := handler.NewAuthHandler(userService)

	// Setup Router
	r := gin.Default()

	// Setup Routes
	route.SetupRoutes(r, &route.HandlerContainer{
		Skill:   skillHandler,
		Project: projectHandler,
		Auth:    userHandler,
	}, cfg.JWTSecret)

	// Run Server
	r.Run(":" + cfg.Port)
}