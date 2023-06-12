package main

import (
	"TheaterCinemaAPISystem/initializers"
	"TheaterCinemaAPISystem/models"
)

func init() {
	initializers.LoadEnvVarables()
	initializers.ConectToDb()
}
func main() {
	initializers.DB.AutoMigrate(&models.Categories{})
	initializers.DB.AutoMigrate(&models.Movies{})
	initializers.DB.AutoMigrate(&models.Halls{})
	initializers.DB.AutoMigrate(&models.MOViE_Hall{})
	initializers.DB.AutoMigrate(&models.User{})
}