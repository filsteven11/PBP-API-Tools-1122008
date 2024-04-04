package Controller

import (
	"fmt"
	"net/smtp"
	"strings"
)

const (
	DefaultSMTPPort = "587"
)

type EmailController struct {
	From     string
	Password string
	SMTPAddr string
}

// NewEmailController membuat instance baru dari EmailController
func NewEmailController(from, password, smtpAddr string) (*EmailController, error) {
	if from == "" || password == "" || smtpAddr == "" {
		return nil, fmt.Errorf("from, password, dan smtpAddr tidak boleh kosong")
	}

	return &EmailController{
		From:     from,
		Password: password,
		SMTPAddr: smtpAddr,
	}, nil
}

// SendEmail mengirim email dengan menggunakan SMTP
func (ec *EmailController) SendEmail(to, subject, body string) error {
	if to == "" || subject == "" || body == "" {
		return fmt.Errorf("to, subject, dan body tidak boleh kosong")
	}

	auth := smtp.PlainAuth("", ec.From, ec.Password, strings.Split(ec.SMTPAddr, ":")[0])

	msg := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s", ec.From, to, subject, body)

	err := smtp.SendMail(ec.SMTPAddr, auth, ec.From, []string{to}, []byte(msg))
	if err != nil {
		return fmt.Errorf("gagal mengirim email: %v", err)
	}

	return nil
}
