package controller

import (
	"TheaterCinemaAPISystem/initializers"
	"TheaterCinemaAPISystem/models"

	"github.com/gin-gonic/gin"
)
func Category_index(c *gin.Context) {
		var categories []models.Categories
		initializers.DB.Find(&categories)
	
	c.JSON(200, gin.H{
		"status": true,
		"data":categories,
		
	})

}
func Category_create(c *gin.Context){
		var body struct{
			Title string
			Slug string 
			Parint_id uint
		}
		c.Bind(&body)

	cat:= models.Categories{Title:body.Title,Slug:body.Slug,Parint_id: body.Parint_id}
	res :=initializers.DB.Create(&cat);
	if res!=nil {
		c.Status(400);
	}
	c.JSON(200, gin.H{
		"status":true,
		"message": "create data",
	})
}
func Category_show(c *gin.Context){

		// get id

		//find id to db
		// initializers.DB.First(&models.Categories, 10)
		//result json 

}