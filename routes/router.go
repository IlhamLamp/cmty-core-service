package routes

import (
	"github.com/IlhamLamp/cmty-project-service/controllers"
	"github.com/gin-gonic/gin"
)

const projectsRoute = "/projects"

func SetupRouter(router *gin.Engine, projectController *controllers.ProjectController) {
	v1 := router.Group("/api/v1")
	{
		project := v1.Group(projectsRoute)
		{
			project.POST("/", projectController.Create)
			project.GET("/", projectController.GetAll)
			project.GET("/:id", projectController.GetByID)
			project.PUT("/:id", projectController.Update)
			project.DELETE("/:id", projectController.Delete)
		}

		// COMMENT FOR PRODUCTION
		internal := v1.Group("/internal/seed")
		{
			internal.POST(projectsRoute, projectController.SeedProjects)
			internal.DELETE(projectsRoute, projectController.CleanProjects)
		}
	}
}
