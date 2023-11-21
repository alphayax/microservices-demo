package handler

import (
	"article-service/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetArticle(ctx *gin.Context) {
	articles, err := repository.GetArticles(ctx, nil)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"articles": articles,
	})
}
