package controller

import (
	"TheaterCinemaAPISystem/initializers"
	"TheaterCinemaAPISystem/models"

	"github.com/gin-gonic/gin"
)

/*
* ----show  all category  resoult json 
*/
func Category_index(c *gin.Context) {
		var categories []models.Categories
		initializers.DB.Find(&categories)
	
	c.JSON(200, gin.H{
		"status": true,
		"data":categories,
		
	})

}

/*
* ----create category  resoult json 
*/
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

/*
* ----show category use id resoult json 
*/
func Category_show(c *gin.Context){

		// get id
        id := c.Param("id")

		//find id to db
		var category models.Categories
		initializers.DB.First(&category, id)
		//result json 
		c.JSON(200, gin.H{
			"status":true,
			"message": "show data id : "+id,
			"data": category,
		})
}

/*
* ----update category use id resoult json 
*/
func Category_update(c *gin.Context){

		// get id
		id := c.Param("id")
		//get data request 
		var body struct{
			Title string
			Slug string 
			Parint_id uint
		}
		c.Bind(&body)
		//find category 
		var categories models.Categories
		initializers.DB.First(&categories, id)
		//update 
		initializers.DB.Model(&categories).Updates(models.Categories{
			Title: body.Title,
			Slug: body.Slug,
			Parint_id: body.Parint_id,
		})
		//resolte
		c.JSON(200, gin.H{
			"status":true,
			"message": "update data id : "+id,
			"data": categories,
		})
}

/*
* ----delete category use id resoult json 
*/
func Category_delete(c *gin.Context){
	// get id
	id := c.Param("id")
	
	//find category 
	var categories models.Categories
	initializers.DB.Delete(&categories,id)

	//resolte
	c.JSON(200, gin.H{
		"status":true,
		"message": "delete data id : "+id,
	
	})
	
}
