package main

import (
	"facebookfix/engine"
	"facebookfix/handlers"
	"html/template"
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

	unEscapeHTML := make(map[string]interface{})
	unEscapeHTML["unescapeHTML"] = func(s string) template.HTML {
		return template.HTML(s)
	}
	mode := gin.DebugMode
	environment := os.Getenv("ENVIRONMENT")
	if environment == "PRODUCTION" {
		mode = gin.ReleaseMode
	}
	gin.SetMode(mode)

	router := gin.New()
	router.SetFuncMap(unEscapeHTML)
	router.Use(engine.CORSMiddleware())
	router.Use(engine.CustomHeaders())
	router.Use(gin.LoggerWithFormatter(engine.HTTPLogger))

	router.LoadHTMLGlob("template/*")

	//routes
	router.GET("/", handlers.HomeHandler)
	router.GET("/watch", handlers.GetVideoHandler)
	router.GET("/:accountName/videos/:id", handlers.GetVideoHandler)
	router.GET("/reel/:id", handlers.GetReelHandler)
	router.GET("/photo", handlers.GetPhotoHandler)
	router.GET("/marketplace/item/:id", handlers.GetMarketHandler)

	port := os.Getenv("PORT")
	if err := router.Run(":" + port); err != nil {
		log.Fatalln(err)
	}
}
