package services

import (
	"crypto/tls"
	"fmt"
	"log"
	"main/models"
	"os"
	"strconv"

	"github.com/go-gomail/gomail"
)

// EMailSend is genric function for sending mail
func EMailSend(p *models.Email) {
	log.Println("Inside Email service")
	m := gomail.NewMessage()
	ctype := "text/plain"
	if p.IsHTML {
		ctype = "text/html"
	}
	// Set E-Mail sender
	m.SetHeader("From", p.From)

	// Set E-Mail receivers
	m.SetHeader("To", p.To)

	// Set E-Mail subject
	m.SetHeader("Subject", p.Subject)

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody(ctype, p.Body)
	port, _ := strconv.Atoi(os.Getenv("SMTP_SERVER_PORT"))
	// Settings for SMTP server
	d := gomail.NewDialer(os.Getenv("SMTP_SERVER_HOST"), port, p.From, os.Getenv("EMAIL_PASSWORD"))

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)

	}

}
