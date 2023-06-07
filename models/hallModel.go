package models

import (
	"gorm.io/gorm"
)

type Halls  struct {
	gorm.Model
	Title string `json:"title"`
	Body string `json:"body"`
	Span uint  `json:"span"`
	Active bool `json:"active"`
	Movies []Movies `gorm:"many2many:MOViE_Hall;"`
}