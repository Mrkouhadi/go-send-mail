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
`
html:
`

<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>
    <h1>Hello {{.Name}}</h1>
</body>
</html>
`
