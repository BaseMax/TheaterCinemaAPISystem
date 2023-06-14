package controller

import (
	"TheaterCinemaAPISystem/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterInput struct {
	Email    string `json:"email" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register_auth(c *gin.Context) {
	//Todo: validation data user
	// دریافت اطلاعات فرم ثبت نام کاربر
	fmt.Println("hi ...")
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

func Login_auth(c *gin.Context) {

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
	// c.JSON(http.StatusBadRequest, gin.H{"error": input.Password, "username": input.Username})
	// c.JSON(http.StatusBadRequest, gin.H{"error model": u.Username, "username": u.Password})
	//create token
	token, err := models.LoginCheck(u.Username, u.Password)
	// c.JSON(http.StatusBadRequest, gin.H{"error": err, "token": token})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//send responce data user
	c.JSON(http.StatusOK, gin.H{"username": u.Username, "token": token})

}

type ForgetPassInput struct {
	Email string `json:"email" binding:"required"`
}

func ForgotPassword_auth(c *gin.Context) {
	//get input
	var input ForgetPassInput
	// cheack email user
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u := models.User{}

	if u.Email == input.Email {
		c.JSON(http.StatusBadRequest, gin.H{"sucsess ": "user in send reset password "})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user in_valide data "})
	}
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
