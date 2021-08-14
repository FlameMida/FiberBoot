package source

import (
	"FiberBoot/global"
	"FiberBoot/model/system"
	"github.com/gookit/color"
	"gorm.io/gorm"
)

var UserAuthority = new(userAuthority)

type userAuthority struct{}

var userAuthorityModel = []system.UserAuthority{
	{1, "888"},
	{1, "8881"},
	{1, "9528"},
	{2, "888"},
}

// Init @description: user_authority 数据初始化
func (a *userAuthority) Init() error {
	return global.DB.Model(&system.UserAuthority{}).Transaction(func(tx *gorm.DB) error {
		if tx.Where("user_id IN (1, 2)").Find(&[]system.UserAuthority{}).RowsAffected == 4 {
			color.Danger.Println("\n[Mysql] --> user_authority 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&userAuthorityModel).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> user_authority 表初始数据成功!")
		return nil
	})
}
