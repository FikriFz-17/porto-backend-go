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
}

func SetupRoutes(r *gin.Engine, h *HandlerContainer, jwtSecret string) {
	api := r.Group("/api")
	{
		setupSkillRoutes(api, h.Skill, jwtSecret)
		setupProjectRoutes(api, h.Project, jwtSecret)
		SetupAuthRoutes(api, h.Auth)
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
	}
}

func SetupAuthRoutes(rg *gin.RouterGroup, h *handler.AuthHandler){
	auth := rg.Group("/auth")
	{
		auth.POST("/login", h.Login)
	}
}
