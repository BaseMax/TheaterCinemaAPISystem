package controller

import (
	"TheaterCinemaAPISystem/initializers"
	"TheaterCinemaAPISystem/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register_auth(c *gin.Context) {
    // دریافت اطلاعات فرم ثبت نام کاربر
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }

    // چک کردن اینکه کاربر با این نام در سیستم وجود ندارد
    var existingUser models.User
    if err := initializers.DB.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
        c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
        return
    }



}
func Login_auth(c *gin.Context){

}
func ForgotPassword_auth(c *gin.Context){

}
func Profile_auth(c *gin.Context){

}