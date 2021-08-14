package system

type Menu struct {
	BaseMenu
	MenuId      string              `json:"menuId" gorm:"comment:菜单ID"`
	AuthorityId string              `json:"-" gorm:"comment:角色ID"`
	Children    []Menu              `json:"children" gorm:"-"`
	Parameters  []BaseMenuParameter `json:"parameters" gorm:"foreignKey:BaseMenuID;references:MenuId"`
}

func (s Menu) TableName() string {
	return "authority_menu"
}
