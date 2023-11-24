package handler

import (
	"cart-service/reporitory"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteCart(ctx *gin.Context) {
	cartId := ctx.Param("cartId")

	err := reporitory.DeleteCart(ctx, cartId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusAccepted)
}
