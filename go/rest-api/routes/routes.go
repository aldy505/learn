package routes

import (
	"rest-api/middlewares"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	router := gin.Default()

	router.Use(middlewares.Secure())

	users := router.Group("/users")
	{
		users.Use(middlewares.CheckApiKey())
		users.GET("/", Fetch)
		users.POST("/login", Login)
		users.POST("/register", Register)
	}

	return router
}
