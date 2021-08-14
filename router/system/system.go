package system

import (
	"FiberBoot/api/v1"
	"FiberBoot/middleware"
	"github.com/gofiber/fiber/v2"
)

type SysRouter struct {
}

func (s *SysRouter) InitSystemRouter(Router fiber.Router) {
	sysRouter := Router.Group("system").Use(middleware.Operations())
	var systems = v1.AppApi.SystemApi.Systems
	{
		sysRouter.Post("getSystemConfig", systems.GetSystemConfig) // 获取配置文件内容
		sysRouter.Post("setSystemConfig", systems.SetSystemConfig) // 设置配置文件内容
		sysRouter.Post("getServerInfo", systems.GetServerInfo)     // 获取服务器信息
		sysRouter.Post("reloadSystem", systems.ReloadSystem)       // 重启服务
	}
}
