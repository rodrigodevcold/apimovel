package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteLicensePlateHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete opening",
	})
}
