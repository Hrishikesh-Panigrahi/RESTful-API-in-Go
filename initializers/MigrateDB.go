package initializers

import (
	"github.com/Hrishikesh-Panigrahi/RESTful-API-in-Go/models"
)

func SyncDb() {
	DB.AutoMigrate(&models.User{})
}
