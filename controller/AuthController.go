package controller

import (
	"TheaterCinemaAPISystem/initializers"
	"TheaterCinemaAPISystem/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


type RegisterInput struct {
	Fullname string `json:"fullname" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
func Register_auth(c *gin.Context) {
	//Todo: validation data user
	// دریافت اطلاعات فرم ثبت نام کاربر
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// چک کردن اینکه کاربر با این نام در سیستم وجود ندارد
	var existingUser models.User
	if err := initializers.DB.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}
	//hash password to database

	//save data

	//send response user

}
func Login_auth(c *gin.Context) {

	//Todo: validation data user
	//check user to system
	//create token
	//send responce data user

}
func ForgotPassword_auth(c *gin.Context) {
	// cheack email user
	//if not  valid email  send response not valid data

	//else
	//genrate link

	// send email  link

}
func CheackVerifyEmail(c *gin.Context) {

// check url
// html password

}
func Profile_auth(c *gin.Context) {
// get id user or token
}

