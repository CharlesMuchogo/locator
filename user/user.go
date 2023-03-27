package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"main.go/databasehandler"
	"main.go/myStructs"
	"main.go/sendmail"
)

func BeforeSave(password string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hashedPassword, nil
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
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

func UpdateProfile(c *gin.Context) {

	var user myStructs.User

	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Printf("error: %s \n ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	status, response := databasehandler.UpdateProfile(user.Profile_photo, user.UserId)

	userData, getdetailsErr := databasehandler.Login(user.Email)
	if getdetailsErr != nil {
		fmt.Printf("error: %s \n ", getdetailsErr.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Server error"})
		return
	}

	if status == 200 {
		c.JSON(http.StatusOK, gin.H{"message": response, "user": userData})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": response})
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
	user, err := databasehandler.Login(loginData.Email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	} else if user.UserId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Incorrect credentials"})
	} else {
		if VerifyPassword(loginData.Password, user.Password) != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Incorrect credentials"})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "login success",
				"user":    user})
		}
	}

}

func RequestPromotion(c *gin.Context) {
	var requestDetails myStructs.LoginData
	if err := c.ShouldBindJSON(&requestDetails); err != nil {
		fmt.Printf("error: %s \n ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	user, err := databasehandler.Login(requestDetails.Email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	} else if user.UserId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "User not found"})
	} else {
		go sendmail.SendMail("charles5muchogo@gmail.com", user)
		c.JSON(http.StatusOK, gin.H{"message": "request sent successfully"})
	}

}

func PromoteUser(c *gin.Context) {
	id := c.Query("id")
	fmt.Printf("userid is updated successfully %s  \n", id)

	requestStatus := databasehandler.PromoteUser(id, true)
	fmt.Printf("userid is updated successfully %d  \n", requestStatus)

	if requestStatus == 200 {

		c.JSON(http.StatusOK, gin.H{"message": "User has been promoted to admin successfully"})

	} else {
		c.JSON(http.StatusOK, gin.H{"message": "could not promote user"})

	}
}
