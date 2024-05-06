package middleware

import (
	"github.com/gin-contrib/cors"
	"time"
	"os"
)

var origins []string

func InitCors() cors.Config {
    if url := os.Getenv("FRONTEND_URL"); url != "" {
        origins = append(origins, url)
    }

    origins = append(origins, "http://localhost:3000")

    return cors.Config{
        AllowOrigins: origins,
        AllowMethods: []string{
            "GET",
            "POST",
        },
        AllowHeaders: []string{
            "Authorization",
            "Content-Type",
        },
        AllowCredentials: true,
        MaxAge:           24 * time.Hour,
    }
}

var CorsSettings = InitCors()