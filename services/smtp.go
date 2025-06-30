package services

import (
	"errors"
	"net/smtp"
	"strings"
	"video-streaming-server/config"
	"video-streaming-server/shared/logger"
)

//SendEmail is a wrapper over smtp.SendEmail.
//
//This function consumes information like host address, username and password from the config.
//
//Recipients is an array of string which is joined using comma (,) 
//separator in the function, creating a comma separated string
func SendEmail(recipients []string, sender string, subject string, body string) error {
	if len(recipients) == 0 {
		return errors.New("Recipients should not be emtpy")
	}

	user := config.AppConfig.SMTPUser
	password := config.AppConfig.SMTPPassword

	host := config.AppConfig.SMTPHost
	port := config.AppConfig.SMTPPort
	addr := host + ":" + port

	auth := smtp.PlainAuth("", user, password, host)

	msg := "From: " + sender + "\r\n" +
	"To: " + strings.Join(recipients, ",") + "\r\n" +
	"Subject: " + subject + "\r\n" + body

	err := smtp.SendMail(addr, auth, user, recipients, []byte(msg))

	if err != nil {
		logger.Log.Error(err.Error())

		return err
	}

	return nil
}
