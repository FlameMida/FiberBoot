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

type Authority struct {
}

// CreateAuthority
//
// @Tags Authority
// @Summary 创建角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.Authority true "权限id, 权限名, 父角色id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /authority/createAuthority [post]
func (a *Authority) CreateAuthority(c *fiber.Ctx) error {
	var authority system.Authority
	_ = c.BodyParser(&authority)
	if err := utils.Verify(authority, utils.AuthorityVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err, authBack := authorityService.CreateAuthority(authority); err != nil {
		global.LOG.Error("创建失败!", zap.Any("err", err))
		return response.FailWithMessage("创建失败"+err.Error(), c)
	} else {
		_ = menuService.AddMenuAuthority(systemReq.DefaultMenu(), authority.AuthorityId)
		_ = casbinService.UpdateCasbin(authority.AuthorityId, systemReq.DefaultCasbin())
		return response.OkWithDetailed(systemRes.SysAuthorityResponse{Authority: authBack}, "创建成功", c)
	}
}

// CopyAuthority
// @Tags Authority
// @Summary 拷贝角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body response.SysAuthorityCopyResponse true "旧角色id, 新权限id, 新权限名, 新父角色id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"拷贝成功"}"
// @Router /authority/copyAuthority [post]
func (a *Authority) CopyAuthority(c *fiber.Ctx) error {
	var copyInfo systemRes.SysAuthorityCopyResponse
	_ = c.BodyParser(&copyInfo)
	if err := utils.Verify(copyInfo, utils.OldAuthorityVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err := utils.Verify(copyInfo.Authority, utils.AuthorityVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err, authBack := authorityService.CopyAuthority(copyInfo); err != nil {
		global.LOG.Error("拷贝失败!", zap.Any("err", err))
		return response.FailWithMessage("拷贝失败"+err.Error(), c)
	} else {
		return response.OkWithDetailed(systemRes.SysAuthorityResponse{Authority: authBack}, "拷贝成功", c)
	}
}

// DeleteAuthority
// @Tags Authority
// @Summary 删除角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.Authority true "删除角色"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /authority/deleteAuthority [post]
func (a *Authority) DeleteAuthority(c *fiber.Ctx) error {
	var authority system.Authority
	_ = c.BodyParser(&authority)
	if err := utils.Verify(authority, utils.AuthorityIdVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err := authorityService.DeleteAuthority(&authority); err != nil { // 删除角色之前需要判断是否有用户正在使用此角色
		global.LOG.Error("删除失败!", zap.Any("err", err))
		return response.FailWithMessage("删除失败"+err.Error(), c)
	} else {
		return response.OkWithMessage("删除成功", c)
	}
}

// UpdateAuthority
// @Tags Authority
// @Summary 更新角色信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.Authority true "权限id, 权限名, 父角色id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /authority/updateAuthority [post]
func (a *Authority) UpdateAuthority(c *fiber.Ctx) error {
	var auth system.Authority
	_ = c.BodyParser(&auth)
	if err := utils.Verify(auth, utils.AuthorityVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err, authority := authorityService.UpdateAuthority(auth); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		return response.FailWithMessage("更新失败"+err.Error(), c)
	} else {
		return response.OkWithDetailed(systemRes.SysAuthorityResponse{Authority: authority}, "更新成功", c)
	}
}

// GetAuthorityList
// @Tags Authority
// @Summary 分页获取角色列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /authority/getAuthorityList [post]
func (a *Authority) GetAuthorityList(c *fiber.Ctx) error {
	var pageInfo request.PageInfo
	_ = c.BodyParser(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err, list, total := authorityService.GetAuthorityInfoList(pageInfo); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		return response.FailWithMessage("获取失败"+err.Error(), c)
	} else {
		return response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// SetDataAuthority
// @Tags Authority
// @Summary 设置角色资源权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.Authority true "设置角色资源权限"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"设置成功"}"
// @Router /authority/setDataAuthority [post]
func (a *Authority) SetDataAuthority(c *fiber.Ctx) error {
	var auth system.Authority
	_ = c.BodyParser(&auth)
	if err := utils.Verify(auth, utils.AuthorityIdVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err := authorityService.SetDataAuthority(auth); err != nil {
		global.LOG.Error("设置失败!", zap.Any("err", err))
		return response.FailWithMessage("设置失败"+err.Error(), c)
	} else {
		return response.OkWithMessage("设置成功", c)
	}
}
