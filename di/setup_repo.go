package di

import (
	"github.com/IlhamLamp/cmty-project-service/controllers"
	"github.com/IlhamLamp/cmty-project-service/repository"
	"github.com/IlhamLamp/cmty-project-service/routes"
	"github.com/IlhamLamp/cmty-project-service/services"
	"gorm.io/gorm"
)

func SetupRepositories(db *gorm.DB) *routes.AppControllers {
	projectRepo := repository.NewProjectRepository(db)
	memberRepo := repository.NewMemberRepository(db)

	projectService := services.NewProjectService(projectRepo)
	memberService := services.NewMemberService(memberRepo)

	projectController := controllers.NewProjectController(projectService)
	memberController := controllers.NewMemberController(memberService)

	return &routes.AppControllers{
		Member:  memberController,
		Project: projectController,
	}
}
