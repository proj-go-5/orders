package mail

import (
	"fmt"
	"orders/internal/config"
	"orders/internal/models"
	"strconv"

	"gopkg.in/gomail.v2"
)

type EmailSender struct {
	From string
}

func NewEmailSender() *EmailSender {
	from := "mailtrap@demomailtrap.com"
	return &EmailSender{
		From: from,
	}
}

func (e *EmailSender) SendEmail(o *models.Order) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", "mailtrap@demomailtrap.com")
	msg.SetHeader("To", o.CustomerInfo.Email)
	msg.SetHeader("Subject", fmt.Sprintf("Order #%d", o.ID))
	msg.SetBody("text/html", "Thank you for your order")

	port, err := strconv.Atoi(config.Env("MAIL_PORT"))
	if err != nil {
		return err
	}
	n := gomail.NewDialer(config.Env("MAIL_HOST"), port, config.Env("MAIL_USERNAME"), config.Env("MAIL_PASSWORD"))

	if err := n.DialAndSend(msg); err != nil {
		return err
	}
	return nil
}
