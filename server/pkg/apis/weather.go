package apis

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type WeatherData struct {
	Current struct {
		WeatherCode int `json:"weather_code"`
	} `json:"current"`
	Daily struct {
		WeatherCode []int `json:"weather_code"`
	} `json:"daily"`
}

func FetchWeather() (WeatherData, error) {
	apiUrl := "https://api.open-meteo.com/v1/forecast?latitude=35.6895&longitude=139.6917&current=temperature_2m,weather_code&daily=weather_code&timezone=Asia%2FTokyo&forecast_days=1"

	var weather WeatherData

	resp, err := http.Get(apiUrl)
	if err != nil {
		log.Println(err)
		return weather, fmt.Errorf("failed to get whether data: %w", err)
	}
	defer resp.Body.Close()

	// bufio.NewScanner(resp.Body)も試す
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return weather, fmt.Errorf("failed to read response body: %w", err)
	}

	err = json.Unmarshal(body, &weather)
	if err != nil {
		log.Println(err)
		return weather, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return weather, nil
}
