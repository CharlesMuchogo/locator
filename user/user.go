package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"main.go/databasehandler"
	"main.go/myStructs"
	"net/http"
)

func BeforeSave(password string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	//remove spaces in username
	//u.Username = .EscapeString(strings.TrimSpace(u.Username))

	return hashedPassword, nil
}

func Register(c *gin.Context) {
	var input myStructs.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	password, er := BeforeSave(input.Password)

	fmt.Printf(" input email address: %s \n ", input.Email)
	fmt.Printf(" input first name: %s \n ", input.First_name)

	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": er.Error(),
		})
		return
	}

	user, status, err := databasehandler.SaveUser(input.First_name, input.Middle_name, input.Email, input.Firebase_id, input.Phone_number, password)

	if status != 200 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Registration successful",
			"user":    user,
		})
	}

}

func Login(c *gin.Context) {
	var loginData myStructs.LoginData
	if err := c.ShouldBindJSON(&loginData); err != nil {
		fmt.Printf("error: %s \n ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	fmt.Printf("error: %s \n ", loginData.Firebase_id)
	password, err := BeforeSave(loginData.Password)
	user, err := databasehandler.Login(loginData.Email, password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "login success",
			"user":    user})
	}

}
