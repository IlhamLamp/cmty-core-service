package main

import (
	"github.com/IlhamLamp/cmty-project-service/config"
	"github.com/IlhamLamp/cmty-project-service/controllers"
	"github.com/IlhamLamp/cmty-project-service/database"
	"github.com/IlhamLamp/cmty-project-service/repository"
	"github.com/IlhamLamp/cmty-project-service/routes"
	"github.com/IlhamLamp/cmty-project-service/services"
	"github.com/gin-gonic/gin"
)

func main() {
	conf := config.LoadEnv()

	dbreq := database.DbConnection{
		Host:     conf.DbHost,
		Name:     conf.DbName,
		User:     conf.DbUsername,
		Password: conf.DbPassword,
	}
	db := database.Connect(dbreq)

	projectRepo := repository.NewProjectRepository(db)
	projectService := services.NewProjectService(projectRepo)
	projectController := controllers.NewProjectController(projectService)

	router := gin.Default()
	routes.SetupRouter(router, projectController)

	router.Run(":" + conf.AppPort)
}
