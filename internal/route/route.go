package route

import (
	"porto-backend-go/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, skillHandler *handler.SkillHandler) {
	api := r.Group("/api")
	{
		api.GET("/skills", skillHandler.GetAllSkills)
	}
}
