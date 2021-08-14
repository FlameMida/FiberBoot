package example

import "FiberBoot/model/system"

type ExcelInfo struct {
	FileName string            `json:"fileName"` // 文件名
	InfoList []system.BaseMenu `json:"infoList"`
}
