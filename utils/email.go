package utils

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"strings"

	"FiberBoot/global"

	"github.com/jordan-wright/email"
)

//@author: Flame
//@function: Email
//@description: Email发送方法
//@param: subject string, body string
//@return: error

func Email(subject string, body string) error {
	to := strings.Split(global.CONFIG.Email.To, ",")
	return send(to, subject, body)
}

//@author: Flame
//@function: ErrorToEmail
//@description: 给email中间件错误发送邮件到指定邮箱
//@param: subject string, body string
//@return: error

func ErrorToEmail(subject string, body string) error {
	to := strings.Split(global.CONFIG.Email.To, ",")
	if to[len(to)-1] == "" { // 判断切片的最后一个元素是否为空,为空则移除
		to = to[:len(to)-1]
	}
	return send(to, subject, body)
}

//@author: Flame
//@function: EmailTest
//@description: Email测试方法
//@param: subject string, body string
//@return: error

func EmailTest(subject string, body string) error {
	to := []string{global.CONFIG.Email.From}
	return send(to, subject, body)
}

//@author: Flame
//@function: send
//@description: Email发送方法
//@param: subject string, body string
//@return: error

func send(to []string, subject string, body string) error {
	from := global.CONFIG.Email.From
	nickname := global.CONFIG.Email.Nickname
	secret := global.CONFIG.Email.Secret
	host := global.CONFIG.Email.Host
	port := global.CONFIG.Email.Port
	isSSL := global.CONFIG.Email.IsSSL

	auth := smtp.PlainAuth("", from, secret, host)
	e := email.NewEmail()
	if nickname != "" {
		e.From = fmt.Sprintf("%s <%s>", nickname, from)
	} else {
		e.From = from
	}
	e.To = to
	e.Subject = subject
	e.HTML = []byte(body)
	var err error
	hostAddr := fmt.Sprintf("%s:%d", host, port)
	if isSSL {
		err = e.SendWithTLS(hostAddr, auth, &tls.Config{ServerName: host})
	} else {
		err = e.Send(hostAddr, auth)
	}
	return err
}
