package main

import (
	"fmt"
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

// SendMailHandler is the handler that we hit when we wanna send an email
func SendMailHandler(w http.ResponseWriter, req *http.Request) {
	name := "Bakr Kouhadi"                        // a name
	someDate := "Friday, 12th july"               // dummy date
	anotherDate := "next Sunday 17 august - 2023" // dummy date
	mailContent := fmt.Sprintf(`
	<strong>Your reservation Confirmation</strong> <br>
	Dear %s, <br>
	Your Reservation from %s to %s  has been done succefully <br>
	Regards  <br>
	`, name, someDate, anotherDate)

	msg := MailData{
		To:       "receiver-bryan-kouhadi@example.com",
		From:     "sender-kouhadi@example.com",
		Subject:  "Confirming your reservation",
		Content:  mailContent,
		Template: "mail-template.html", // or any name of your Mail template
	}
	app.MailsChannel <- msg
	fmt.Fprintf(w, "SENDING MAILS !: Install mailhog and run paste this link on your browser: http://localhost:8025/")
}
