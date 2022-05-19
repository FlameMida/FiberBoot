package middleware

import (
	"FiberBoot/global"
	"FiberBoot/model/system"
	"FiberBoot/model/system/request"
	"FiberBoot/service"
	"FiberBoot/utils"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"strconv"
	"sync"
	"time"
)

var userService = service.AppService.SystemService.UserService

func ErrorToEmail() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var username string
		if claims := c.Locals("claims"); claims != nil {
			waitUse := claims.(*request.CustomClaims)
			username = waitUse.Username
		} else {
			id, _ := strconv.Atoi(c.Get("x-user-id"))
			err, user := userService.FindUserById(id)
			if err != nil {
				username = "Unknown"
			}
			username = user.Username
		}
		body := c.Request().Body()
		record := system.Operations{
			Ip:     c.IP(),
			Method: c.Method(),
			Path:   c.OriginalURL(),
			Agent:  c.Get("User-Agent"),
			Body:   string(body),
		}
		now := time.Now()

		var (
			once       sync.Once
			errHandler fiber.ErrorHandler
			errPadding = 15
		)

		// Set error handler once
		once.Do(func() {
			stack := c.App().Stack()
			for m := range stack {
				for r := range stack[m] {
					if len(stack[m][r].Path) > errPadding {
						errPadding = len(stack[m][r].Path)

					}
				}
			}
			// override error handler
			errHandler = c.App().Config().ErrorHandler
		})

		chainErr := c.Next()
		if chainErr != nil {
			if err := errHandler(c, chainErr); err != nil {
				_ = c.SendStatus(fiber.StatusInternalServerError)
			}
			record.ErrorMessage = chainErr.Error()
		}

		latency := time.Now().Sub(now)
		status := c.Response().StatusCode()
		str := "接收到的请求为" + record.Body + "\n" + "请求方式为" + record.Method + "\n" + "报错信息如下" + record.ErrorMessage + "\n" + "耗时" + latency.String() + "\n"
		if status != 200 {
			subject := username + "" + record.Ip + "调用了" + record.Path + "报错了"
			if err := utils.ErrorToEmail(subject, str); err != nil {
				global.LOG.Error("ErrorToEmail Failed, err:", zap.Any("err", err))
			}
		}
		return nil
	}
}
