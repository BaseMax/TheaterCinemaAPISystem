package main

import (
	"TheaterCinemaAPISystem/controller"
	"TheaterCinemaAPISystem/initializers"

	"github.com/gin-gonic/gin"
)
	
		func init()  {
		initializers.LoadEnvVarables()
		initializers.ConectToDb()
		
		}

		func main() {
	
		
			r := gin.Default()
			r.GET("/categories", controller.Category_index)
			r.GET("/categories/:id", controller.Category_show)
			r.POST("/categories", controller.Category_create)

			r.Run() // listen and serve on 0.0.0.0:8080
		}
