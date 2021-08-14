package system

import (
	"FiberBoot/global"
	"FiberBoot/model/common/request"
	"FiberBoot/model/common/response"
	"FiberBoot/model/system"
	systemReq "FiberBoot/model/system/request"
	"FiberBoot/utils"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Operations struct {
}

// CreateOperations
//
// @Tags Operations
// @Summary 创建Operations
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.Operations true "创建Operations"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /operations/createOperations [post]
func (s *Operations) CreateOperations(c *fiber.Ctx) error {
	var operations system.Operations
	_ = c.BodyParser(&operations)
	if err := operationsService.CreateOperations(operations); err != nil {
		global.LOG.Error("创建失败!", zap.Any("err", err))
		return response.FailWithMessage("创建失败", c)
	} else {
		return response.OkWithMessage("创建成功", c)
	}
}

// DeleteOperations
//
// @Tags Operations
// @Summary 删除Operations
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.Operations true "Operations模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /operations/deleteOperations [delete]
func (s *Operations) DeleteOperations(c *fiber.Ctx) error {
	var operations system.Operations
	_ = c.BodyParser(&operations)
	if err := operationsService.DeleteOperations(operations); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		return response.FailWithMessage("删除失败", c)
	} else {
		return response.OkWithMessage("删除成功", c)
	}
}

// DeleteOperationsByIds
// @Tags Operations
// @Summary 批量删除Operations
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Operations"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /operations/deleteOperationsByIds [delete]
func (s *Operations) DeleteOperationsByIds(c *fiber.Ctx) error {
	var IDS request.IdsReq
	_ = c.BodyParser(&IDS)
	if err := operationsService.DeleteOperationsByIds(IDS); err != nil {
		global.LOG.Error("批量删除失败!", zap.Any("err", err))
		return response.FailWithMessage("批量删除失败", c)
	} else {
		return response.OkWithMessage("批量删除成功", c)
	}
}

// FindOperations
//
// @Tags Operations
// @Summary 用id查询Operations
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.Operations true "ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /operations/findOperations [get]
func (s *Operations) FindOperations(c *fiber.Ctx) error {
	var operations system.Operations
	_ = c.QueryParser(&operations)
	if err := utils.Verify(operations, utils.IdVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err, reOperations := operationsService.GetOperations(operations.ID); err != nil {
		global.LOG.Error("查询失败!", zap.Any("err", err))
		return response.FailWithMessage("查询失败", c)
	} else {
		return response.OkWithDetailed(fiber.Map{"reOperations": reOperations}, "查询成功", c)
	}
}

// GetOperationsList
//
// @Tags Operations
// @Summary 分页获取Operations列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.OperationsSearch true "页码, 每页大小, 搜索条件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /operations/getOperationsList [get]
func (s *Operations) GetOperationsList(c *fiber.Ctx) error {
	var pageInfo systemReq.OperationsSearch
	_ = c.QueryParser(&pageInfo)
	if err, list, total := operationsService.GetOperationsInfoList(pageInfo); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		return response.FailWithMessage("获取失败", c)
	} else {
		return response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
