package main

import (
	"bytes"
	"log"
	"net/smtp"
	"text/template"
)

func sendMailSimpleHtml(subject, templatePath string, receivers []string) {
	// prepare the mail html template
	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Println("Failed to parse mail template (html)")
	}
	// mail data and authentication
	from := "contact@myapp.com"
	t.Execute(&body, struct{ Name string }{Name: "Bryan kouhadi"})
	auth := smtp.PlainAuth("", from, "", "localhost") // identity, senderEmail, password,serverName
	// headers := "MIME-version:1.0, \nContent-Type:text/html; charset=\"UTF-8\";"
	msg := "Subject: " + subject + "\n" + "\n\n" + body.String()
	err = smtp.SendMail("localhost:1025", auth, from, receivers, []byte(msg))
	if err != nil {
		log.Println("Failed to send the mail")
	}
}

func main() {
	sendMailSimpleHtml("Notification of registragion Email", "./mail-template.html", []string{"receiver1@k.k", "receiver2@k.k"})
}
