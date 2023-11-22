package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

func main() {
	log.Infoln("-= Article Cart =-")
	loadConfig()
	initDatabase()
	loadApiServer()
}

// loadConfig define the default values and loads the user configuration from config.yaml
func loadConfig() {
	viper.SetDefault("listen", ":8081")
	//viper.SetDefault("redisUri", "")
	err := viper.BindEnv("redisUri", "REDIS_URI")
	if err != nil {
		log.Warnln(err)
	}
	viper.SetConfigFile("config.yaml")
	viper.AddConfigPath("/etc/article-cart/")
	if err := viper.ReadInConfig(); err != nil {
		log.Warnln(err)
	}
}

// initDatabase initialize the database connection
func initDatabase() {
	// init redis database connection
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

	//Router.GET("/", handler.GetArticle)
	//Router.POST("/", handler.AddArticle)
	//Router.DELETE("/:articleId/", handler.DeleteArticle)

	listenAddress := viper.GetString("listen")
	err := Router.Run(listenAddress)
	log.Panicln(err)
}
