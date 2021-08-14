package system

import (
	"FiberBoot/utils"
)

type EmailService struct {
}

//@author: Flame
//@function: EmailTest
//@description: 发送邮件测试
//@return: err error

func (e *EmailService) EmailTest() (err error) {
	subject := "test"
	body := "test"
	err = utils.EmailTest(subject, body)
	return err
}
