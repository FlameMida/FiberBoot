package system

import (
	"FiberBoot/global"
	"FiberBoot/model/common/response"
	"FiberBoot/model/system"
	systemRes "FiberBoot/model/system/response"
	"FiberBoot/utils"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Systems struct {
}

// GetSystemConfig
// @Tags System
// @Summary 获取配置文件内容
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /system/getSystemConfig [post]
func (s *Systems) GetSystemConfig(c *fiber.Ctx) error {
	if err, config := configService.GetSystemConfig(); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		return response.FailWithMessage("获取失败", c)
	} else {
		return response.OkWithDetailed(systemRes.SysConfigResponse{Config: config}, "获取成功", c)
	}
}

// SetSystemConfig
// @Tags System
// @Summary 设置配置文件内容
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body system.System true "设置配置文件内容"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"设置成功"}"
// @Router /system/setSystemConfig [post]
func (s *Systems) SetSystemConfig(c *fiber.Ctx) error {
	var sys system.System
	_ = c.BodyParser(&sys)
	if err := configService.SetSystemConfig(sys); err != nil {
		global.LOG.Error("设置失败!", zap.Any("err", err))
		return response.FailWithMessage("设置失败", c)
	} else {
		return response.OkWithData("设置成功", c)
	}
}

// ReloadSystem
// @Tags System
// @Summary 重启系统
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"code":0,"data":{},"msg":"重启系统成功"}"
// @Router /system/reloadSystem [post]
func (s *Systems) ReloadSystem(c *fiber.Ctx) error {
	err := utils.Reload()
	if err != nil {
		global.LOG.Error("重启系统失败!", zap.Any("err", err))
		return response.FailWithMessage("重启系统失败", c)
	} else {
		return response.OkWithMessage("重启系统成功", c)
	}
}

// GetServerInfo
// @Tags System
// @Summary 获取服务器信息
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /system/getServerInfo [post]
func (s *Systems) GetServerInfo(c *fiber.Ctx) error {
	if server, err := configService.GetServerInfo(); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		return response.FailWithMessage("获取失败", c)
	} else {
		return response.OkWithDetailed(fiber.Map{"server": server}, "获取成功", c)
	}
}
