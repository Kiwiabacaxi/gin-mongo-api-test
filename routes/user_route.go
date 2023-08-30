package routes

import (
	"gin-mongo-api/controllers" // importa o modulo de controllers
	"github.com/gin-gonic/gin"
)

// UserRoute : All routes related to users comes here

func UserRoute(router *gin.Engine) {
	//All routes related to users comes here
	router.POST("/user", controllers.CreateUser()) // cria um usuario
	router.GET("/user/:userId", controllers.GetAUser())    // pega um usuario
}
