package models

import (
	"gorm.io/gorm"
)

type Movies  struct {
	gorm.Model
	Title string `json:"title"`
	Synopsis string `json:"synopsis"`
	Body string `json:"body"`
	Category_id uint `json:"category_id"`
	Price string `json:"price"`
	Year string `json:"year"`
	View uint   `json:"view_movie"`
	Image string `json:"img_url"`
	Actived bool `json:"active"`

}