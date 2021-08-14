package middleware

import (
	"FiberBoot/global"
	"FiberBoot/model/common/response"
	"github.com/gofiber/fiber/v2"
)

// NeedInit 处理跨域请求,支持options访问
func NeedInit() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if global.DB == nil {
			return response.OkWithDetailed(fiber.Map{
				"needInit": true,
			}, "前往初始化数据库", c)
		} else {
			return c.Next()
		}
		// 处理请求
	}
}
