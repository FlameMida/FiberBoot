package system

import (
	"FiberBoot/global"
	"FiberBoot/model/common/response"
	"FiberBoot/model/system/request"
	"go.uber.org/zap"

	"github.com/gofiber/fiber/v2"
)

type DB struct {
}

// InitDB
// @Tags InitDB
// @Summary 初始化用户数据库
// @Produce  application/json
// @Param data body request.InitDB true "初始化数据库参数"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"自动创建数据库成功"}"
// @Router /init/initDB [post]
func (i *DB) InitDB(c *fiber.Ctx) error {
	if global.DB != nil {
		global.LOG.Error("已存在数据库配置!")
		return response.FailWithMessage("已存在数据库配置", c)
	}
	var dbInfo request.InitDB
	if err := c.BodyParser(&dbInfo); err != nil {
		global.LOG.Error("参数校验不通过!", zap.Any("err", err))
		return response.FailWithMessage("参数校验不通过", c)
	}
	if err := initDBService.InitDB(dbInfo); err != nil {
		global.LOG.Error("自动创建数据库失败!", zap.Any("err", err))
		return response.FailWithMessage("自动创建数据库失败，请查看后台日志，检查后在进行初始化", c)
	}
	return response.OkWithData("自动创建数据库成功", c)
}

// CheckDB
// @Tags CheckDB
// @Summary 初始化用户数据库
// @Produce  application/json
// @Success 200 {string} string "{"code":0,"data":{},"msg":"探测完成"}"
// @Router /init/checkDB [post]
func (i *DB) CheckDB(c *fiber.Ctx) error {
	if global.DB != nil {
		global.LOG.Info("数据库无需初始化")
		return response.OkWithDetailed(fiber.Map{"needInit": false}, "数据库无需初始化", c)
	} else {
		global.LOG.Info("前往初始化数据库")
		return response.OkWithDetailed(fiber.Map{"needInit": true}, "前往初始化数据库", c)
	}
}
