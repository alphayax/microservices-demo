package handler

import (
	"article-service/repository"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func DeleteArticle(ctx *gin.Context) {
	articleId := ctx.Param("articleId")

	if err := repository.DeleteArticle(ctx, articleId); err != nil {
		log.Warnf("DeleteArticle Error: %s", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":     err.Error(),
			"articleId": articleId,
		})
		return
	}

	ctx.Status(http.StatusAccepted)
}
