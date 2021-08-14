package response

import "FiberBoot/model/system"

type SysMenusResponse struct {
	Menus []system.Menu `json:"menus"`
}

type SysBaseMenusResponse struct {
	Menus []system.BaseMenu `json:"menus"`
}

type SysBaseMenuResponse struct {
	Menu system.BaseMenu `json:"menu"`
}
