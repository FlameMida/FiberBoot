package system

import (
	"FiberBoot/global"
	"FiberBoot/model/common/response"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// EmailTest
//
// @Tags System
// @Summary 发送测试邮件
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /email/emailTest [post]
func (s *Systems) EmailTest(c *fiber.Ctx) error {
	if err := emailService.EmailTest(); err != nil {
		global.LOG.Error("发送失败!", zap.Any("err", err))
		return response.FailWithMessage("发送失败", c)
	} else {
		return response.OkWithData("发送成功", c)
	}
}
