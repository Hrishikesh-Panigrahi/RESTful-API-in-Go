package main

import (
	// "net/http"

	// "net/http"

	"github.com/Hrishikesh-Panigrahi/RESTful-API-in-Go/connectors"
	"github.com/Hrishikesh-Panigrahi/RESTful-API-in-Go/initializers"
	"github.com/gin-gonic/gin"
	// "github.com/Hrishikesh-Panigrahi/RESTful-API-in-Go/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
	initializers.SyncDb()
}

// func Hello(c *gin.Context)
// {

// c.JSON({

// "msg":"hello"

// })

// }

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {

		// Call the HTML method of the Context to render a template
		c.JSON(201, gin.H{
			"msg": "hello",
		})

	})

	router.POST("/register", connectors.Register)
	// router.POST("/login", connectors.Login)

	router.Run()
}
