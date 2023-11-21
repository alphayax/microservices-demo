package handler

import (
	"article-service/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteArticle(ctx *gin.Context) {
	articleId := ctx.Param("articleId")

	if err := repository.DeleteArticle(ctx, articleId); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":     err.Error(),
			"articleId": articleId,
		})
		return
	}

	ctx.Status(http.StatusAccepted)
}
