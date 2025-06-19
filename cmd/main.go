package main

import (
	"github.com/IlhamLamp/cmty-project-service/config"
	"github.com/IlhamLamp/cmty-project-service/database"
	"github.com/IlhamLamp/cmty-project-service/routes"
)

func main() {
	conf := config.LoadEnv()
	dbConn := database.DbConnection{
		Host:     conf.DbHost,
		Name:     conf.DbName,
		User:     conf.DbUsername,
		Password: conf.DbPassword,
	}
	database.Connect(dbConn)

	// database.AutoMigrate(db)

	r := routes.SetupRouter()
	r.Run(":" + conf.AppPort)
}
