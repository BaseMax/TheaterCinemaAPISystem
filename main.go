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
			r.PUT("movies/:id",controller.Movies_update)
			r.GET("movies/:id",controller.Movies_show)
			r.DELETE("movies/:id",controller.Movies_delete)
			r.PATCH("movies/:id/active",controller.Movies_active)

			r.Run() // listen and serve on 0.0.0.0:8080
		}
