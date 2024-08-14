package connectors

import (
	"fmt"
	"net/http"

	"github.com/Hrishikesh-Panigrahi/RESTful-API-in-Go/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"github.com/Hrishikesh-Panigrahi/RESTful-API-in-Go/initializers"
)

func Register(c *gin.Context) {
	type body struct {
		Email    string `json:"Email"`
		Password string `json:"password"`
		Name     string `json:"name"` // A regular string field
		Role     string `json:"role"` // A regular string field
	}

	var b body

	if c.Bind(&b) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to read body",
		})
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(b.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to hash password",
		})
		return
	}

	//create the user
	var user = &models.User{Email: b.Email, Password: string(passwordHash), Name: b.Name, Role: b.Role}

	fmt.Print((user))
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{

			"message": "failed to create user",
		})

		return
	}

	// code to register a user
	fmt.Print("Registering a user...")

	c.JSON(http.StatusOK, gin.H{})
}
