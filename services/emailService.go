package services

import (
	"fmt"
	"log"
	"stori/config"
)

type EmailService struct {
	host string
	port string
}

func NewEmailService(emailConfig config.EmailConfig) *EmailService {
	return &EmailService{
		host: emailConfig.Host,
		port: emailConfig.Port,
	}
}

func (es *EmailService) Send(from string, password string, email string, summary string) {
	log.Println()
	log.Println("Creating Message...")

	to := []string{email}
	log.Printf("Message to %s \n", to)
	log.Println("--------------------")

	message := []byte(summary)

	log.Printf(string(message))
	log.Println("--------------------")

	log.Println("Sending.....")

	/*
		// Authentication.
		auth := smtp.PlainAuth("", from, password, es.host)

		// Sending email.
		err := smtp.SendMail(es.host+":"+es.port, auth, from, to, message)
		if err != nil {
			fmt.Println(err)
			return
		}
	*/

	fmt.Println("Email Sent Successfully!")
}
