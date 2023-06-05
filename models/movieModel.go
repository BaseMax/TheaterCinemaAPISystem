package models

import (
	"gorm.io/gorm"
)

type Movies  struct {
	gorm.Model
	Title string
	Synopsis string
	Body string
	Category_id uint
	Price string
	Year string
	View uint
	Image string

}