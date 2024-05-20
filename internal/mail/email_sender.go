package mail

import (
	"bytes"
	"fmt"
	"orders/internal/config"
	"orders/internal/dto"
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

func (e *EmailSender) SendEmail(o *models.Order, products []*dto.Product) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", e.From)
	msg.SetHeader("To", o.CustomerInfo.Email)
	msg.SetHeader("Subject", fmt.Sprintf("Order #%d", o.ID))
	msg.SetBody("text/html", e.buildBodyText(o.Products, products))

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

func (e *EmailSender) buildBodyText(products []models.OrderProduct, productInfo []*dto.Product) string {
	var result bytes.Buffer
	result.WriteString("Thank you for your order.\n")

	for _, p := range products {
		for _, pi := range productInfo {
			if p.ProductID == pi.ID {
				str := fmt.Sprintf("%s: %d\n", pi.Title, p.Quantity)
				result.WriteString(str)
			}
		}
	}
	return result.String()
}
