package controller

import (
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}

	u.Username = input.Username
	u.Password = input.Password

	_, err := u.SaveUser()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success"})

}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login_auth(c *gin.Context)  {

	//Todo: validation data user
	var input LoginInput
	//check user to system
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}
	u := models.User{}

	u.Username = input.Username
	u.Password = input.Password
	//create token
	token, err := models.LoginCheck(u.Username, u.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	//send responce data user
	c.JSON(http.StatusOK, gin.H{"token":token})

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
