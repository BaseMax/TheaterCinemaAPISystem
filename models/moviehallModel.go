package models

import (
	"gorm.io/gorm"
)

type MOViE_Hall struct {
	gorm.Model
	Hall_ID Halls  `gorm:"primaryKey"`
	Movie_id Movies `gorm:"primaryKey"`
	Year_start string
    Month_start string
	Day_start string
	Hour_start string
	Minutes_start string
	Span uint
	Year_end string
    Month_end string
	Day_end string
	Hour_end string
	Minutes_end string
	Number_sales uint
}