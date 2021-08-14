package system

import (
	"FiberBoot/api/v1"
	"FiberBoot/middleware"
	"github.com/gofiber/fiber/v2"
)

type ApiRouter struct {
}

func (s *ApiRouter) InitApiRouter(Router fiber.Router) {
	apiRouter := Router.Group("api").Use(middleware.Operations())
	var API = v1.AppApi.SystemApi.Api
	{
		apiRouter.Post("createApi", API.CreateApi)               // 创建Api
		apiRouter.Post("deleteApi", API.DeleteApi)               // 删除Api
		apiRouter.Post("getApiList", API.GetApiList)             // 获取Api列表
		apiRouter.Post("getApiById", API.GetApiById)             // 获取单条Api消息
		apiRouter.Post("updateApi", API.UpdateApi)               // 更新api
		apiRouter.Post("getAllApis", API.GetAllApis)             // 获取所有api
		apiRouter.Delete("deleteApisByIds", API.DeleteApisByIds) // 删除选中api
	}
}
