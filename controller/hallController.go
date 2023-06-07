package controller

import (
	"TheaterCinemaAPISystem/initializers"
	"TheaterCinemaAPISystem/models"

	"github.com/gin-gonic/gin"
)

/*
* ----show  all Hall  resoult json 
*/
func Hall_index(c *gin.Context) {
		var hall []models.Halls
		initializers.DB.Find(&hall)
	
	c.JSON(200, gin.H{
		"status": true,
		"data":hall,
		
	})

}

/*
* ----create Hall  resoult json 
*/
func Hall_create(c *gin.Context){
		var body struct{
			Title string
			Body string 
			Active bool
		}
		c.Bind(&body)

	hall:= models.Halls{Title:body.Title,Body:body.Body,Active: body.Active}
	res :=initializers.DB.Create(&hall);
	if res!=nil {
		c.Status(400);
	}
	c.JSON(200, gin.H{
		"status":true,
		"message": "create data",
	})
}

/*
* ----show Hall use id resoult json 
*/
func Hall_show(c *gin.Context){

		// get id
        id := c.Param("id")

		//find id to db
		var Hall models.Halls
		initializers.DB.First(&Hall, id)
		//result json 
		c.JSON(200, gin.H{
			"status":true,
			"message": "show data id : "+id,
			"data": Hall,
		})
}

/*
* ----update Hall use id resoult json 
*/
func Hall_update(c *gin.Context){

		// get id
		id := c.Param("id")
		//get data request 
		var body struct{
			Title string
			Body string 
			Active bool
	
		}
		c.Bind(&body)
		//find Hall 
		var hall models.Halls
		initializers.DB.First(&hall, id)
		//update 
		initializers.DB.Model(&hall).Updates(models.Halls{Title:body.Title,Body:body.Body,Active: body.Active,})
		//resolte
		c.JSON(200, gin.H{
			"status":true,
			"message": "update data id : "+id,
			"data": hall,
		})
}

/*
* ----delete Hall use id resoult json 
*/
func Hall_delete(c *gin.Context){
	// get id
	id := c.Param("id")
	
	//find Hall 
	var hall models.Halls
	initializers.DB.Delete(&hall,id)

	//resolte
	c.JSON(200, gin.H{
		"status":true,
		"message": "delete data id : "+id,
	
	})
	
}
