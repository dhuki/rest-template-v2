package utils

import (
	"bytes"
	"context"
	"fmt"
	"net/smtp"
	"text/template"

	"github.com/dhuki/rest-template-v2/common"
	"github.com/dhuki/rest-template-v2/pkg/testing/domain/entity"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type Email interface {
	SendEmail(ctx context.Context, data entity.TestTable) error
}

type emailUtil struct {
	smtpAuth smtp.Auth
	logger   log.Logger
}

func NewEmailUtil(smtpAuth smtp.Auth, logger log.Logger) Email {
	return emailUtil{
		smtpAuth: smtpAuth,
		logger:   logger,
	}
}

// parse html
func (s emailUtil) parseTemplate(templateFile string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templateFile)
	if err != nil {
		return "", err
	}

	buffer := new(bytes.Buffer)
	if err = t.Execute(buffer, data); err != nil {
		return "", err
	}

	return buffer.String(), nil
}

// making content
func (s emailUtil) makeContentMsg(data entity.TestTable) ([]byte, error) {
	// this is takes two argument, template html path and placeholder for dynamic content
	// that initially with "{{ Name }}".
	message, err := s.parseTemplate("file_template_email.html",
		struct {
			Name string
		}{
			Name: data.Name,
		})

	if err != nil {
		return nil, err
	}

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	to := "To:" + "something@gmail.com"
	subject := "Subject: " + "Something\n"
	body := []byte(to + "\n" + subject + mime + "\n" + message)

	return body, nil
}

// send email
func (s emailUtil) SendEmail(ctx context.Context, data entity.TestTable) error {
	// google smtp address with port
	smtpAddr := fmt.Sprint(common.EmailSmtpHost, ":", common.EmailSmtpPort)
	msg, err := s.makeContentMsg(data)
	if err != nil {
		return level.Error(s.logger).Log("description", "send email", "message", err)
	}

	// receiver email but it will insert it to bcc (means he get copy but another recipent doens't know)
	// so it must use "TO" in body message
	to := []string{"something@gmail.com"}

	// sendMail function
	err = smtp.SendMail(smtpAddr, s.smtpAuth, common.EmailUsername, to, msg)
	if err != nil {
		return level.Error(s.logger).Log("description", "send email", "message", err)
	}

	return nil
}
