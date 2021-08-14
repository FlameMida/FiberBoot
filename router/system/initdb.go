package system

import (
	"FiberBoot/api/v1"
	"github.com/gofiber/fiber/v2"
)

type InitRouter struct {
}

func (s *InitRouter) InitInitRouter(Router fiber.Router) {
	initRouter := Router.Group("init")
	var dbApi = v1.AppApi.SystemApi.DB
	{
		initRouter.Post("initDB", dbApi.InitDB)   // 创建Api
		initRouter.Post("checkDB", dbApi.CheckDB) // 创建Api
	}
}
