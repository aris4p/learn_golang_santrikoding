package main

import (
	"github.com/aris4p/config"
	"github.com/aris4p/database"
	"github.com/aris4p/routes"
)

func main() {

	//load config .env
	config.LoadEnv()

	//inisialisasi database
	database.InitDB()

	// setup router
	r := routes.SetupRouter()

	//mulai server dengan port 3000
	r.Run(":" + config.GetEnv("APP_PORT", "3001"))
}
