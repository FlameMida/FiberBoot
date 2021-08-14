package response

import "FiberBoot/model/example"

type FilePathResponse struct {
	FilePath string `json:"filePath"`
}

type FileResponse struct {
	File example.File `json:"file"`
}
