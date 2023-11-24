package handler

import (
	"cart-service/reporitory"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateCart(ctx *gin.Context) {
	cartId := ctx.Param("cartId")

	var updateCartRequest UpdateCartRequest
	if err := ctx.ShouldBindJSON(&updateCartRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := reporitory.UpdateCart(ctx, cartId, updateCartRequest.Items)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusAccepted)
}

type UpdateCartRequest struct {
	Items []string `json:"items"`
}
