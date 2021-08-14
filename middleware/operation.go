package middleware

import (
	"FiberBoot/global"
	"FiberBoot/model/system"
	"FiberBoot/model/system/request"
	"FiberBoot/service"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"strconv"
	"sync"
	"time"
)

var OperationsService = service.AppService.SystemService.OperationsService

func Operations() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var body []byte
		var userId int
		if c.Method() != fiber.MethodGet {

			body = c.Request().Body()
		}
		if claims := c.Locals("claims"); claims != nil {
			waitUse := claims.(*request.CustomClaims)
			userId = int(waitUse.ID)
		} else {
			id, err := strconv.Atoi(c.Get("x-user-id"))
			if err != nil {
				userId = 0
			}
			userId = id
		}
		record := system.Operations{
			Ip:     c.IP(),
			Method: c.Method(),
			Path:   c.OriginalURL(),
			Agent:  c.Get("User-Agent"),
			Body:   string(body),
			UserID: userId,
		}
		now := time.Now()
		var (
			once          sync.Once
			errHandler    fiber.ErrorHandler
			errPadding    = 15
			errPaddingStr = strconv.Itoa(errPadding)
		)

		// Set error handler once
		once.Do(func() {
			stack := c.App().Stack()
			for m := range stack {
				for r := range stack[m] {
					if len(stack[m][r].Path) > errPadding {
						errPadding = len(stack[m][r].Path)
						errPaddingStr = strconv.Itoa(errPadding)
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
		record.Status = c.Response().StatusCode()
		record.Latency = latency
		record.Resp = string(c.Response().Body())
		if err := OperationsService.CreateOperations(record); err != nil {
			global.LOG.Error("create operation record error:", zap.Any("err", err))
		}
		return nil
	}
}
