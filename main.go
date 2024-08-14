package main

import (
	"github.com/Hrishikesh-Panigrahi/RESTful-API-in-Go/connectors"
	"github.com/Hrishikesh-Panigrahi/RESTful-API-in-Go/initializers"
	"github.com/Hrishikesh-Panigrahi/RESTful-API-in-Go/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
	initializers.SyncDb()
}

func main() {
	router := gin.Default()

	router.POST("/register", connectors.Register)
	router.POST("/login", connectors.Login)

	authorized := router.Group("/", middleware.AuthMiddleware)
	{
		authorized.GET("/users", connectors.GetUsers)
	}

	router.Run()
}
