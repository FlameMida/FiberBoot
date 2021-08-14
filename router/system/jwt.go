package system

import (
	"FiberBoot/api/v1"
	"FiberBoot/middleware"
	"github.com/gofiber/fiber/v2"
)

type JwtRouter struct {
}

func (s *JwtRouter) InitJwtRouter(Router fiber.Router) {
	jwtRouter := Router.Group("jwt").Use(middleware.Operations())
	var jwtApi = v1.AppApi.SystemApi.Jwt
	{
		jwtRouter.Post("jsonInBlacklist", jwtApi.JsonInBlacklist) // jwt加入黑名单
	}
}
