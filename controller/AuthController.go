package controller

import (
	"TheaterCinemaAPISystem/initializers"
	"TheaterCinemaAPISystem/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(c *gin.Context) {
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

    // رمزگذاری رمز عبور با الگوریتم bcrypt
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not process password"})
        return
    }

}