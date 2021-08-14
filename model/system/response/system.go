package response

import "FiberBoot/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
