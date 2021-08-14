package system

import "FiberBoot/service"

type ApiGroup struct {
	Systems
	Authority
	Base
	Casbin
	Api
	DB
	Jwt
	Operations
	AuthorityMenu
}

var authorityService = service.AppService.SystemService.AuthorityService
var apiService = service.AppService.SystemService.ApiService
var menuService = service.AppService.SystemService.MenuService
var casbinService = service.AppService.SystemService.CasbinService
var emailService = service.AppService.SystemService.EmailService
var initDBService = service.AppService.SystemService.InitDBService
var jwtService = service.AppService.SystemService.JwtService
var baseMenuService = service.AppService.SystemService.BaseMenuService
var operationsService = service.AppService.SystemService.OperationsService
var userService = service.AppService.SystemService.UserService
var configService = service.AppService.SystemService.ConfigService
