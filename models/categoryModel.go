package models

import "gorm.io/gorm"

type Categories struct {
	gorm.Model
	Title string
	Slug string
	Parint_id uint
}