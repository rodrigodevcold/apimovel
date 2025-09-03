package router

import "github.com/gin-gonic/gin"

func Initialize() {
	//initialize gin router
	router := gin.Default()
	//initialize routes
	initializeRoutes(router)

	//run the server
	router.Run(":8080") // listen and serve on
}
