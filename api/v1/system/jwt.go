package system

import (
	"FiberBoot/global"
	"FiberBoot/model/common/response"
	"FiberBoot/model/system"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Jwt struct {
}

// JsonInBlacklist
// @Tags Jwt
// @Summary jwt加入黑名单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"拉黑成功"}"
// @Router /jwt/jsonInBlacklist [post]
func (j *Jwt) JsonInBlacklist(c *fiber.Ctx) error {
	token := c.Get("x-token")
	jwt := system.JwtBlacklist{Jwt: token}
	if err := jwtService.JsonInBlacklist(jwt); err != nil {
		global.LOG.Error("jwt作废失败!", zap.Any("err", err))
		return response.FailWithMessage("jwt作废失败", c)
	} else {
		return response.OkWithMessage("jwt作废成功", c)
	}
}
