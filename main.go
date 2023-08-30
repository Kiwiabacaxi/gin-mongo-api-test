package main

import (
	"gin-mongo-api/configs"    // import the configs module
	"gin-mongo-api/routes"     // import the routes module
	"github.com/gin-gonic/gin" // import gin framework
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "Hello from Gin-gonic & mongoDB Teste 1 :D",
		})
	})

	//run database
	configs.ConnectDB()

	// run routes
	routes.UserRoute(router)

	router.Run("localhost:6000")
}
