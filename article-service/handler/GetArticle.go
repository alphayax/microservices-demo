package handler

import (
	"article-service/repository"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func GetArticle(ctx *gin.Context) {
	articles, err := repository.GetArticles(ctx, nil)
	if err != nil {
		log.Warnf("GetArticle Error: %s", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"articles": articles,
	})
}
