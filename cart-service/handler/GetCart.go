package handler

import (
	"cart-service/model"
	"cart-service/reporitory"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCart(ctx *gin.Context) {
	cartId := ctx.Param("cartId")

	cart, err := reporitory.GetCart(ctx, cartId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// If cart is nil, create a new one
	if cart == nil {
		cart = &model.Cart{
			Id: cartId,
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"cart": cart,
	})
}
