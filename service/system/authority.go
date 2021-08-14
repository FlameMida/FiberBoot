package system

import (
	"FiberBoot/global"
	"FiberBoot/model/common/request"
	"FiberBoot/model/system"
	"FiberBoot/model/system/response"
	"errors"
	"gorm.io/gorm"
	"strconv"
)

//@author: Flame
//@function: CreateAuthority
//@description: 创建一个角色
//@param: auth model.Authority
//@return: err error, authority model.Authority

type AuthorityService struct {
}

var AuthorityServiceApp = new(AuthorityService)

func (authorityService *AuthorityService) CreateAuthority(auth system.Authority) (err error, authority system.Authority) {
	var authorityBox system.Authority
	if !errors.Is(global.DB.Where("authority_id = ?", auth.AuthorityId).First(&authorityBox).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同角色id"), auth
	}
	err = global.DB.Create(&auth).Error
	return err, auth
}

//@author: Flame
//@function: CopyAuthority
//@description: 复制一个角色
//@param: copyInfo response.SysAuthorityCopyResponse
//@return: err error, authority model.Authority

func (authorityService *AuthorityService) CopyAuthority(copyInfo response.SysAuthorityCopyResponse) (err error, authority system.Authority) {
	var authorityBox system.Authority
	if !errors.Is(global.DB.Where("authority_id = ?", copyInfo.Authority.AuthorityId).First(&authorityBox).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同角色id"), authority
	}
	copyInfo.Authority.Children = []system.Authority{}
	err, menus := MenuServiceApp.GetMenuAuthority(&request.GetAuthorityId{AuthorityId: copyInfo.OldAuthorityId})
	var baseMenu []system.BaseMenu
	for _, v := range menus {
		intNum, _ := strconv.Atoi(v.MenuId)
		v.BaseMenu.ID = uint(intNum)
		baseMenu = append(baseMenu, v.BaseMenu)
	}
	copyInfo.Authority.SysBaseMenus = baseMenu
	err = global.DB.Create(&copyInfo.Authority).Error

	paths := CasbinServiceApp.GetPolicyPathByAuthorityId(copyInfo.OldAuthorityId)
	err = CasbinServiceApp.UpdateCasbin(copyInfo.Authority.AuthorityId, paths)
	if err != nil {
		_ = authorityService.DeleteAuthority(&copyInfo.Authority)
	}
	return err, copyInfo.Authority
}

//@author: Flame
//@function: UpdateAuthority
//@description: 更改一个角色
//@param: auth model.Authority
//@return: err error, authority model.Authority

func (authorityService *AuthorityService) UpdateAuthority(auth system.Authority) (err error, authority system.Authority) {
	err = global.DB.Where("authority_id = ?", auth.AuthorityId).First(&system.Authority{}).Updates(&auth).Error
	return err, auth
}

//@author: Flame
//@function: DeleteAuthority
//@description: 删除角色
//@param: auth *model.Authority
//@return: err error

func (authorityService *AuthorityService) DeleteAuthority(auth *system.Authority) (err error) {
	if !errors.Is(global.DB.Where("authority_id = ?", auth.AuthorityId).First(&system.User{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色有用户正在使用禁止删除")
	}
	if !errors.Is(global.DB.Where("parent_id = ?", auth.AuthorityId).First(&system.Authority{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色存在子角色不允许删除")
	}
	db := global.DB.Preload("SysBaseMenus").Where("authority_id = ?", auth.AuthorityId).First(auth)
	err = db.Unscoped().Delete(auth).Error
	if len(auth.SysBaseMenus) > 0 {
		err = global.DB.Model(auth).Association("SysBaseMenus").Delete(auth.SysBaseMenus)
		//err = db.Association("SysBaseMenus").Delete(&auth)
	} else {
		err = db.Error
	}
	err = global.DB.Delete(&[]system.UserAuthority{}, "authority_authority_id = ?", auth.AuthorityId).Error
	CasbinServiceApp.ClearCasbin(0, auth.AuthorityId)
	return err
}

//@author: Flame
//@function: GetAuthorityInfoList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: err error, list interface{}, total int64

func (authorityService *AuthorityService) GetAuthorityInfoList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.DB
	var authority []system.Authority
	err = db.Limit(limit).Offset(offset).Preload("DataAuthorityId").Where("parent_id = 0").Find(&authority).Error
	if len(authority) > 0 {
		for k := range authority {
			err = authorityService.findChildrenAuthority(&authority[k])
		}
	}
	return err, authority, total
}

//@author: Flame
//@function: GetAuthorityInfo
//@description: 获取所有角色信息
//@param: auth model.Authority
//@return: err error, sa model.Authority

func (authorityService *AuthorityService) GetAuthorityInfo(auth system.Authority) (err error, sa system.Authority) {
	err = global.DB.Preload("DataAuthorityId").Where("authority_id = ?", auth.AuthorityId).First(&sa).Error
	return err, sa
}

//@author: Flame
//@function: SetDataAuthority
//@description: 设置角色资源权限
//@param: auth model.Authority
//@return: error

func (authorityService *AuthorityService) SetDataAuthority(auth system.Authority) error {
	var s system.Authority
	global.DB.Preload("DataAuthorityId").First(&s, "authority_id = ?", auth.AuthorityId)
	err := global.DB.Model(&s).Association("DataAuthorityId").Replace(&auth.DataAuthorityId)
	return err
}

//@author: Flame
//@function: SetMenuAuthority
//@description: 菜单与角色绑定
//@param: auth *model.Authority
//@return: error

func (authorityService *AuthorityService) SetMenuAuthority(auth *system.Authority) error {
	var s system.Authority
	global.DB.Preload("SysBaseMenus").First(&s, "authority_id = ?", auth.AuthorityId)
	err := global.DB.Model(&s).Association("SysBaseMenus").Replace(&auth.SysBaseMenus)
	return err
}

//@author: Flame
//@function: findChildrenAuthority
//@description: 查询子角色
//@param: authority *model.Authority
//@return: err error

func (authorityService *AuthorityService) findChildrenAuthority(authority *system.Authority) (err error) {
	err = global.DB.Preload("DataAuthorityId").Where("parent_id = ?", authority.AuthorityId).Find(&authority.Children).Error
	if len(authority.Children) > 0 {
		for k := range authority.Children {
			err = authorityService.findChildrenAuthority(&authority.Children[k])
		}
	}
	return err
}
