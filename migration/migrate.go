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
}