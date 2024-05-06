package models

import "gorm.io/gorm"

type Bots struct {
	gorm.Model
	UserInput         string `json:"user_input" gorm:"column:user_input"`
	BotResponse       string `json:"bot_response" gorm:"column:bot_response"`
	ResponseTimestamp string `json:"response_timestamp" gorm:"column:response_timestamp"`
}
