package main

import (
	"github.com/IlhamLamp/cmty-project-service/config"
	"github.com/IlhamLamp/cmty-project-service/database"
	"github.com/IlhamLamp/cmty-project-service/di"
	"github.com/IlhamLamp/cmty-project-service/routes"
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

	router := gin.Default()
	repos := di.SetupRepositories(db)
	routes.SetupRouter(router, repos)

	router.Run(":" + conf.AppPort)
}
