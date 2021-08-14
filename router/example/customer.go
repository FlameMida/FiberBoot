package example

import (
	v1 "FiberBoot/api/v1"
	"FiberBoot/middleware"
	"github.com/gofiber/fiber/v2"
)

type CustomerRouter struct {
}

func (e *CustomerRouter) InitCustomerRouter(Router fiber.Router) {
	customerRouter := Router.Group("customer").Use(middleware.Operations())
	var customerApi = v1.AppApi.ExampleApi.CustomerApi
	{
		customerRouter.Post("customer", customerApi.CreateCustomer)     // 创建客户
		customerRouter.Put("customer", customerApi.UpdateCustomer)      // 更新客户
		customerRouter.Delete("customer", customerApi.DeleteCustomer)   // 删除客户
		customerRouter.Get("customer", customerApi.GetCustomer)         // 获取单一客户信息
		customerRouter.Get("customerList", customerApi.GetCustomerList) // 获取客户列表
	}
}
