package handlers

import (
	"log"

	"github.com/soysan/chatter_n/server/pkg/models"
	"gorm.io/gorm"
)

type Db struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *Db {
	log.Println("db connected")
	return &Db{db}
}

func (d *Db) InsertOne(data models.Res) error {
	result := d.DB.Create(&models.Bots{
		UserInput:         data.UserInput,
		BotResponse:       data.BotResponse,
		ResponseTimestamp: data.ResponseTimestamp,
	})

	if result.Error != nil {
		return result.Error
	}

	log.Println("data inserted")
	return nil
}

func (d *Db) FetchLatest10() ([]models.Bots, error) {
	var rows []models.Bots
	result := d.DB.Order("response_timestamp desc").Limit(10).Find(&rows)
	if result.Error != nil {
		return nil, result.Error
	}

	log.Println("data fetched")
	return rows, nil
}
