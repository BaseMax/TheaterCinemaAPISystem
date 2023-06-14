package controller

import (
	"TheaterCinemaAPISystem/initializers"
	"TheaterCinemaAPISystem/models"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

/*
* ----show  all Movies  resoult json
 */
func Movies_index(c *gin.Context) {
	var movies []models.Movies
	initializers.DB.Find(&movies)

	c.JSON(200, gin.H{
		"status": true,
		"data":   movies,
	})

}

/*
* ----create Movies  resoult json
 */
func Movies_create(c *gin.Context) {
	var body struct {
		Title       string
		Body        string
		Year        string
		Category_id uint
		Price       string
		Small_body  string
		Actived     bool
		Minutes     uint
	}
	fmt.Println(body.Title ,body.Actived);
	c.Bind(&body)

	movies := models.Movies{Title: body.Title, Synopsis: body.Small_body, Minutes: body.Minutes, Body: body.Body, Year: body.Year, Category_id: body.Category_id, Price: body.Price, View: 1, Actived: body.Actived}

	res  := initializers.DB.Create(&movies)
	if res != nil {
		c.Status(400)
		c.JSON(200, gin.H{
			"status":  false,
			
		})
	}
	c.JSON(200, gin.H{
		"status":  true,
		"message": "create data",
	})
}

/*
* ----show Movies use id resoult json
 */
func Movies_show(c *gin.Context) {

	// get id
	id := c.Param("id")

	//find id to db
	var Movies models.Movies
	initializers.DB.First(&Movies, id)
	//result json
	c.JSON(200, gin.H{
		"status":  true,
		"message": "show data id : " + id,
		"data":    Movies,
	})
}

/*
* ----update Movies use id resoult json
 */
func Movies_update(c *gin.Context) {

	// get id
	id := c.Param("id")
	//get data request
	var body struct {
		Title       string
		Synopsis    string
		Body        string
		Year        string
		Category_id uint
		Price       string
		Actived     bool
		Minutes     uint
	}
	c.Bind(&body)
	//find Movies
	var movie models.Movies
	initializers.DB.First(&movie, id)
	//update
	initializers.DB.Model(&movie).Updates(models.Movies{
		Title:       body.Title,
		Body:        body.Body,
		Synopsis:    body.Synopsis,
		Price:       body.Price,
		Actived:     body.Actived,
		Category_id: body.Category_id,
		Minutes:     body.Minutes,
	})

	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	// Upload the file to specific dst.
	c.SaveUploadedFile(file, "./../public/image/"+id+"/")

	//resolte
	c.JSON(200, gin.H{
		"status":  true,
		"message": "update data id : " + id,
		"data":    movie,
	})
}

/*
* ----delete Movies use id resoult json
 */
func Movies_delete(c *gin.Context) {
	// get id
	id := c.Param("id")

	//find Movies and delete
	var movie models.Movies
	initializers.DB.Delete(&movie, id)

	//resolte
	c.JSON(200, gin.H{
		"status":  true,
		"message": "delete data id : " + id,
	})

}
func Movies_active(c *gin.Context) {
	// get id
	id := c.Param("id")

	//find Movies
	var movie models.Movies
	initializers.DB.First(&movie, id)
	if movie.Actived == false {
		movie.Actived = true
		initializers.DB.Model(&movie).Updates(models.Movies{
			Actived: movie.Actived,
		})
	} else if movie.Actived == true {
		movie.Actived = false
		initializers.DB.Model(&movie).Updates(models.Movies{
			Actived: movie.Actived,
		})
	}

	//resolte
	c.JSON(200, gin.H{
		"status":  true,
		"message": "delete data id : " + id,
		"data":    movie.Actived,
	})

}
