package mail

import (
	"os"

	notify "github.com/anandvijayan/email-blast/util/notify"
	"gopkg.in/gomail.v2"
)

func SendBlast(message_content string) {
	m := gomail.NewMessage()
	email_from := os.Getenv("EMAIL_ID")
	email_password := os.Getenv("EMAIL_PASSWORD")
	for _, recipient := range notify.GetRecipients() {
		println("Emailing to ", recipient.Name+"<"+recipient.Email+">")
		email_to := recipient.Email
		email_subject := notify.GetSubject()
		email_body := notify.GenerateMessageBody(recipient, message_content)
		m.SetHeader("From", email_from)
		m.SetHeader("To", email_to)
		m.SetHeader("Subject", email_subject)
		m.SetBody("text/plain", email_body)
		//m.Attach("../reports/dcp_report.png")
		d := gomail.NewDialer("smtp.gmail.com", 587, email_from, email_password)

		if err := d.DialAndSend(m); err != nil {
			panic(err)
		}
	}
}
