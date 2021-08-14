package core

import (
	"FiberBoot/global"
	"FiberBoot/initialize"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"time"
)

type Server interface {
	ServeAsync(string, *fiber.App) error
}

func RunServer() {
	if global.CONFIG.System.UseMultipoint {
		// 初始化redis服务
		initialize.Redis()
	}
	Router := initialize.Routers()
	address := fmt.Sprintf(":%d", global.CONFIG.System.Addr)

	time.Sleep(10 * time.Microsecond)

	global.LOG.Info("server run success on ", zap.String("address", address))
	fmt.Printf("[GIN-STARTER]文档地址:http://127.0.0.1%s/swagger/index.html \n", address)

	global.LOG.Error(newServer().ServeAsync(address, Router).Error())
}

//@author: Flame
//@function: NewOss
//@description: OSS接口
//@description: OSS的实例化方法
//@return: OSS

func newServer() Server {
	return &ServerImpl{}
}
