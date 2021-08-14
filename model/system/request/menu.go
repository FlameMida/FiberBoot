package request

import (
	"FiberBoot/global"
	"FiberBoot/model/system"
)

// AddMenuAuthorityInfo Add menu authority info structure
type AddMenuAuthorityInfo struct {
	Menus       []system.BaseMenu
	AuthorityId string // 角色ID
}

func DefaultMenu() []system.BaseMenu {
	return []system.BaseMenu{{
		MODEL:     global.MODEL{ID: 1},
		ParentId:  "0",
		Path:      "dashboard",
		Name:      "dashboard",
		Component: "view/dashboard/index.vue",
		Sort:      1,
		Meta: system.Meta{
			Title: "仪表盘",
			Icon:  "setting",
		},
	}}
}
