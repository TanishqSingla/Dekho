package services

import (
	"net/smtp"
	"video-streaming-server/config"
	"video-streaming-server/shared/logger"
)

func SendEmail(to []string, body string) error {
	user := config.AppConfig.SMTPUser
	password := config.AppConfig.SMTPPAssword

	host := config.AppConfig.SMTPHost
	port := config.AppConfig.SMTPPort
	addr := host + ":" + port

	auth := smtp.PlainAuth("", user, password, host)

	err := smtp.SendMail(addr, auth, user, to, []byte(body))

	if err != nil {
		logger.Log.Error(err.Error())

		return err
	}

	return nil
}
