package system

import (
	"FiberBoot/config"
	"FiberBoot/global"
	"FiberBoot/model/system"
	"FiberBoot/utils"
	"go.uber.org/zap"
)

//@author: Flame
//@function: GetSystemConfig
//@description: 读取配置文件
//@return: err error, conf config.Server

type ConfigService struct {
}

func (configService *ConfigService) GetSystemConfig() (err error, conf config.Server) {
	return nil, global.CONFIG
}

// @description   set system config,
//@author: Flame
//@function: SetSystemConfig
//@description: 设置配置文件
//@param: system model.System
//@return: err error

func (configService *ConfigService) SetSystemConfig(system system.System) (err error) {
	cs := utils.StructToMap(system.Config)
	for k, v := range cs {
		global.VP.Set(k, v)
	}
	err = global.VP.WriteConfig()
	return err
}

//@author: Flame
//@function: GetServerInfo
//@description: 获取服务器信息
//@return: server *utils.Server, err error

func (configService *ConfigService) GetServerInfo() (server *utils.Server, err error) {
	var s utils.Server
	s.Os = utils.InitOS()
	if s.Cpu, err = utils.InitCPU(); err != nil {
		global.LOG.Error("func utils.InitCPU() Failed", zap.String("err", err.Error()))
		return &s, err
	}
	if s.Rrm, err = utils.InitRAM(); err != nil {
		global.LOG.Error("func utils.InitRAM() Failed", zap.String("err", err.Error()))
		return &s, err
	}
	if s.Disk, err = utils.InitDisk(); err != nil {
		global.LOG.Error("func utils.InitDisk() Failed", zap.String("err", err.Error()))
		return &s, err
	}

	return &s, nil
}
