package main

import (
	"log"
	"net/smtp"
)

func main() {
	// Set up authentication information
	//
	// PlainAuth returns an Auth that implements the PLAIN authentication
	// mechanism as defined in RFC 4616.
	// The returned Auth uses the given username and password to
	// authenticate on TLS connections to host and act as identity.
	// Usually identity will be left blank to act as username.
	auth := smtp.PlainAuth("", "example@gmail.com", "password", "smtp.gmail.com")

	// Connect to the server, authenticate, set the sender and the
	// recipient and then send the email, all in one single step
	//
	// Set up the relevant information needed to send the email
	to := []string{"example@gmail.com"} // TO
	msg := []byte("To: example@gmail.com\r\n" +
		"Subject: Hello From Golang\r\n" +
		"\r\n" +
		"Sending an email using golang over SMTP",
	)

	// Now send the mail using all of the above information
	//
	// SendMail connects to the server at addr, switches to TLS if
	// possible, authenticates with the optional mechanism a if possible,
	// and then sends an email from address from, to addresses to, with
	// message msg.
	// The addr must include a port, as in "mail.example.com:smtp".
	//
	// The addresses in the to parameter are the SMTP RCPT addresses.
	//
	// The msg parameter should be an RFC 822-style email with headers
	// first, a blank line, and then the message body. The lines of msg
	// should be CRLF terminated.  The msg headers should usually include
	// fields such as "From", "To", "Subject", and "Cc".  Sending "Bcc"
	// messages is accomplished by including an email address in the to
	// parameter but not including it in the msg headers.
	//
	// The SendMail function and the the net/smtp package are low-level
	// mechanisms and provide no support for DKIM signing, MIME
	// attachments (see the mime/multipart package), or other mail
	// functionality. Higher-level packages exist outside of the standard
	// library.
	err := smtp.SendMail("smtp.gmail.com:587", auth,
		"example@gmail.com", to, msg)
	if err != nil {
		log.Fatalln(err)
	}
}
