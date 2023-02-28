package initializers

import "github.com/mrfade/go-jwt/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
