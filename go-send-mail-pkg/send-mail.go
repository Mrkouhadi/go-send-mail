package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	mail "github.com/xhit/go-simple-mail/v2"
)

func ListenForMails() {
	// an anonymous function that runs in the background
	go func() {
		for {
			msg := <-app.MailsChannel
			sendMsg(msg)
		}
	}()
}

func sendMsg(msg MailData) {
	server := mail.NewSMTPClient()
	server.Username = "test@example.com"
	server.Password = "examplepass"
	server.Encryption = mail.EncryptionSTARTTLS
	server.Host = "127.0.0.1" // or "smtp.example.com"
	server.Port = 1025        // or 587 as in the docs
	server.KeepAlive = true
	server.ConnectTimeout = 60 * time.Second
	server.SendTimeout = 60 * time.Second

	client, err := server.Connect()
	if err != nil {
		// handle errors
		return
	}
	// Read the email template file outside of the loop
	var templateContent string
	if msg.Template != "" {
		data, err := os.ReadFile(fmt.Sprintf("./email-templates/%s", msg.Template))
		if err != nil {
			// handle errors
			return
		}
		// Convert template data to string
		templateContent = string(data)
	}
	for _, recipient := range msg.To {
		email := mail.NewMSG()
		email.SetFrom(msg.From).AddTo(recipient).AddCc(msg.Cc).SetSubject(msg.Subject)
		// Add inline attachments

		// FIXME: can't add attachements
		// email.Attach(&mail.File{FilePath: "./img.png", Name: "img.png", Inline: true})
		// email.Attach(&mail.File{FilePath: "./doc.pdf", Name: "doc.pdf", Inline: true})

		// 	// also you can set Delivery Status Notification (DSN) (only is set when server supports DSN)
		// email.SetDSN([]mail.DSN{mail.SUCCESS, mail.FAILURE}, false)

		// Set custom headers if needed
		// email.AddHeader("X-Priority", "1") // Example of adding a custom header

		// Set email body using template content or provided content
		if msg.Template == "" {
			email.SetBody(mail.TextHTML, msg.Content)
		} else {
			// Replace template placeholders with actual content
			msgToSend := strings.Replace(templateContent, "[%body%]", msg.Content, 1)
			email.SetBody(mail.TextHTML, msgToSend)
		}
		// Send the email
		err = email.Send(client)
		if err != nil {
			log.Println("error: ", err)
		} else {
			log.Println("EMAIL HAS BEEN SENT SUCCESSFULLY")
		}
	}
}
