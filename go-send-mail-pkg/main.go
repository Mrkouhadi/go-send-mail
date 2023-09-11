package main

import (
	"net/http"
)

var app AppConfig

func main() {
	// this part is responsible for listening to mails received by our mail Channel
	mailingChan := make(chan MailData)
	app.MailsChannel = mailingChan
	defer close(app.MailsChannel)
	ListenForMails()
	// run our local server
	http.HandleFunc("/send-mail", SendMailHandler)
	http.ListenAndServe(":8080", nil)
}
