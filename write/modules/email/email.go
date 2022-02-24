package email

import (
	"errors"
	"fmt"
	"net/smtp"
	"strings"

	"git.kicoe.com/blog/write/modules/setting"
	"github.com/sirupsen/logrus"
)

type Email struct {
	To      []string
	Subject string
	Body    string
}

var logger *logrus.Logger

func SetLog(l *logrus.Logger) {
	logger = l
}

func (e *Email) Parse() []byte {
	return []byte(
		"To: " + strings.Join(e.To, ",") + "\n" +
			"From: notice<" + setting.SMTP.User + ">\n" +
			"Subject: " + e.Subject + "\n" +
			"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n" +
			e.Body,
	)
}

func (e *Email) Send() error {
	if setting.SMTP.User == "" {
		return errors.New("email, no config")
	}

	logger.WithFields(logrus.Fields{
		"email": e.To,
	}).Info("send")

	return smtp.SendMail(
		fmt.Sprintf("%s:%s", setting.SMTP.Host, setting.SMTP.Port),
		smtp.PlainAuth("", setting.SMTP.User, setting.SMTP.Pass, setting.SMTP.Host),
		setting.SMTP.User,
		e.To,
		e.Parse(),
	)
}
