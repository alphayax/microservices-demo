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
	log.Println("-= Article Alpha =-")
	loadConfig()
	initDatabase()
	loadApiServer()
}

func loadConfig() {
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Println(err)
	}
}

func initDatabase() {
	mongodbUri := viper.GetString("mongodbUri")
	if err := repository.Initialize(mongodbUri); err != nil {
		log.Panic(err)
	}
}

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

	err := Router.Run(":8080")
	log.Panicln(err)
}
