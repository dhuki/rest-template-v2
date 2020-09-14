package config

import (
	"net/smtp"

	"github.com/dhuki/rest-template-v2/common"
)

func NewEmail() smtp.Auth {
	// using smtp.gmail.com for server smtp has limit (100 sendMail/day)
	// so by default in centOs7, it provide smtp server external
	// for send email that has not limit.
	// Usually identity should be the empty string, to act as username.
	return smtp.PlainAuth("", common.EmailUsername, common.EmailPassword, common.EmailSmtpHost)
}
