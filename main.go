package main

import (
	"TheaterCinemaAPISystem/controller"
	"TheaterCinemaAPISystem/initializers"
	"TheaterCinemaAPISystem/middlewares"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVarables()
	initializers.ConectToDb()


}

func main() {
	r := gin.Default()
	// max  memory 3
	r.MaxMultipartMemory = 3 << 20
	// category
	r.GET("/categories", controller.Category_index)
	r.GET("/categories/:id", controller.Category_show)
	r.POST("/categories", controller.Category_create)
	r.PUT("/categories/:id", controller.Category_update)
	r.DELETE("/categories/:id", controller.Category_delete)
	//movies
	r.GET("/movies", controller.Movies_index)
	r.POST("/movies", controller.Movies_create)
	r.PUT("movies/:id", controller.Movies_update)
	r.GET("movies/:id", controller.Movies_show)
	r.DELETE("movies/:id", controller.Movies_delete)
	r.PATCH("movies/:id/active", controller.Movies_active)

	//halls
	r.GET("/halls", controller.Hall_index)
	r.GET("/halls/:id", controller.Hall_show)
	r.POST("/halls", controller.Hall_create)
	r.PUT("/halls/:id", controller.Hall_update)
	r.DELETE("/halls/:id", controller.Hall_delete)

	//schedule
	r.GET("/movie/:Id/schedule", controller.Schedule_index)
	r.POST("/movie/:Id/schedule", controller.Schedule_create)
	//auth
	r.POST("/auth/register", controller.Register_auth)
	r.POST("/auth/login", controller.Login_auth)
	
	// ایجاد یک روت برای بازیابی رمز عبور کاربر
	r.POST("/auth/forgot_password", controller.ForgotPassword_auth)

	admin := r.Group("/api/admin")
	admin.POST("/dashbord", controller.Dashbord)
	admin.GET("/profile", controller.Profile_auth)
	admin.Use(middlewares.JwtAuthMiddleware())
	

	r.Run()
}
