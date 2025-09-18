package main

import (
	"github.com/IlhamLamp/cmty-core-service/config"
	"github.com/IlhamLamp/cmty-core-service/database"
	"github.com/IlhamLamp/cmty-core-service/di"
	"github.com/IlhamLamp/cmty-core-service/routes"
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
