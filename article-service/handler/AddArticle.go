package handler

import (
	"article-service/model"
	"article-service/repository"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddArticle(ctx *gin.Context) {
	var addArticleRequest AddArticleRequest
	if err := ctx.ShouldBindJSON(&addArticleRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := addArticleRequest.validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := repository.AddArticle(ctx, addArticleRequest.toArticle())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

type AddArticleRequest struct {
	ArticleName        string `json:"name"`
	ArticleDescription string `json:"description"`
}

func (r AddArticleRequest) validate() error {
	if r.ArticleName == "" {
		return fmt.Errorf("articleName is required")
	}
	if r.ArticleDescription == "" {
		return fmt.Errorf("articleDescription is required")
	}
	return nil
}

func (r AddArticleRequest) toArticle() *model.Article {
	return &model.Article{
		Title:       r.ArticleName,
		Description: r.ArticleDescription,
	}
}
