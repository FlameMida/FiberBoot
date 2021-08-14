package example

import (
	"FiberBoot/global"
	"FiberBoot/model/system"
)

type Customer struct {
	global.MODEL
	CustomerName      string      `json:"customerName" form:"customerName" gorm:"comment:客户名"`             // 客户名
	CustomerPhoneData string      `json:"customerPhoneData" form:"customerPhoneData" gorm:"comment:客户手机号"` // 客户手机号
	UserID            uint        `json:"userId" form:"userId" gorm:"comment:管理ID"`                        // 管理ID
	UserAuthorityID   string      `json:"userAuthorityID" form:"userAuthorityID" gorm:"comment:管理角色ID"`    // 管理角色ID
	User              system.User `json:"user" form:"user" gorm:"comment:管理详情"`                            // 管理详情
}
