package user

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"main.go/databasehandler"
	"net/http"
)

type RegisterInput struct {
	Email        string `json:"email" binding:"required"`
	First_name   string `json:"first_name" binding:"required"`
	Middle_name  string `json:"middle_name" binding:"required"`
	Phone_number string `json:"phone_number" binding:"required"`
	Firebase_id  string `json:"firebase_id" binding:"required"`
	Password     string `json:"password" binding:"required"`
}

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
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	password, er := BeforeSave(input.Password)

	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": er,
		})
		return
	}

	_, err := databasehandler.SaveUser(input.First_name, input.Middle_name, input.Email, input.Firebase_id, input.Phone_number, password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Registration successful",
		})
	}

}
