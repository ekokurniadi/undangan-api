package main

import (
	"undangan_online_api/config"
	"undangan_online_api/migrationconfig"
	"undangan_online_api/routes"
)

func main() {
	database, _ := config.InitialDatabaseConfig()
	migrationconfig.NewMigrationList(database).MigrateAllEntities()
	routes.NewRoutesInjector(database).RunAppRouters()

}
