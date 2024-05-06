package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/soysan/chatter_n/server/pkg/apis"
	"github.com/soysan/chatter_n/server/pkg/consts"
	"github.com/soysan/chatter_n/server/pkg/models"
)

var timeLayout = "15:04:05"

func createResponse(userInput string, botResponse string) models.Res {
	return models.Res{
		UserInput:         userInput,
		BotResponse:       botResponse,
		ResponseTimestamp: time.Now().Format(time.RFC3339),
	}
}

func (d *Db) Chat(c *gin.Context) {
	var req models.Req
	if err := c.BindJSON(&req); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println(req)

	var res models.Res
	switch req.UserInput {
	case "こんにちは":
		res = createResponse(req.UserInput, "こんにちは。Botです。")
	case "今何時？":
		res = createResponse(req.UserInput, fmt.Sprintf("現在の時刻は %v です。", time.Now().Format(timeLayout)))
	case "今日の東京の天気は？":
		weather, err := apis.FetchWeather()
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		res = createResponse(req.UserInput, fmt.Sprintf("東京の天気は %v です。", consts.Weather[weather.Daily.WeatherCode[0]]))
	default:
		res = createResponse(req.UserInput, "すみません、理解できませんでした。")
	}

	err := d.InsertOne(res)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}
