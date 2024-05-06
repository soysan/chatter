package middleware

import (
	"github.com/gin-contrib/cors"
	"time"
)

var CorsSettings = cors.Config{
	AllowOrigins: []string{
		"http://localhost:3000",
	},
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
