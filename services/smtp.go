package services

import (
	"errors"
	"net/smtp"
	"strings"
	"video-streaming-server/config"
	"video-streaming-server/shared/logger"
)

type message struct {
	sender     string
	recipients []string
	isHTML     bool
	subject    string
	body       string
}

// SendEmail is a wrapper over smtp.SendEmail.
//
// This function consumes information like host address, username and password from the config.
//
// Recipients is an array of string which is joined using comma (,)
// separator in the function, creating a comma separated string
func SendEmail(msg *message) error {
	if len(msg.recipients) == 0 {
		return errors.New("Recipients should not be emtpy")
	}

	user := config.AppConfig.SMTPUser
	password := config.AppConfig.SMTPPassword

	host := config.AppConfig.SMTPHost
	port := config.AppConfig.SMTPPort
	addr := host + ":" + port

	auth := smtp.PlainAuth("", user, password, host)

	message := "From: " + msg.sender + "\r\n" +
		"To: " + strings.Join(msg.recipients, ",") + "\r\n" +
		"Subject: " + msg.subject + "\r\n"

	if msg.isHTML {
		message += "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	}

	message += msg.body

	err := smtp.SendMail(addr, auth, user, msg.recipients, []byte(message))

	if err != nil {
		logger.Log.Error(err.Error())

		return err
	}

	return nil
}
