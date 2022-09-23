package bootstrap

import "un1versum.com/clutchers/models"

func migrateDB() {
	DB.AutoMigrate(&models.User{})
}
