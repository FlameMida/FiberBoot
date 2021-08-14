package system

import (
	"FiberBoot/api/v1"
	"FiberBoot/middleware"
	"github.com/gofiber/fiber/v2"
)

type CasbinRouter struct {
}

func (s *CasbinRouter) InitCasbinRouter(Router fiber.Router) {
	casbinRouter := Router.Group("casbin").Use(middleware.Operations())
	var casbinApi = v1.AppApi.SystemApi.Casbin
	{
		casbinRouter.Post("updateCasbin", casbinApi.UpdateCasbin)
		casbinRouter.Post("getPolicyPathByAuthorityId", casbinApi.GetPolicyPathByAuthorityId)
	}
}
