package middleware

import (
	"FiberBoot/global"
	"FiberBoot/model/common/response"
	"FiberBoot/model/system/request"
	"FiberBoot/service"
	"github.com/gofiber/fiber/v2"
)

var casbinService = service.AppService.SystemService.CasbinService

// CasbinHandler 拦截器
func CasbinHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims := c.Locals("claims")
		waitUse := claims.(*request.CustomClaims)
		// 获取请求的URI
		obj := c.OriginalURL()
		// 获取请求方法
		act := c.Method()
		// 获取用户的角色
		sub := waitUse.AuthorityId
		e := casbinService.Casbin()
		// 判断策略中是否存在
		success, _ := e.Enforce(sub, obj, act)
		if global.CONFIG.System.Env == "develop" || success {
			return c.Next()
		} else {
			return response.FailWithDetailed(fiber.Map{}, "权限不足", c)
		}
	}
}
