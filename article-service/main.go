package main

import (
	"article-service/handler"
	"article-service/repository"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"time"
)

func main() {
	log.Println("-= Article Service =-")
	loadConfig()
	initDatabase()
	loadApiServer()
}

// loadConfig define the default values and loads the user configuration from config.yaml
func loadConfig() {
	viper.SetDefault("listen", ":8080")
	viper.SetDefault("mongodbUri", "mongodb://localhost:27017/alpha-articles")
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Println(err)
	}
}

// initDatabase initialize the database connection
func initDatabase() {
	mongodbUri := viper.GetString("mongodbUri")
	if err := repository.Initialize(mongodbUri); err != nil {
		log.Panic(err)
	}
}

// loadApiServer initialize the API server with a cors middleware and define routes to be served.
// This function is blocking: it will wait until the server returns an error
func loadApiServer() {
	Router := gin.New()
	Router.Use(cors.New(cors.Config{
		//AllowOrigins:     []string{"http://localhost:3001"},
		AllowMethods:     []string{"GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Accept, Origin, Cache-Control, X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true //return origin == "http://localhost:3001"
		},
		MaxAge: 12 * time.Hour,
	}))

	Router.GET("/", handler.GetArticle)
	Router.POST("/", handler.AddArticle)
	Router.DELETE("/:articleId/", handler.DeleteArticle)

	listenAddress := viper.GetString("listen")
	err := Router.Run(listenAddress)
	log.Panicln(err)
}
