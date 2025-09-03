package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateLicensePlateHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Update opening",
	})
}
