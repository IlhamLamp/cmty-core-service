package routes

import (
	"github.com/IlhamLamp/cmty-project-service/controllers"
	"github.com/gin-gonic/gin"
)

type AppControllers struct {
	Project *controllers.ProjectController
	Member  *controllers.MemberController
}

const projectsRoute = "/projects"
const membersRoute = "/members"

func SetupRouter(router *gin.Engine, ctl *AppControllers) {
	v1 := router.Group("/api/v1")
	{
		project := v1.Group(projectsRoute)
		{
			project.POST("/", ctl.Project.Create)
			project.GET("/", ctl.Project.GetAll)
			project.GET("/:id", ctl.Project.GetByID)
			project.PUT("/:id", ctl.Project.Update)
			project.DELETE("/:id", ctl.Project.Delete)
		}

		member := v1.Group(membersRoute)
		{
			member.POST("/", ctl.Member.Create)
			member.GET("/", ctl.Member.GetAll)
		}

		// COMMENT FOR PRODUCTION
		internal := v1.Group("/internal/seed")
		{
			internal.POST(projectsRoute, ctl.Project.SeedProjects)
			internal.DELETE(projectsRoute, ctl.Project.CleanProjects)
			internal.POST(membersRoute, ctl.Member.SeedMembers)
			internal.DELETE(membersRoute, ctl.Member.CleanMembers)
		}
	}
}
