package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (d *Db) Healthz(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Health Check OK!",
	})
}
