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
		setupProjectRoutes(api, h.Project)
		SetupAuthRoutes(api, h.Auth)
	}
}

func setupSkillRoutes(rg *gin.RouterGroup, h *handler.SkillHandler, jwtSecret string) {
	skills := rg.Group("/skills")
	{
		skills.GET("", h.GetAllSkills)
		skills.POST("", middleware.AuthMiddleware(jwtSecret), h.CreateSkill)
	}
}

func setupProjectRoutes(rg *gin.RouterGroup, h *handler.ProjectHandler) {
	projects := rg.Group("/projects")
	{
		projects.GET("", h.GetAllProjects)
	}
}

func SetupAuthRoutes(rg *gin.RouterGroup, h *handler.AuthHandler){
	auth := rg.Group("/auth")
	{
		auth.POST("/login", h.Login)
	}
}
