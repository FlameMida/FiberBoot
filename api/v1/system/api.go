package system

import (
	"FiberBoot/global"
	"FiberBoot/model/common/request"
	"FiberBoot/model/common/response"
	"FiberBoot/model/system"
	systemReq "FiberBoot/model/system/request"
	systemRes "FiberBoot/model/system/response"
	"FiberBoot/utils"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Api struct {
}

// CreateApi
//
// @Tags Api
// @Summary 创建基础api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.Api true "api路径, api中文描述, api组, 方法"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /api/createApi [post]
func (s *Api) CreateApi(c *fiber.Ctx) error {
	var api system.Api
	_ = c.BodyParser(&api)
	if err := utils.Verify(api, utils.ApiVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err := apiService.CreateApi(api); err != nil {
		global.LOG.Error("创建失败!", zap.Any("err", err))
		return response.FailWithMessage("创建失败", c)
	} else {
		return response.OkWithMessage("创建成功", c)
	}
}

// DeleteApi
// @Tags Api
// @Summary 删除api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.Api true "ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /api/deleteApi [post]
func (s *Api) DeleteApi(c *fiber.Ctx) error {
	var api system.Api
	_ = c.BodyParser(&api)
	if err := utils.Verify(api.MODEL, utils.IdVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err := apiService.DeleteApi(api); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		return response.FailWithMessage("删除失败", c)
	} else {
		return response.OkWithMessage("删除成功", c)
	}
}

// GetApiList
// @Tags Api
// @Summary 分页获取API列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body systemReq.SearchApiParams true "分页获取API列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getApiList [post]
func (s *Api) GetApiList(c *fiber.Ctx) error {
	var pageInfo systemReq.SearchApiParams
	_ = c.BodyParser(&pageInfo)
	if err := utils.Verify(pageInfo.PageInfo, utils.PageInfoVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err, list, total := apiService.GetAPIInfoList(pageInfo.Api, pageInfo.PageInfo, pageInfo.OrderKey, pageInfo.Desc); err != nil {
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

// GetApiById
// @Tags Api
// @Summary 根据id获取api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "根据id获取api"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getApiById [post]
func (s *Api) GetApiById(c *fiber.Ctx) error {
	var idInfo request.GetById
	_ = c.BodyParser(&idInfo)
	if err := utils.Verify(idInfo, utils.IdVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	err, api := apiService.GetApiById(idInfo.ID)
	if err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		return response.FailWithMessage("获取失败", c)
	} else {
		return response.OkWithData(systemRes.SysAPIResponse{Api: api}, c)
	}
}

// UpdateApi
// @Tags Api
// @Summary 创建基础api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.Api true "api路径, api中文描述, api组, 方法"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /api/updateApi [post]
func (s *Api) UpdateApi(c *fiber.Ctx) error {
	var api system.Api
	_ = c.BodyParser(&api)
	if err := utils.Verify(api, utils.ApiVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err := apiService.UpdateApi(api); err != nil {
		global.LOG.Error("修改失败!", zap.Any("err", err))
		return response.FailWithMessage("修改失败", c)
	} else {
		return response.OkWithMessage("修改成功", c)
	}
}

// GetAllApis
// @Tags Api
// @Summary 获取所有的Api 不分页
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getAllApis [post]
func (s *Api) GetAllApis(c *fiber.Ctx) error {
	if err, apis := apiService.GetAllApis(); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		return response.FailWithMessage("获取失败", c)
	} else {
		return response.OkWithDetailed(systemRes.SysAPIListResponse{Apis: apis}, "获取成功", c)
	}
}

// DeleteApisByIds
// @Tags Api
// @Summary 删除选中Api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /api/deleteApisByIds [delete]
func (s *Api) DeleteApisByIds(c *fiber.Ctx) error {
	var ids request.IdsReq
	_ = c.BodyParser(&ids)
	if err := apiService.DeleteApisByIds(ids); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		return response.FailWithMessage("删除失败", c)
	} else {
		return response.OkWithMessage("删除成功", c)
	}
}
