package example

import (
	"FiberBoot/api/v1"
	"github.com/gofiber/fiber/v2"
)

type SimpleUploaderRouter struct{}

func (e *SimpleUploaderRouter) InitSimpleUploaderRouter(Router fiber.Router) {
	simpleUploadRouter := Router.Group("simpleUpload")
	var simpleUploadApi = v1.AppApi.ExampleApi.SimpleUploadApi
	{
		simpleUploadRouter.Post("upload", simpleUploadApi.SimpleUpload)      // 上传功能
		simpleUploadRouter.Get("checkFileMd5", simpleUploadApi.CheckFileMd5) // 文件完整度验证
		simpleUploadRouter.Get("mergeFileMd5", simpleUploadApi.MergeFileMd5) // 合并文件
	}
}
