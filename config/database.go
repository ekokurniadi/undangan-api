package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitialDatabaseConfig() (*gorm.DB, error) {

	env := godotenv.Load()
	if env != nil {
		log.Fatal(env.Error())
	}

	constantName := DeclareConstant(
		os.Getenv("HOST"),
		os.Getenv("DATABASE"),
		os.Getenv("USER_DB"),
		os.Getenv("PASSWORD_DB"),
		os.Getenv("PORT_DB"),
	)

	dsn := "host=" + constantName.Host + " user=" + constantName.Username + " password=" + constantName.Password + " dbname=" + constantName.Database + " port=" + constantName.Port + " sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
