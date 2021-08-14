package router

import (
	"FiberBoot/router/example"
	"FiberBoot/router/system"
)

type Router struct {
	System  system.RouterGroup
	Example example.RouterGroup
}

var AppRouter = new(Router)
