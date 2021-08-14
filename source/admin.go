package source

import (
	"FiberBoot/global"
	"FiberBoot/model/system"
	"github.com/gookit/color"
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

var Admin = new(admin)

type admin struct{}

var admins = []system.User{
	{MODEL: global.MODEL{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(), Username: "admin", Password: "e10adc3949ba59abbe56e057f20f883e", NickName: "超级管理员", Avatar: "", AuthorityId: "888"},
	{MODEL: global.MODEL{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(), Username: "manager", Password: "3ec063004a6f31642261936a379fde3d", NickName: "管理员", Avatar: "", AuthorityId: "9528"},
}

// Init
//@author: Flame
//@description: users 表数据初始化
func (a *admin) Init() error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 2}).Find(&[]system.User{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> users 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&admins).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> users 表初始数据成功!")
		return nil
	})
}
