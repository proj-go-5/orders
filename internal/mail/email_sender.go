package mail

import "fmt"

type EmailSender struct {
}

func (e *EmailSender) SendEmail(userID int, mailID int) {
	fmt.Printf("EmailSender: Sending email for user %d; Mail ID: %d", userID, mailID)
}
