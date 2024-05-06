package main

import (
    "os"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/soysan/chatter_n/server/pkg/db"
	"github.com/soysan/chatter_n/server/pkg/handlers"
	"github.com/soysan/chatter_n/server/pkg/middleware"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}

	db := db.Init()
	// inject db to handlers
	h := handlers.New(db)

	router := gin.Default()
	router.Use(cors.New(middleware.CorsSettings))

	router.GET("/healthz", h.Healthz)
	router.POST("/chat", h.Chat)
	router.GET("/history/list", h.HistoryList)

    port := os.Getenv("PORT")
	router.Run(":" + port)
}
