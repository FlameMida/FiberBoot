package system

import (
	"FiberBoot/global"
	"FiberBoot/model/common/response"
	"FiberBoot/model/system/request"
	systemRes "FiberBoot/model/system/response"
	"FiberBoot/utils"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Casbin struct {
}

// UpdateCasbin
//
// @Tags Casbin
// @Summary 更新角色api权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CasbinInReceive true "权限id, 权限模型列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /casbin/UpdateCasbin [post]
func (cas *Casbin) UpdateCasbin(c *fiber.Ctx) error {
	var cmr request.CasbinInReceive
	_ = c.BodyParser(&cmr)
	if err := utils.Verify(cmr, utils.AuthorityIdVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)

	}
	if err := casbinService.UpdateCasbin(cmr.AuthorityId, cmr.CasbinInfos); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		return response.FailWithMessage("更新失败", c)
	} else {
		return response.OkWithMessage("更新成功", c)
	}
}

// GetPolicyPathByAuthorityId
// @Tags Casbin
// @Summary 获取权限列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CasbinInReceive true "权限id, 权限模型列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /casbin/getPolicyPathByAuthorityId [post]
func (cas *Casbin) GetPolicyPathByAuthorityId(c *fiber.Ctx) error {
	var casbin request.CasbinInReceive
	_ = c.BodyParser(&casbin)
	if err := utils.Verify(casbin, utils.AuthorityIdVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	paths := casbinService.GetPolicyPathByAuthorityId(casbin.AuthorityId)
	return response.OkWithDetailed(systemRes.PolicyPathResponse{Paths: paths}, "获取成功", c)
}
