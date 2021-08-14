package system

import (
	"FiberBoot/api/v1"
	"FiberBoot/middleware"
	"github.com/gofiber/fiber/v2"
)

type MenuRouter struct {
}

func (s *MenuRouter) InitMenuRouter(Router fiber.Router) (R fiber.Router) {
	menuRouter := Router.Group("menu").Use(middleware.Operations())
	var authorityMenuApi = v1.AppApi.SystemApi.AuthorityMenu
	{
		menuRouter.Post("getMenu", authorityMenuApi.GetMenu)                   // 获取菜单树
		menuRouter.Post("getMenuList", authorityMenuApi.GetMenuList)           // 分页获取基础menu列表
		menuRouter.Post("addBaseMenu", authorityMenuApi.AddBaseMenu)           // 新增菜单
		menuRouter.Post("getBaseMenuTree", authorityMenuApi.GetBaseMenuTree)   // 获取用户动态路由
		menuRouter.Post("addMenuAuthority", authorityMenuApi.AddMenuAuthority) //	增加menu和角色关联关系
		menuRouter.Post("getMenuAuthority", authorityMenuApi.GetMenuAuthority) // 获取指定角色menu
		menuRouter.Post("deleteBaseMenu", authorityMenuApi.DeleteBaseMenu)     // 删除菜单
		menuRouter.Post("updateBaseMenu", authorityMenuApi.UpdateBaseMenu)     // 更新菜单
		menuRouter.Post("getBaseMenuById", authorityMenuApi.GetBaseMenuById)   // 根据id获取菜单
	}
	return menuRouter
}
