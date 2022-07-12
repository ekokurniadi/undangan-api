package migrationconfig

import (
	"undangan_online_api/entity"

	"gorm.io/gorm"
)

type migrationsList struct {
	db *gorm.DB
}

func NewMigrationList(db *gorm.DB) *migrationsList {
	return &migrationsList{db}
}

func (m *migrationsList) MigrateAllEntities() {
	var migrationsList MigrationList

	entities := make([]interface{}, 0)

	entities = append(
		entities,
		entity.Galeri{},
		entity.GuestBook{},
		entity.Hadiah{},
		entity.LoveStory{},
		entity.Undangan{},
		entity.UndanganDetail{},
		entity.User{},
		entity.RootLink{},
		entity.TurutMengundang{},
	)

	migrationsList.Entities = entities

	for _, myEntities := range migrationsList.Entities {
		m.db.AutoMigrate(&myEntities)
	}

}
