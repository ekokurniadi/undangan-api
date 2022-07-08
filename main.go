package main

import (
	"undangan_online_api/config"
	"undangan_online_api/routes"
)

func main() {
	database, _ := config.InitialDatabaseConfig()

	routes.NewRoutesInjector(database).RunAppRouters()

}
