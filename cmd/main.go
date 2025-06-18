package main

import (
	"github.com/IlhamLamp/cmty-project-service/config"
	"github.com/IlhamLamp/cmty-project-service/database"
	"github.com/IlhamLamp/cmty-project-service/routes"
)

func main() {
	config.LoadEnv()
	database.Connect()

	// database.AutoMigrate(db)

	r := routes.SetupRouter()
	r.Run(":3010 ")
}
