package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/takumines/go-sandbox-api/config"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	logger := config.NewLogger()

	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/health-check", func(c *gin.Context) {
			logger.Debug().Msg("Starting application")
			c.JSON(200, gin.H{
				"message": "OK",
			})
		})
	}

	r.Run()
}
