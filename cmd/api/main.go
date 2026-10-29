package main

import (
	"net/http"
	"porto-backend-go/internal/config"
	"porto-backend-go/internal/database"
	"porto-backend-go/internal/handler"
	"porto-backend-go/internal/middleware"
	"porto-backend-go/internal/repository"
	"porto-backend-go/internal/route"
	"porto-backend-go/internal/service"
	"porto-backend-go/pkg/storage"
	"time"

	_ "porto-backend-go/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"golang.org/x/time/rate"

	"github.com/gin-gonic/gin"
)

// @title                      Portfolio CMS API
// @version                    1.0
// @description                Dokumentasi REST API Portfolio CMS Backend Golang
// @BasePath                   /api
// @securityDefinitions.apikey BearerAuth
// @in                         header
// @name                       Authorization
// @description                format: Bearer <token_jwt>
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

	experienceRepo := repository.NewExperienceRepository(db)
	experienceService := service.NewExperienceService(experienceRepo)
	experienceHandler := handler.NewExperienceHandler(experienceService)

	// Setup Router
	r := gin.Default()
	_ = r.SetTrustedProxies([]string{
		"127.0.0.1",    // Localhost IPv4
		"::1",          // Localhost IPv6
		"10.0.0.0/8",   // Subnet private / Docker network
		"172.16.0.0/12",// Subnet default Docker
	})

	// Add CORS middleware
	r.Use(middleware.CORSMiddleware())

	// Add Rate Limiting middleware (60 requests per minute per IP)
	r.Use(middleware.RateLimiter(rate.Every(time.Minute/60), 10))

	// Swagger Docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/api/docs", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})

	// Setup Routes
	route.SetupRoutes(r, &route.HandlerContainer{
		Skill:   skillHandler,
		Project: projectHandler,
		Auth:    userHandler,
		Experience: experienceHandler,
	}, cfg.JWTSecret)

	// Run Server
	r.Run(":" + cfg.Port)
}