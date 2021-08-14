package system

import (
	"FiberBoot/api/v1"
	"FiberBoot/middleware"
	"github.com/gofiber/fiber/v2"
)

type AuthorityRouter struct {
}

func (s *AuthorityRouter) InitAuthorityRouter(Router fiber.Router) {
	authorityRouter := Router.Group("authority").Use(middleware.Operations())
	var authorityApi = v1.AppApi.SystemApi.Authority
	{
		authorityRouter.Post("createAuthority", authorityApi.CreateAuthority)   // 创建角色
		authorityRouter.Post("deleteAuthority", authorityApi.DeleteAuthority)   // 删除角色
		authorityRouter.Put("updateAuthority", authorityApi.UpdateAuthority)    // 更新角色
		authorityRouter.Post("copyAuthority", authorityApi.CopyAuthority)       // 更新角色
		authorityRouter.Post("getAuthorityList", authorityApi.GetAuthorityList) // 获取角色列表
		authorityRouter.Post("setDataAuthority", authorityApi.SetDataAuthority) // 设置角色资源权限
	}
}
