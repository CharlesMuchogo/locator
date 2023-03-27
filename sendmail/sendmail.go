package sendmail

import (
	"fmt"
	"net/smtp"

	"main.go/databasehandler"
	"main.go/myStructs"
)

func SendMail(recipient string, user myStructs.User) {
	from := "muchpaul2@gmail.com"
	password := "ykynxuelmugnzpat"
	to := recipient

	link := fmt.Sprintf("%s%s",databasehandler.GoDotEnvVariable("ACCEPT_REQUEST_LINK"),user.UserId ) 
	sendBody := fmt.Sprintf("%v %v requests to be admin.\n Click  <a href=\"%v\">here</a> to accept the request ", user.First_name, user.Middle_name, link)

	msg := []byte(fmt.Sprintf("To: %s\r\n", to) +
		"Subject: Request to be Admin\r\n" +
		"Content-Type: text/html\r\n" +
		fmt.Sprintf("\r\n%s", sendBody))

	err := smtp.SendMail("smtp.gmail.com:587",

		smtp.PlainAuth("", from, password, "smtp.gmail.com"),
		from, []string{to}, msg)

	if err != nil {
		fmt.Println("Error sending email: ", err)
		return
	}

	fmt.Println("Email sent successfully.")
}
