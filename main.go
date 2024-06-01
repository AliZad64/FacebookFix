package main

import (
	"facebookfix/engine"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mode := gin.DebugMode
	environment := os.Getenv("ENVIRONMENT")
	if environment == "PRODUCTION" {
		mode = gin.ReleaseMode
	}
	gin.SetMode(mode)

	router := gin.New()
	router.Use(engine.CORSMiddleware())
	router.Use(engine.CustomHeaders())
	router.Use(gin.LoggerWithFormatter(engine.HTTPLogger))
	port := os.Getenv("PORT")
	if err := router.Run(":" + port); err != nil {
		log.Fatalln(err)
	}
}
