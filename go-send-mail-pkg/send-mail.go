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
			SendMail(msg)
		}
	}()
}
func SendMail(msg MailData) {
	server := mail.NewSMTPClient()
	server.Username = "test@example.com"
	server.Password = "examplepass"
	server.Encryption = mail.EncryptionSTARTTLS
	server.Host = "localhost" // or "smtp.example.com"
	server.Port = 1025        // or 587 as in the docs
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	client, err := server.Connect()
	if err != nil {
		fmt.Println(err)
	}

	email := mail.NewMSG()
	email.SetFrom(msg.From).AddTo(msg.To).SetSubject(msg.Subject).AddCc(msg.Cc)
	// add inline
	email.Attach(&mail.File{FilePath: "./img.png", Name: "img.png", Inline: true})
	email.Attach(&mail.File{FilePath: "./doc.pdf", Name: "doc.pdf", Inline: true})
	// also you can set Delivery Status Notification (DSN) (only is set when server supports DSN)
	// email.SetDSN([]mail.DSN{mail.SUCCESS, mail.FAILURE}, false)

	// we haven't specified the template do this
	if msg.Template == "" {
		email.SetBody(mail.TextHTML, msg.Content)
	} else {
		// otherwise get the specified template from disk
		data, err := os.ReadFile(fmt.Sprintf("./email-templates/%s", msg.Template))
		if err != nil {
			fmt.Println(err)
		}
		// since data is of type array of bytes ([]bytes) we need to convert it
		mailTmpl := string(data)
		//args:  string to be replace in, txt to be replaced, content, how many times
		msgToSend := strings.Replace(mailTmpl, "[%body%]", msg.Content, 1)
		email.SetBody(mail.TextHTML, msgToSend)
	}

	err = email.Send(client)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("EMAIL HAS BEEN SENT SUCCESFULLY")
	}
}
