package db

import (
	"log"
	"os"

	"github.com/soysan/chatter_n/server/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbUrl := os.Getenv("DB_URL")

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&models.Bots{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}
