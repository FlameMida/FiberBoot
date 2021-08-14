package system

import (
	"FiberBoot/api/v1"
	"FiberBoot/middleware"
	"github.com/gofiber/fiber/v2"
)

type UserRouter struct {
}

func (s *UserRouter) InitUserRouter(Router fiber.Router) {
	userRouter := Router.Group("user").Use(middleware.Operations())
	var baseApi = v1.AppApi.SystemApi.Base
	{
		userRouter.Post("register", baseApi.Register)                     // 用户注册账号
		userRouter.Post("changePassword", baseApi.ChangePassword)         // 用户修改密码
		userRouter.Post("getUserList", baseApi.GetUserList)               // 分页获取用户列表
		userRouter.Post("setUserAuthority", baseApi.SetUserAuthority)     // 设置用户权限
		userRouter.Delete("deleteUser", baseApi.DeleteUser)               // 删除用户
		userRouter.Put("setUserInfo", baseApi.SetUserInfo)                // 设置用户信息
		userRouter.Post("setUserAuthorities", baseApi.SetUserAuthorities) // 设置用户权限组
		userRouter.Get("getUserInfo", baseApi.GetUserInfo)                // 获取自身信息
	}
}
