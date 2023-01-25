package controllers

import (
	"bytes"
	"html/template"
	"net/http"
	"net/smtp"
	"os"

	"github.com/gin-gonic/gin"
)

var authh smtp.Auth

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

	var from = os.Getenv("MAIL_FROM")
	var password = os.Getenv("MAIL_PASSWORD")
	var toEmailAddress = os.Getenv("MAIL_TO")
	var to = []string{toEmailAddress}

	var host = os.Getenv("MAIL_HOST")
	var serverPath = os.Getenv("SERVER_PATH")

	authh = smtp.PlainAuth("", from, password, host)
	templateData := struct {
		Name        string
		Mail        string
		Message     string
		PhoneNumber string
	}{
		Name:        mail.FullName,
		Mail:        mail.Email,
		Message:     mail.Letter,
		PhoneNumber: mail.PhoneNumber,
	}
	r := NewRequest(to, "Salam. Men "+templateData.Name, "Salam Isleg Market!")
	if err := r.ParseTemplate(serverPath+"templates/template.html", templateData); err == nil {
		_, err := r.SendEmail()
		if err != nil {
			c.JSON(400, gin.H{"message": err.Error()})
			return
		}
	} else {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	// subject := "Subject: " + mail.FullName + " - den hat geldi\n"
	// letter := fmt.Sprintf("Mail adresi: %s\r\n", mail.Email)
	// letter += fmt.Sprintf("Telefon belgisi: %s\r\n", mail.PhoneNumber)
	// letter += fmt.Sprintf("Haty: %s\r\n", mail.Letter)
	// body := fmt.Sprintf("\r\n%s\r\n", letter)
	// message := []byte(subject + body)

	// auth := smtp.PlainAuth("", from, password, host)
	// err := smtp.SendMail(address, auth, from, to, message)
	// if err != nil {
	// 	c.JSON(500, gin.H{"message": err.Error()})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "mail successfully send",
	})

}

// Request struct
type Request struct {
	// from    string
	to      []string
	subject string
	body    string
}

func NewRequest(to []string, subject, body string) *Request {
	return &Request{
		to:      to,
		subject: subject,
		body:    body,
	}
}

func (r *Request) SendEmail() (bool, error) {
	var port = os.Getenv("MAIL_PORT")
	var from = os.Getenv("MAIL_FROM")
	var host = os.Getenv("MAIL_HOST")

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + r.subject + "!\n"
	msg := []byte(subject + mime + "\n" + r.body)
	addr := host + ":" + port

	if err := smtp.SendMail(addr, authh, from, r.to, msg); err != nil {
		return false, err
	}
	return true, nil
}

func (r *Request) ParseTemplate(templateFileName string, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	r.body = buf.String()
	return nil
}
