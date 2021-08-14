package system

import (
	"FiberBoot/api/v1"
	"FiberBoot/middleware"
	"github.com/gofiber/fiber/v2"
)

type EmailRouter struct {
}

func (s *EmailRouter) InitEmailRouter(Router fiber.Router) {
	emailRouter := Router.Group("email").Use(middleware.Operations())
	var systemApi = v1.AppApi.SystemApi.Systems
	{
		emailRouter.Post("emailTest", systemApi.EmailTest) // 发送测试邮件
	}
}
