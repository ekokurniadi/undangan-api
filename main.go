package main

import (
	"undangan_online_api/config"
	"undangan_online_api/entity"
	"undangan_online_api/generator"
	"undangan_online_api/migrationconfig"
	"undangan_online_api/routes"
)

func main() {
	database, _ := config.InitialDatabaseConfig()
	generator.NewGeneratorConfig(database, &entity.TurutMengundang{})
	migrationconfig.NewMigrationList(database).MigrateAllEntities()
	routes.NewRoutesInjector(database).RunAppRouters()

}
