package router

import "github.com/gin-gonic/gin"

func initializeRoutes(router *gin.Engine) {
	vi := router.Group("/api/v1")
	{
		//show API status
		vi.GET("/opening", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"GET message": "API is up and running!",
			})
		})
	}
	{
		vi.POST("/opening", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"POST message": "API is up and running!",
			})
		})
	}
	{
		vi.DELETE("/opening", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"DELETE message": "API is up and running!",
			})
		})
	}
	{
		vi.PUT("/opening", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"PUT message": "API is up and running!",
			})
		})
	}
	{
		vi.GET("/openings", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"GET message": "API is up and running!",
			})
		})
	}

}
