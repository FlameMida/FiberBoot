package service

import (
	"FiberBoot/service/example"
	"FiberBoot/service/system"
)

type Service struct {
	ExampleService example.ServiceGroup
	SystemService  system.ServiceGroup
}

var AppService = new(Service)
