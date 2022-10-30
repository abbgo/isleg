package controllers

import (
	"fmt"
	"net/http"
	"net/smtp"
	"os"

	"github.com/gin-gonic/gin"
)

type ForMail struct {
	FullName    string `json:"full_name" binding:"required,min=3"`
	Email       string `json:"email" binding:"email"`
	PhoneNumber string `json:"phone_number" binding:"required,e164,len=12"`
	Letter      string `json:"letter" binding:"required,min=3"`
}

func SendMail(c *gin.Context) {

	var mail ForMail

	if err := c.BindJSON(&mail); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	from := os.Getenv("MAIL_FROM")
	password := os.Getenv("MAIL_PASSWORD")
	toEmailAddress := os.Getenv("MAIL_TO")
	to := []string{toEmailAddress}

	host := os.Getenv("MAIL_HOST")
	port := os.Getenv("MAIL_PORT")
	address := host + ":" + port

	subject := "Subject: " + mail.FullName + " - den hat geldi\n"
	letter := fmt.Sprintf("Mail adresi: %s\r\n", mail.Email)
	letter += fmt.Sprintf("Telefon belgisi: %s\r\n", mail.PhoneNumber)
	letter += fmt.Sprintf("Haty: %s\r\n", mail.Letter)
	body := fmt.Sprintf("\r\n%s\r\n", letter)
	message := []byte(subject + body)

	auth := smtp.PlainAuth("", from, password, host)
	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "mail successfully send",
	})

}
