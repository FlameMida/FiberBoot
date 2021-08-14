package system

import (
	"FiberBoot/global"
	"FiberBoot/model/common/request"
	"FiberBoot/model/system"
	"FiberBoot/utils"
	"errors"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

//@author: Flame
//@function: Register
//@description: 用户注册
//@param: u model.User
//@return: err error, userInter model.User

type UserService struct{}

func (userService *UserService) Register(u system.User) (err error, userInter system.User) {
	var user system.User
	if !errors.Is(global.DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("用户名已注册"), userInter
	}
	// 否则 附加uuid 密码md5简单加密 注册
	u.Password = utils.MD5V([]byte(u.Password))
	u.UUID = uuid.NewV4()
	err = global.DB.Create(&u).Error
	return err, u
}

//@author: Flame
//@function: Login
//@description: 用户登录
//@param: u *model.User
//@return: err error, userInter *model.User

func (userService *UserService) Login(u *system.User) (err error, userInter *system.User) {
	var user system.User
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.DB.Where("username = ? AND password = ?", u.Username, u.Password).Preload("Authorities").Preload("Authority").First(&user).Error
	return err, &user
}

//@author: Flame
//@function: ChangePassword
//@description: 修改用户密码
//@param: u *model.User, newPassword string
//@return: err error, userInter *model.User

func (userService *UserService) ChangePassword(u *system.User, newPassword string) (err error, userInter *system.User) {
	var user system.User
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Update("password", utils.MD5V([]byte(newPassword))).Error
	return err, u
}

//@author: Flame
//@function: GetUserInfoList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: err error, list interface{}, total int64

func (userService *UserService) GetUserInfoList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.DB.Model(&system.User{})
	var userList []system.User
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Preload("Authorities").Preload("Authority").Find(&userList).Error
	return err, userList, total
}

//@author: Flame
//@function: SetUserAuthority
//@description: 设置一个用户的权限
//@param: uuid uuid.UUID, authorityId string
//@return: err error

func (userService *UserService) SetUserAuthority(id uint, uuid uuid.UUID, authorityId string) (err error) {
	assignErr := global.DB.Where("user_id = ? AND authority_authority_id = ?", id, authorityId).First(&system.UserAuthority{}).Error
	if errors.Is(assignErr, gorm.ErrRecordNotFound) {
		return errors.New("该用户无此角色")
	}
	err = global.DB.Where("uuid = ?", uuid).First(&system.User{}).Update("authority_id", authorityId).Error
	return err
}

//@author: Flame
//@function: SetUserAuthorities
//@description: 设置一个用户的权限
//@param: id uint, authorityIds []string
//@return: err error

func (userService *UserService) SetUserAuthorities(id uint, authorityIds []string) (err error) {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		TxErr := tx.Delete(&[]system.UserAuthority{}, "user_id = ?", id).Error
		if TxErr != nil {
			return TxErr
		}
		var useAuthority []system.UserAuthority
		for _, v := range authorityIds {
			useAuthority = append(useAuthority, system.UserAuthority{
				UserId: id, AuthorityAuthorityId: v,
			})
		}
		TxErr = tx.Create(&useAuthority).Error
		if TxErr != nil {
			return TxErr
		}
		// 返回 nil 提交事务
		return nil
	})
}

//@author: Flame
//@function: DeleteUser
//@description: 删除用户
//@param: id float64
//@return: err error

func (userService *UserService) DeleteUser(id float64) (err error) {
	var user system.User
	err = global.DB.Where("id = ?", id).Delete(&user).Error
	err = global.DB.Delete(&[]system.UserAuthority{}, "user_id = ?", id).Error
	return err
}

//@author: Flame
//@function: SetUserInfo
//@description: 设置用户信息
//@param: reqUser model.User
//@return: err error, user model.User

func (userService *UserService) SetUserInfo(reqUser system.User) (err error, user system.User) {
	err = global.DB.Updates(&reqUser).Error
	return err, reqUser
}

//@author: Flame
//@function: GetUserInfo
//@description: 获取用户信息
//@param: uuid uuid.UUID
//@return: err error, user system.User

func (userService *UserService) GetUserInfo(uuid uuid.UUID) (err error, user system.User) {
	var reqUser system.User
	err = global.DB.Preload("Authorities").Preload("Authority").First(&reqUser, "uuid = ?", uuid).Error
	return err, reqUser
}

//@author: Flame
//@function: FindUserById
//@description: 通过id获取用户信息
//@param: id int
//@return: err error, user *model.User

func (userService *UserService) FindUserById(id int) (err error, user *system.User) {
	var u system.User
	err = global.DB.Where("`id` = ?", id).First(&u).Error
	return err, &u
}

//@author: Flame
//@function: FindUserByUuid
//@description: 通过uuid获取用户信息
//@param: uuid string
//@return: err error, user *model.User

func (userService *UserService) FindUserByUuid(uuid string) (err error, user *system.User) {
	var u system.User
	if err = global.DB.Where("`uuid` = ?", uuid).First(&u).Error; err != nil {
		return errors.New("用户不存在"), &u
	}
	return nil, &u
}
