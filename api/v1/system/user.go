package system

import (
	"FiberBoot/global"
	"FiberBoot/middleware"
	"FiberBoot/model/common/request"
	"FiberBoot/model/common/response"
	"FiberBoot/model/system"
	systemReq "FiberBoot/model/system/request"
	systemRes "FiberBoot/model/system/response"
	"FiberBoot/utils"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
)

// Login
// @Tags Base
// @Summary 用户登录
// @Produce  application/json
// @Param data body systemReq.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func (b *Base) Login(c *fiber.Ctx) error {
	var l systemReq.Login
	_ = c.BodyParser(&l)
	if err := utils.Verify(l, utils.LoginVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if store.Verify(l.CaptchaId, l.Captcha, true) {
		u := &system.User{Username: l.Username, Password: l.Password}
		if err, user := userService.Login(u); err != nil {
			global.LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Any("err", err))
			return response.FailWithMessage("用户名不存在或者密码错误", c)
		} else {
			return b.tokenNext(c, *user)
		}
	} else {
		return response.FailWithMessage("验证码错误", c)
	}
}

// 登录以后签发jwt
func (b *Base) tokenNext(c *fiber.Ctx, user system.User) error {
	j := &middleware.JWT{SigningKey: []byte(global.CONFIG.JWT.SigningKey)} // 唯一签名
	claims := systemReq.CustomClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		NickName:    user.NickName,
		Username:    user.Username,
		AuthorityId: user.AuthorityId,
		BufferTime:  global.CONFIG.JWT.BufferTime, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                          // 签名生效时间
			ExpiresAt: time.Now().Unix() + global.CONFIG.JWT.ExpiresTime, // 过期时间 7天  配置文件
			Issuer:    "yourKeys",                                        // 签名的发行者
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		global.LOG.Error("获取token失败!", zap.Any("err", err))
		return response.FailWithMessage("获取token失败", c)

	}
	if !global.CONFIG.System.UseMultipoint {
		return response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)

	}
	if err, jwtStr := jwtService.GetRedisJWT(user.Username); err == redis.Nil {
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			global.LOG.Error("设置登录状态失败!", zap.Any("err", err))
			return response.FailWithMessage("设置登录状态失败", c)

		}
		return response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	} else if err != nil {
		global.LOG.Error("设置登录状态失败!", zap.Any("err", err))
		return response.FailWithMessage("设置登录状态失败", c)
	} else {
		var blackJWT system.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := jwtService.JsonInBlacklist(blackJWT); err != nil {
			return response.FailWithMessage("jwt作废失败", c)

		}
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			return response.FailWithMessage("设置登录状态失败", c)

		}
		return response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	}
}

// Register
// @Tags User
// @Summary 用户注册账号
// @Produce  application/json
// @Param data body systemReq.Register true "用户名, 昵称, 密码, 角色ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /user/register [post]
func (b *Base) Register(c *fiber.Ctx) error {
	var r systemReq.Register
	_ = c.BodyParser(&r)
	if err := utils.Verify(r, utils.RegisterVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	var authorities []system.Authority
	for _, v := range r.AuthorityIds {
		authorities = append(authorities, system.Authority{
			AuthorityId: v,
		})
	}
	user := &system.User{Username: r.Username, NickName: r.NickName, Password: r.Password, Avatar: r.Avatar, AuthorityId: r.AuthorityId, Authorities: authorities}
	err, userReturn := userService.Register(*user)
	if err != nil {
		global.LOG.Error("注册失败!", zap.Any("err", err))
		return response.FailWithDetailed(systemRes.UserResponse{User: userReturn}, "注册失败", c)
	} else {
		return response.OkWithDetailed(systemRes.UserResponse{User: userReturn}, "注册成功", c)
	}
}

// ChangePassword
//
// @Tags User
// @Summary 用户修改密码
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body systemReq.ChangePasswordStruct true "用户名, 原密码, 新密码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/changePassword [put]
func (b *Base) ChangePassword(c *fiber.Ctx) error {
	var user systemReq.ChangePasswordStruct
	_ = c.BodyParser(&user)
	if err := utils.Verify(user, utils.ChangePasswordVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)

	}
	u := &system.User{Username: user.Username, Password: user.Password}
	if err, _ := userService.ChangePassword(u, user.NewPassword); err != nil {
		global.LOG.Error("修改失败!", zap.Any("err", err))
		return response.FailWithMessage("修改失败，原密码与当前账户不符", c)
	} else {
		return response.OkWithMessage("修改成功", c)
	}
}

// GetUserList
// @Tags User
// @Summary 分页获取用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/getUserList [post]
func (b *Base) GetUserList(c *fiber.Ctx) error {
	var pageInfo request.PageInfo
	_ = c.BodyParser(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)

	}
	if err, list, total := userService.GetUserInfoList(pageInfo); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		return response.FailWithMessage("获取失败", c)
	} else {
		return response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// SetUserAuthority
// @Tags User
// @Summary 更改用户权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body systemReq.SetUserAuth true "用户UUID, 角色ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/setUserAuthority [post]
func (b *Base) SetUserAuthority(c *fiber.Ctx) error {
	var sua systemReq.SetUserAuth
	_ = c.BodyParser(&sua)
	if UserVerifyErr := utils.Verify(sua, utils.SetUserAuthorityVerify); UserVerifyErr != nil {
		return response.FailWithMessage(UserVerifyErr.Error(), c)
	}
	userID := utils.GetUserID(c)
	uuid := utils.GetUserUuid(c)
	if err := userService.SetUserAuthority(userID, uuid, sua.AuthorityId); err != nil {
		global.LOG.Error("修改失败!", zap.Any("err", err))
		return response.FailWithMessage(err.Error(), c)
	} else {
		claims := utils.GetUserInfo(c)
		j := &middleware.JWT{SigningKey: []byte(global.CONFIG.JWT.SigningKey)} // 唯一签名
		claims.AuthorityId = sua.AuthorityId
		if token, err := j.CreateToken(*claims); err != nil {
			global.LOG.Error("修改失败!", zap.Any("err", err))
			return response.FailWithMessage(err.Error(), c)
		} else {
			c.Set("new-token", token)
			c.Set("new-expires-at", strconv.FormatInt(claims.ExpiresAt, 10))
			return response.OkWithMessage("修改成功", c)
		}

	}
}

// SetUserAuthorities
// @Tags User
// @Summary 设置用户权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body systemReq.SetUserAuthorities true "用户UUID, 角色ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/setUserAuthorities [post]
func (b *Base) SetUserAuthorities(c *fiber.Ctx) error {
	var sua systemReq.SetUserAuthorities
	_ = c.BodyParser(&sua)
	if err := userService.SetUserAuthorities(sua.ID, sua.AuthorityIds); err != nil {
		global.LOG.Error("修改失败!", zap.Any("err", err))
		return response.FailWithMessage("修改失败", c)
	} else {
		return response.OkWithMessage("修改成功", c)
	}
}

// DeleteUser
// @Tags User
// @Summary 删除用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "用户ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /user/deleteUser [delete]
func (b *Base) DeleteUser(c *fiber.Ctx) error {
	var reqId request.GetById
	_ = c.BodyParser(&reqId)
	if err := utils.Verify(reqId, utils.IdVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)

	}
	jwtId := utils.GetUserID(c)
	if jwtId == uint(reqId.ID) {
		return response.FailWithMessage("删除失败, 自杀失败", c)
	}
	if err := userService.DeleteUser(reqId.ID); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		return response.FailWithMessage("删除失败", c)
	} else {
		return response.OkWithMessage("删除成功", c)
	}
}

// SetUserInfo
// @Tags User
// @Summary 设置用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.User true "ID, 用户名, 昵称, 头像链接"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"设置成功"}"
// @Router /user/setUserInfo [put]
func (b *Base) SetUserInfo(c *fiber.Ctx) error {
	var user system.User
	_ = c.BodyParser(&user)
	if err := utils.Verify(user, utils.IdVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if err, ReqUser := userService.SetUserInfo(user); err != nil {
		global.LOG.Error("设置失败!", zap.Any("err", err))
		return response.FailWithMessage("设置失败", c)
	} else {
		return response.OkWithDetailed(fiber.Map{"userInfo": ReqUser}, "设置成功", c)
	}
}

// GetUserInfo
// @Tags User
// @Summary 获取用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/getUserInfo [get]
func (b *Base) GetUserInfo(c *fiber.Ctx) error {
	uuid := utils.GetUserUuid(c)
	if err, ReqUser := userService.GetUserInfo(uuid); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		return response.FailWithMessage("获取失败", c)
	} else {
		return response.OkWithDetailed(fiber.Map{"userInfo": ReqUser}, "获取成功", c)
	}
}
