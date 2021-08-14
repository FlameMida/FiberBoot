package v1

import (
	"FiberBoot/api/v1/example"
	"FiberBoot/api/v1/system"
)

type ApiGroup struct {
	ExampleApi example.ApiGroup
	SystemApi  system.ApiGroup
}

var AppApi = new(ApiGroup)
