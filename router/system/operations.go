package system

import (
	"FiberBoot/api/v1"
	"FiberBoot/middleware"
	"github.com/gofiber/fiber/v2"
)

type OperationsRouter struct {
}

func (s *OperationsRouter) InitOperationsRouter(Router fiber.Router) {
	operationsRouter := Router.Group("operations").Use(middleware.Operations())
	var operations = v1.AppApi.SystemApi.Operations
	{
		operationsRouter.Post("createOperations", operations.CreateOperations)             // 新建Operations
		operationsRouter.Delete("deleteOperations", operations.DeleteOperations)           // 删除Operations
		operationsRouter.Delete("deleteOperationsByIds", operations.DeleteOperationsByIds) // 批量删除Operations
		operationsRouter.Get("findOperations", operations.FindOperations)                  // 根据ID获取Operations
		operationsRouter.Get("getOperationsList", operations.GetOperationsList)            // 获取Operations列表

	}
}
