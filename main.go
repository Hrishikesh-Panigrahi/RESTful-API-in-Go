package main

import (
	// "net/http"

	"github.com/gin-gonic/gin"
	// "github.com/Hrishikesh-Panigrahi/RESTful-API-in-Go/connectors"
	"github.com/Hrishikesh-Panigrahi/RESTful-API-in-Go/initializers"
	// "github.com/Hrishikesh-Panigrahi/RESTful-API-in-Go/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
	initializers.SyncDb()
}

func main() {
	router := gin.Default()
	// router.POST("/register", connectors.Register)
	// router.POST("/login", connectors.Login)

	router.Run()
}
