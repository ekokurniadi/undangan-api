package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitialDatabaseConfig() (*gorm.DB, error) {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
