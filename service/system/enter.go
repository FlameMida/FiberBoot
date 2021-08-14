package system

type ServiceGroup struct {
	JwtService
	ApiService
	AuthorityService
	BaseMenuService
	CasbinService
	EmailService
	InitDBService
	MenuService
	OperationsService
	ConfigService
	UserService
}
