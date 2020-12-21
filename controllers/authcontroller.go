package controllers

import (
	"fmt"
	"goseed/models/entity"
	"goseed/models/service"
	"log"

	"github.com/gin-gonic/gin"

	"golang.org/x/crypto/bcrypt"
)

//AuthController is for auth logic
type AuthController struct{}

//Login is to process login request
func (auth *AuthController) Login(c *gin.Context) {

	var loginInfo entity.User
	if err := c.ShouldBindJSON(&loginInfo); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	//TODO
	userservice := service.Userservice{}
	user, errf := userservice.Find(&loginInfo)
	if errf != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": "Not found"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInfo.Password))
	if err != nil {
		c.AbortWithStatusJSON(403, gin.H{"error": "Email or password is invalid."})
		return
	}

	fmt.Println("user email is ", user.Email)
	token, err := user.GetJwtToken()
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}
	//-------
	c.JSON(200, gin.H{
		"token": token,
	})
}

//Profile is to provide current user info
func (auth *AuthController) Profile(c *gin.Context) {
	user := c.MustGet("user").(*(entity.User))

	c.JSON(200, gin.H{
		"name":     user.Name,
		"lastname": user.LastName,
		"email":    user.Email,
		"cpf":      user.Cpf,
	})
}

//Signup is for user signup
func (auth *AuthController) Signup(c *gin.Context) {

	type signupInfo struct {
		Email     string `json:"email" binding:"required"`
		Password  string `json:"password" binding:"required"`
		Name      string `json:"name"`
		Cpf       string `json:"cpf"`
		Celphone  string `json:"celphone"`
		LastName  string `json:"lastname"`
		Address   string `json:"address"`
		Number    int32  `json:"number"`
		City      string `json:"city"`
		State     string `json:"state"`
		UserAgent string `json:"useragent"`
	}
	var info signupInfo
	fmt.Println("INFO ->>> ", info)
	if err := c.ShouldBindJSON(&info); err != nil {
		c.AbortWithStatusJSON(402, gin.H{"error": "Please input all fields"})
		return
	}
	user := entity.User{}
	user.Email = info.Email
	hash, err := bcrypt.GenerateFromPassword([]byte(info.Password), bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
		return
	}

	user.Password = string(hash)
	user.Name = info.Name
	user.LastName = info.LastName
	user.Address = info.Address
	user.Number = info.Number
	user.City = info.City
	user.State = info.State
	user.Celphone = info.Celphone
	user.Cpf = info.Cpf
	user.UserAgent = c.GetHeader("User-Agent")
	userservice := service.Userservice{}
	err = userservice.Create(&user)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"result": "ok"})
	}
	return
}
