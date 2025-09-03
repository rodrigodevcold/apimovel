package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListLicensePlateHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "List opening",
	})
}
