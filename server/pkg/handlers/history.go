package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (d *Db) HistoryList(c *gin.Context) {
	rows, err := d.FetchLatest10()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, rows)
}
