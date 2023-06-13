package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Fullname string `gorm:"not null"`
    Username string `gorm:"unique;not null"`
    Password string `gorm:"not null"`
    Email    string `gorm:"unique;not null"`
    Active   uint   `gorm:"default:0"`
    Status   uint  `gorm:"default:1"`
    CodeActive uint
}