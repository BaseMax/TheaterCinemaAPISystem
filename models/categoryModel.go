package models

import "gorm.io/gorm"

type Categories struct {
	gorm.Model
	Title string `json:"title"`
	Slug string `json:"slug"`
	Parint_id uint `json:"parint_id"`
}