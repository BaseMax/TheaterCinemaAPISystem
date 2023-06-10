package controller

import (
	"TheaterCinemaAPISystem/initializers"
	"TheaterCinemaAPISystem/models"

	"github.com/gin-gonic/gin"
)

/*
* ----show  all category  resoult json 
*/
func Schedule_index(c *gin.Context) {
		var moviehall []models.MOViE_Hall
		initializers.DB.Find(&moviehall)
	
	c.JSON(200, gin.H{
		"status": true,
		"data":moviehall,
		
	})

}

/*
* ----create category  resoult json 
*/
func Schedule_create(c *gin.Context){

		var body struct{
			Hall_id  uint
			Movie_id uint 
	
		}
		c.Bind(&body)

		cat:= models.MOViE_Hall{Hall_id:body.Hall_id ,Movie_id:body.Movie_id }
		res :=initializers.DB.Create(&cat);
		if res!=nil {
			c.Status(400);
		}
		c.JSON(200, gin.H{
			"status":true,
			"message": "create data",
		})
}
