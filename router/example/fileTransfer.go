package example

import (
	"FiberBoot/api/v1"
	"github.com/gofiber/fiber/v2"
)

type FileTransferRouter struct {
}

func (e *FileTransferRouter) InitFileTransferRouter(Router fiber.Router) {
	fileTransferRouter := Router.Group("fileTransfer")
	var exaFileTransferApi = v1.AppApi.ExampleApi.FileTransferApi
	{
		fileTransferRouter.Post("/upload", exaFileTransferApi.UploadFile)                                 // 上传文件
		fileTransferRouter.Post("/getFileList", exaFileTransferApi.GetFileList)                           // 获取上传文件列表
		fileTransferRouter.Post("/deleteFile", exaFileTransferApi.DeleteFile)                             // 删除指定文件
		fileTransferRouter.Post("/breakpointContinue", exaFileTransferApi.BreakpointContinue)             // 断点续传
		fileTransferRouter.Get("/findFile", exaFileTransferApi.FindFile)                                  // 查询当前文件成功的切片
		fileTransferRouter.Post("/breakpointContinueFinish", exaFileTransferApi.BreakpointContinueFinish) // 查询当前文件成功的切片
		fileTransferRouter.Post("/removeChunk", exaFileTransferApi.RemoveChunk)                           // 查询当前文件成功的切片
	}
}
