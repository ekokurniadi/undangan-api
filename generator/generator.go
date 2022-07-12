package generator

import (
	"github.com/ekokurniadi/micagen"
	"gorm.io/gorm"
)

type generatorConfig struct {
	db    *gorm.DB
	model interface{}
}

func NewGeneratorConfig(db *gorm.DB, model interface{}) *generatorConfig {
	return &generatorConfig{db, model}
}

func (g *generatorConfig) PlayGenerator() {
	micagen.CreateFormatter(g.db, g.model)
	micagen.CreateHandler(g.db, g.model)
	micagen.CreateRepository(g.db, g.model)
	micagen.CreateService(g.db, g.model)
	micagen.CreateStructInput(g.model)
}
