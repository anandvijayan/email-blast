package mail

import (
	"os"
	"time"

	notify "github.com/anandvijayan/email-blast/util/notify"
	"gopkg.in/gomail.v2"
)

func SendBlast(message_content string) {
	m := gomail.NewMessage()
	email_from := os.Getenv("EMAIL_ID")
	email_user := os.Getenv("EMAIL_USER")
	email_password := os.Getenv("EMAIL_PASSWORD")
	for _, recipient := range notify.GetRecipients() {
		time.Sleep(2 * time.Second)
		println("Emailing to ", recipient.Name+"<"+recipient.Email+">")
		email_to := recipient.Email
		email_subject := notify.GetSubject()
		email_body := notify.GenerateMessageBody(recipient, message_content)
		m.SetHeader("From", email_from)
		m.SetHeader("To", email_to)
		m.SetHeader("Subject", email_subject)
		m.SetBody("text/html", email_body)
		for _, attachment := range notify.GetAttachments() {
			m.Attach(attachment)
		}
		d := gomail.NewDialer("email-smtp.us-west-2.amazonaws.com", 587, email_user, email_password)

		if err := d.DialAndSend(m); err != nil {
			panic(err)
		}
	}
}
