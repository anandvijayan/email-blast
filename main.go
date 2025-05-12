package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/c-bata/go-prompt"

	sendmail "github.com/anandvijayan/email-blast/send_mail"
	notify "github.com/anandvijayan/email-blast/util/notify"
)

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "proceed", Description: "Proceed and send blast emails"},
		{Text: "cancel", Description: "Cancel the blast emails"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func main() {
	testOnly := os.Getenv("TEST_ONLY")
	if testOnly != "" {
		notify.Test()
		return
	}
	notify.LoadEmailBlastConfig()
	message_content := notify.GetMessageContent()
	enhance_message := os.Getenv("ENHANCE_MESSAGE")
	if strings.ToUpper(enhance_message) == "YES" {
		message_content = notify.GetMessageContentEnhanced()
	}
	println(message_content)
	fmt.Println("Please select.")
	choice := prompt.Input("> ", completer)
	fmt.Println("You selected: " + choice)
	if choice == "proceed" {
		sendmail.SendBlast(message_content)
		return
	}
	fmt.Println("Cancelling blast emails..")
}
