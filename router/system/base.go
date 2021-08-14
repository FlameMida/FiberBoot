package system

import (
	"FiberBoot/api/v1"
	"github.com/gofiber/fiber/v2"
)

type BaseRouter struct {
}

func (s *BaseRouter) InitBaseRouter(Router fiber.Router) (R fiber.Router) {
	baseRouter := Router.Group("base")
	var baseApi = v1.AppApi.SystemApi.Base
	{
		baseRouter.Post("login", baseApi.Login)
		baseRouter.Post("captcha", baseApi.Captcha)
	}
	return baseRouter
}
