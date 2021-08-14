package system

import (
	"FiberBoot/global"
	"FiberBoot/model/common/request"
	"FiberBoot/model/common/response"
	"FiberBoot/model/system"
	systemReq "FiberBoot/model/system/request"
	systemRes "FiberBoot/model/system/response"
	"FiberBoot/utils"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type AuthorityMenu struct {
}

// GetMenu
// @Tags AuthorityMenu
// @Summary 获取用户动态路由
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.Empty true "空"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/getMenu [post]
func (a *AuthorityMenu) GetMenu(c *fiber.Ctx) error {
	if err, menus := menuService.GetMenuTree(utils.GetUserAuthorityId(c)); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		return response.FailWithMessage("获取失败", c)
	} else {
		if menus == nil {
			menus = []system.Menu{}
		}
		return response.OkWithDetailed(systemRes.SysMenusResponse{Menus: menus}, "获取成功", c)
	}
}

// GetBaseMenuTree
// @Tags AuthorityMenu
// @Summary 获取用户动态路由
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.Empty true "空"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/getBaseMenuTree [post]
func (a *AuthorityMenu) GetBaseMenuTree(c *fiber.Ctx) error {
	if err, menus := menuService.GetBaseMenuTree(); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		return response.FailWithMessage("获取失败", c)
	} else {
		return response.OkWithDetailed(systemRes.SysBaseMenusResponse{Menus: menus}, "获取成功", c)
	}
}

// AddMenuAuthority
// @Tags AuthorityMenu
// @Summary 增加menu和角色关联关系
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body systemReq.AddMenuAuthorityInfo true "角色ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"添加成功"}"
// @Router /menu/addMenuAuthority [post]
func (a *AuthorityMenu) AddMenuAuthority(c *fiber.Ctx) error {
	var authorityMenu systemReq.AddMenuAuthorityInfo
	_ = c.BodyParser(&authorityMenu)
	if err := utils.Verify(authorityMenu, utils.AuthorityIdVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err := menuService.AddMenuAuthority(authorityMenu.Menus, authorityMenu.AuthorityId); err != nil {
		global.LOG.Error("添加失败!", zap.Any("err", err))
		return response.FailWithMessage("添加失败", c)
	} else {
		return response.OkWithMessage("添加成功", c)
	}
}

// GetMenuAuthority
// @Tags AuthorityMenu
// @Summary 获取指定角色menu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetAuthorityId true "角色ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/GetMenuAuthority [post]
func (a *AuthorityMenu) GetMenuAuthority(c *fiber.Ctx) error {
	var param request.GetAuthorityId
	_ = c.BodyParser(&param)
	if err := utils.Verify(param, utils.AuthorityIdVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err, menus := menuService.GetMenuAuthority(&param); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		return response.FailWithDetailed(systemRes.SysMenusResponse{Menus: menus}, "获取失败", c)
	} else {
		return response.OkWithDetailed(fiber.Map{"menus": menus}, "获取成功", c)
	}
}

// AddBaseMenu
//
// @Tags Menu
// @Summary 新增菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.BaseMenu true "路由path, 父菜单ID, 路由name, 对应前端文件路径, 排序标记"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"添加成功"}"
// @Router /menu/addBaseMenu [post]
func (a *AuthorityMenu) AddBaseMenu(c *fiber.Ctx) error {
	var menu system.BaseMenu
	_ = c.BodyParser(&menu)
	if err := utils.Verify(menu, utils.MenuVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err := utils.Verify(menu.Meta, utils.MenuMetaVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err := menuService.AddBaseMenu(menu); err != nil {
		global.LOG.Error("添加失败!", zap.Any("err", err))

		return response.FailWithMessage("添加失败", c)
	} else {
		return response.OkWithMessage("添加成功", c)
	}
}

// DeleteBaseMenu
//
// @Tags Menu
// @Summary 删除菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "菜单id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /menu/deleteBaseMenu [post]
func (a *AuthorityMenu) DeleteBaseMenu(c *fiber.Ctx) error {
	var menu request.GetById
	_ = c.BodyParser(&menu)
	if err := utils.Verify(menu, utils.IdVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err := baseMenuService.DeleteBaseMenu(menu.ID); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		return response.FailWithMessage("删除失败", c)
	} else {
		return response.OkWithMessage("删除成功", c)
	}
}

// UpdateBaseMenu
// @Tags Menu
// @Summary 更新菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.BaseMenu true "路由path, 父菜单ID, 路由name, 对应前端文件路径, 排序标记"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /menu/updateBaseMenu [post]
func (a *AuthorityMenu) UpdateBaseMenu(c *fiber.Ctx) error {
	var menu system.BaseMenu
	_ = c.BodyParser(&menu)
	if err := utils.Verify(menu, utils.MenuVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err := utils.Verify(menu.Meta, utils.MenuMetaVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err := baseMenuService.UpdateBaseMenu(menu); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		return response.FailWithMessage("更新失败", c)
	} else {
		return response.OkWithMessage("更新成功", c)
	}
}

// GetBaseMenuById
// @Tags Menu
// @Summary 根据id获取菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "菜单id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/getBaseMenuById [post]
func (a *AuthorityMenu) GetBaseMenuById(c *fiber.Ctx) error {
	var idInfo request.GetById
	_ = c.BodyParser(&idInfo)
	if err := utils.Verify(idInfo, utils.IdVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err, menu := baseMenuService.GetBaseMenuById(idInfo.ID); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		return response.FailWithMessage("获取失败", c)
	} else {
		return response.OkWithDetailed(systemRes.SysBaseMenuResponse{Menu: menu}, "获取成功", c)
	}
}

// GetMenuList
// @Tags Menu
// @Summary 分页获取基础menu列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/getMenuList [post]
func (a *AuthorityMenu) GetMenuList(c *fiber.Ctx) error {
	var pageInfo request.PageInfo
	_ = c.BodyParser(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err, menuList, total := menuService.GetInfoList(); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		return response.FailWithMessage("获取失败", c)
	} else {
		return response.OkWithDetailed(response.PageResult{
			List:     menuList,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
