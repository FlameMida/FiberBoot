package example

import (
	"FiberBoot/global"
	"FiberBoot/model/common/request"
	"FiberBoot/model/common/response"
	"FiberBoot/model/example"
	exampleRes "FiberBoot/model/example/response"
	"FiberBoot/utils"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type CustomerApi struct {
}

// CreateCustomer
//
// @Tags Customer
// @Summary 创建客户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body example.Customer true "客户用户名, 客户手机号码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /customer/customer [post]
func (e *CustomerApi) CreateCustomer(c *fiber.Ctx) error {
	var customer example.Customer
	_ = c.BodyParser(&customer)
	if err := utils.Verify(customer, utils.CustomerVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	customer.UserID = utils.GetUserID(c)
	customer.UserAuthorityID = utils.GetUserAuthorityId(c)
	if err := customerService.CreateCustomer(customer); err != nil {
		global.LOG.Error("创建失败!", zap.Any("err", err))
		return response.FailWithMessage("创建失败", c)
	} else {
		return response.OkWithMessage("创建成功", c)
	}
}

// DeleteCustomer
// @Tags Customer
// @Summary 删除客户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body example.Customer true "客户ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /customer/customer [delete]
func (e *CustomerApi) DeleteCustomer(c *fiber.Ctx) error {
	var customer example.Customer
	_ = c.BodyParser(&customer)
	if err := utils.Verify(customer.MODEL, utils.IdVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err := customerService.DeleteCustomer(customer); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		return response.FailWithMessage("删除失败", c)
	} else {
		return response.OkWithMessage("删除成功", c)
	}
}

// UpdateCustomer
// @Tags Customer
// @Summary 更新客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body example.Customer true "客户ID, 客户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /customer/customer [put]
func (e *CustomerApi) UpdateCustomer(c *fiber.Ctx) error {
	var customer example.Customer
	_ = c.BodyParser(&customer)
	if err := utils.Verify(customer.MODEL, utils.IdVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err := utils.Verify(customer, utils.CustomerVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err := customerService.UpdateCustomer(&customer); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		return response.FailWithMessage("更新失败", c)
	} else {
		return response.OkWithMessage("更新成功", c)
	}
}

// GetCustomer
// @Tags Customer
// @Summary 获取单一客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body example.Customer true "客户ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/customer [get]
func (e *CustomerApi) GetCustomer(c *fiber.Ctx) error {
	var customer example.Customer
	_ = c.QueryParser(&customer)
	if err := utils.Verify(customer.MODEL, utils.IdVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	err, data := customerService.GetCustomer(customer.ID)
	if err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		return response.FailWithMessage("获取失败", c)
	} else {
		return response.OkWithDetailed(exampleRes.CustomerResponse{Customer: data}, "获取成功", c)
	}
}

// GetCustomerList
// @Tags Customer
// @Summary 分页获取权限客户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/customerList [get]
func (e *CustomerApi) GetCustomerList(c *fiber.Ctx) error {
	var pageInfo request.PageInfo
	_ = c.QueryParser(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	err, customerList, total := customerService.GetCustomerInfoList(utils.GetUserAuthorityId(c), pageInfo)
	if err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		return response.FailWithMessage("获取失败"+err.Error(), c)
	} else {
		return response.OkWithDetailed(response.PageResult{
			List:     customerList,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
