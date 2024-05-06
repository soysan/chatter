package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/soysan/chatter_n/server/pkg/db"
	"github.com/soysan/chatter_n/server/pkg/handlers"
	"github.com/soysan/chatter_n/server/pkg/middleware"
	"os"
)

func main() {
	// Render.com用の無駄 todo fix if they fixed
	_, ok := os.LookupEnv("PRODUCTION_CHAT")
	if !ok {
		err := godotenv.Load()
		if err != nil {
			panic(err)
		}
	}

	dbIns := db.Init()
	// inject db to handlers
	h := handlers.New(dbIns)

	router := gin.Default()
	router.Use(cors.New(middleware.CorsSettings))

	router.GET("/healthz", h.Healthz)
	router.POST("/chat", h.Chat)
	router.GET("/history/list", h.HistoryList)

	port := os.Getenv("PORT")
	if port == "" {
		port = "6565"
	}
	router.Run(":" + port)
}
