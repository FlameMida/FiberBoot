package initialize

import (
	_ "FiberBoot/docs"
	"FiberBoot/global"
	"FiberBoot/middleware"
	"FiberBoot/router"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

// 初始化总路由

func Routers() *fiber.App {
	var Router = fiber.New()
	Router.Static(global.CONFIG.Local.Path, global.CONFIG.Local.Path) // 为用户头像和文件提供静态地址
	global.LOG.Info("use middleware logger")
	Router.Use(middleware.Logger())

	global.LOG.Info("use middleware recover")
	Router.Use(middleware.Recover())
	// 跨域
	global.LOG.Info("use middleware cors")
	Router.Use(middleware.Cors())

	global.LOG.Info("register swagger handler")
	Router.Get("/swagger/*", swagger.HandlerDefault)
	// 方便统一添加路由组前缀 多服务器上线使用

	//获取路由组实例
	systemRouter := router.AppRouter.System
	exampleRouter := router.AppRouter.Example
	PublicGroup := Router.Group("")
	{
		systemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
		systemRouter.InitInitRouter(PublicGroup) // 自动初始化相关
	}
	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		systemRouter.InitApiRouter(PrivateGroup)             // 注册功能api路由
		systemRouter.InitJwtRouter(PrivateGroup)             // jwt相关路由
		systemRouter.InitUserRouter(PrivateGroup)            // 注册用户路由
		systemRouter.InitMenuRouter(PrivateGroup)            // 注册menu路由
		systemRouter.InitEmailRouter(PrivateGroup)           // 邮件相关路由
		systemRouter.InitSystemRouter(PrivateGroup)          // system相关路由
		systemRouter.InitCasbinRouter(PrivateGroup)          // 权限相关路由
		systemRouter.InitAuthorityRouter(PrivateGroup)       // 注册角色路由
		systemRouter.InitOperationsRouter(PrivateGroup)      // 操作记录
		exampleRouter.InitFileTransferRouter(PrivateGroup)   // 文件上传下载功能路由
		exampleRouter.InitExcelRouter(PrivateGroup)          // 表格导入导出
		exampleRouter.InitSimpleUploaderRouter(PrivateGroup) // 断点续传（插件版）
		exampleRouter.InitCustomerRouter(PrivateGroup)       // 客户路由

	}
	global.LOG.Info("router register success")
	return Router
}
