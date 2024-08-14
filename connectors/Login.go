package connectors

import (
	"net/http"
	"os"
	"time"

	"github.com/Hrishikesh-Panigrahi/RESTful-API-in-Go/initializers"
	"github.com/Hrishikesh-Panigrahi/RESTful-API-in-Go/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input",
		})
		return
	}

	// find the user
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid credentials",
		})
		return
	}

	ValidatePassword(user.Password, body.Password, c)
	JwtToken(c, user)
}

func ValidatePassword(userPassword string, bodyPassword string, c *gin.Context) {
	// take the password and hash it and compare it to the password in the database
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(bodyPassword))

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid password",
		})
	}
}

func JwtToken(c *gin.Context, user models.User) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to generate token",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("token", tokenString, 3600*24*30, "", "", false, true)
}
