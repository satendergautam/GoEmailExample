package controllers

import (
	"main/models"
	"main/services"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// Cors is for allowing request from different Origin.
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

// EmailAPI is for Sending Email to client
func EmailAPI(c *gin.Context) {
	var p models.Email
	var emailRequest models.EmailRequest
	if err := c.BindJSON(&emailRequest); err != nil {
		return
	}
	p.To = emailRequest.To
	p.From = os.Getenv("EMAIL_FROM")
	p.Subject = emailRequest.Subject
	p.Body = emailRequest.Body
	services.EMailSend(&p)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})

}
