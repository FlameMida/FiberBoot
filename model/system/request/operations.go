package request

import (
	"FiberBoot/model/common/request"
	"FiberBoot/model/system"
)

type OperationsSearch struct {
	system.Operations
	request.PageInfo
}
