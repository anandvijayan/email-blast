package notify

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	util "github.com/newhorizonsarizona/tmi-status-checker/util"
	"gopkg.in/yaml.v3"
)

// Create an instance of the Config struct
var emailBlastConfig EmailBlastConfig

type Recipient struct {
	Name  string `yaml:"Name"`
	Email string `yaml:"Email"`
}

// Define a struct that matches the structure of your YAML file
type EmailBlastConfig struct {
	EmailBlast struct {
		Recipients []Recipient `yaml:"Recepients"`
		Message    struct {
			SalutationPrefix string   `yaml:"SalutationPrefix"`
			MessageTone      string   `yaml:"MessageTone"`
			MessageSubject   string   `yaml:"MessageSubject"`
			MessageContent   string   `yaml:"MessageContent"`
			Attachments      []string `yaml:"Attachments"`
			Signature        string   `yaml:"Signature"`
		} `yaml:"Message"`
	} `yaml:"EmailBlast"`
}

func LoadEmailBlastConfig() {
	// Open the YAML file
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	configPath := filepath.Join(currentDir, "config", "email_blast.yaml")
	file, err := os.Open(configPath)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	defer file.Close()

	// Create a new decoder
	decoder := yaml.NewDecoder(file)

	// Decode the YAML into the struct
	err = decoder.Decode(&emailBlastConfig)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}

func GetRecipients() []Recipient {
	return emailBlastConfig.EmailBlast.Recipients
}

func GetSalutation(recepient Recipient) string {
	firstName, _, _ := strings.Cut(recepient.Name, " ")
	return emailBlastConfig.EmailBlast.Message.SalutationPrefix + " " + firstName
}

func GetSubject() string {
	return emailBlastConfig.EmailBlast.Message.MessageSubject
}

func GetMessageContent() string {
	return emailBlastConfig.EmailBlast.Message.MessageContent
}

func GetMessageContentEnhanced() string {
	messageTone := emailBlastConfig.EmailBlast.Message.MessageTone
	messageContent := emailBlastConfig.EmailBlast.Message.MessageContent
	enhancedMessage := util.Chat(fmt.Sprintf("%s %s", messageTone, messageContent))
	return enhancedMessage
}

func GetAttachments() []string {
	return emailBlastConfig.EmailBlast.Message.Attachments
}

func GetSignature() string {
	return emailBlastConfig.EmailBlast.Message.Signature
}

func GenerateMessageBody(recipient Recipient, messageContent string) string {
	return fmt.Sprintf(`
%s,

%s

%s
`, GetSalutation(recipient), messageContent, GetSignature())
}

func Test() {
	LoadEmailBlastConfig()
	messageContent := GetMessageContentEnhanced()
	for _, recipient := range GetRecipients() {
		println("Emailing to ", recipient.Name+"<"+recipient.Email+">")
		println(GenerateMessageBody(recipient, messageContent))
	}
}
