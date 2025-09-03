package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrigodevcold/apimovel/handler"
)

func initializeRoutes(router *gin.Engine) {
	vi := router.Group("/api/v1")
	{
		//CRUD - Create, Read, Update, Delete
		vi.GET("/license-plate", handler.ShowLicensePlateHandler)
		vi.POST("/license-plate", handler.CreateLicensePlateHandler)
		vi.DELETE("/license-plate", handler.DeleteLicensePlateHandler)
		vi.PUT("/license-plate", handler.UpdateLicensePlateHandler)
		vi.GET("/license-plate", handler.ListLicensePlateHandler)
	}
}
