package main

import (
	"bytes"
	"log"
	"net/smtp"
	"text/template"
)

func sendMailSimpleHtml(subject, templatePath string, receivers []string) {
	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Println("Failed to parse htm mail template")
	}
	t.Execute(&body, struct{ Name string }{Name: "Bryan kouhadi"})
	auth := smtp.PlainAuth("", "kouhadibakr@gmail.com", "", "localhost") // identity, senderEmail, password,serverName
	headers := "MIME-version:1.0, \nContent-Type:text/html; charset=\"UTF-8\";"
	message := "Subject: " + "\n" + headers + "\n\n" + body.String()
	err = smtp.SendMail("localhost:1025", auth, "kouhadibakr@gmail.com", receivers, []byte(message))
	if err != nil {
		log.Println("Failed to send the mail")
	}
}

func main() {
	sendMailSimpleHtml("Notification of registragion Email", "./mail-template.html", []string{"k@k.com"})
}
