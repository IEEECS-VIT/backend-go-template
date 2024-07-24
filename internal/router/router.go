package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sanjayheaven/ggb/internal/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	SetupTeamRoutes(r)

	return r
}
func SetupTeamRoutes(r *gin.Engine) {
	teamRoutes := r.Group("/team")
	{
		teamRoutes.POST("/create", controllers.CreateTeam)
	}
}
