package response

import "FiberBoot/model/example"

type ExaFileResponse struct {
	File example.FileTransfer `json:"file"`
}
