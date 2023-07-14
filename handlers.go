package main

import (
	"fmt"
	"net/http"
)

// SendMailHandler is the handler that we hit when we wanna send an email
func SendMailHandler(w http.ResponseWriter, req *http.Request) {
	name := "Bakr Kouhadi"      // a name
	someDate := "2023-09-14"    // dummy date
	anotherDate := "2023-11-27" // dummy date
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
