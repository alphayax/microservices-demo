package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HealthZ(c *gin.Context) {
	c.Status(http.StatusOK)
}
