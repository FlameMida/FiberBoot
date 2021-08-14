package system

type UserAuthority struct {
	UserId               uint   `gorm:"column:user_id"`
	AuthorityAuthorityId string `gorm:"column:authority_authority_id"`
}

func (s *UserAuthority) TableName() string {
	return "user_authority"
}
