package route

import (
	"porto-backend-go/internal/handler"
	"porto-backend-go/internal/middleware"

	"github.com/gin-gonic/gin"
)

type HandlerContainer struct {
	Skill   *handler.SkillHandler
	Project *handler.ProjectHandler
	Auth    *handler.AuthHandler
	Experience *handler.ExperienceHandler
}

func SetupRoutes(r *gin.Engine, h *HandlerContainer, jwtSecret string) {
	api := r.Group("/api")
	{
		setupSkillRoutes(api, h.Skill, jwtSecret)
		setupProjectRoutes(api, h.Project, jwtSecret)
		SetupAuthRoutes(api, h.Auth)
		setupExperienceRoutes(api, h.Experience, jwtSecret)
	}
}

func setupSkillRoutes(rg *gin.RouterGroup, h *handler.SkillHandler, jwtSecret string) {
	skills := rg.Group("/skills")
	{
		skills.GET("", h.GetAllSkills)
		skills.POST("", middleware.AuthMiddleware(jwtSecret), h.CreateSkill)
		skills.DELETE("/:uuid", middleware.AuthMiddleware(jwtSecret), h.DeleteSkill)
		skills.GET("/:uuid", h.GetSkillByUUID)
		skills.PUT("/:uuid", middleware.AuthMiddleware(jwtSecret), h.UpdateSkill)
	}
}

func setupProjectRoutes(rg *gin.RouterGroup, h *handler.ProjectHandler, jwtSecret string) {
	projects := rg.Group("/projects")
	{
		projects.GET("", h.GetAllProjects)
		projects.POST("", middleware.AuthMiddleware(jwtSecret), h.CreateProject)
		projects.DELETE("/:uuid", middleware.AuthMiddleware(jwtSecret), h.DeleteProject)
		projects.GET("/:uuid", h.GetProjectByUUID)
		projects.PUT("/:uuid", middleware.AuthMiddleware(jwtSecret), h.UpdateProject)
		projects.GET("/latest", h.GetLatestProjects)
	}
}

func SetupAuthRoutes(rg *gin.RouterGroup, h *handler.AuthHandler){
	auth := rg.Group("/auth")
	{
		auth.POST("/login", h.Login)
	}
}

func setupExperienceRoutes(rg *gin.RouterGroup, h *handler.ExperienceHandler, jwtSecret string) {
	experiences := rg.Group("/experiences")
	{
		experiences.GET("", h.GetAllExperience)
		experiences.POST("", middleware.AuthMiddleware(jwtSecret), h.CreateExperience)
		experiences.DELETE("/:uuid", middleware.AuthMiddleware(jwtSecret), h.DeleteExperience)
		experiences.GET("/:uuid", h.GetExperienceByUUID)
		experiences.PUT("/:uuid", middleware.AuthMiddleware(jwtSecret), h.UpdateExperience)
	}
}
