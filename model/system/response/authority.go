package response

import "FiberBoot/model/system"

type SysAuthorityResponse struct {
	Authority system.Authority `json:"authority"`
}

type SysAuthorityCopyResponse struct {
	Authority      system.Authority `json:"authority"`
	OldAuthorityId string           `json:"oldAuthorityId"` // 旧角色ID
}
