package system

type RouterGroup struct {
	ApiRouter
	AuthorityRouter
	BaseRouter
	CasbinRouter
	EmailRouter
	InitRouter
	JwtRouter
	MenuRouter
	OperationsRouter
	SysRouter
	UserRouter
}
