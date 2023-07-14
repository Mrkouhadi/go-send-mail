### Description

- A guide of how to send mails in go(Golang) efficiently with the help of a very light an ammazing package `go-simple-mail`

### Instructions

- Clone the repository, or download it
- Run : `go run *.go`
- Install Mailhog on your computer and run it.
- Paste this link in the browser: http://localhost:8025/
- Head to this link: http://localhost:8080/send-mail

### How to send a mail with ONLY standard library :

`
import (
"log"
"net/smtp"
)

    func main() {

        from := "me@here.com"

        auth := smtp.PlainAuth("", from, "", "localhost")  // identity, senderEmail, password,serverName


        err := smtp.SendMail("localhost:1025", auth, from, []string{"you@there.com"}, []byte("hello world"))  // adress, auth, senderEmail, recipientsEmails, content
        if err != nil {
            log.Println(err)
        }
    }

`
