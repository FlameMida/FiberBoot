package example

import (
	"FiberBoot/global"
	"FiberBoot/model/common/response"
	exampleRes "FiberBoot/model/example/response"
	"FiberBoot/utils"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"io/ioutil"
	"mime/multipart"
	"strconv"
)

// BreakpointContinue
//
// @Tags FileTransfer
// @Summary 断点续传到服务器
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "an example for breakpoint resume, 断点续传示例"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"切片创建成功"}"
// @Router /fileTransfer/breakpointContinue [post]
func (u *FileTransferApi) BreakpointContinue(c *fiber.Ctx) error {
	fileMd5 := c.FormValue("fileMd5")
	fileName := c.FormValue("fileName")
	chunkMd5 := c.FormValue("chunkMd5")
	chunkNumber, _ := strconv.Atoi(c.FormValue("chunkNumber"))
	chunkTotal, _ := strconv.Atoi(c.FormValue("chunkTotal"))
	FileHeader, err := c.FormFile("file")
	if err != nil {
		global.LOG.Error("接收文件失败!", zap.Any("err", err))
		return response.FailWithMessage("接收文件失败", c)
	}
	f, err := FileHeader.Open()
	if err != nil {
		global.LOG.Error("文件读取失败!", zap.Any("err", err))
		return response.FailWithMessage("文件读取失败", c)
	}
	defer func(f multipart.File) {
		_ = f.Close()
	}(f)
	cen, _ := ioutil.ReadAll(f)
	if !utils.CheckMd5(cen, chunkMd5) {
		global.LOG.Error("检查md5失败!", zap.Any("err", err))
		return response.FailWithMessage("检查md5失败", c)
	}
	err, file := fileTransferService.FindOrCreateFile(fileMd5, fileName, chunkTotal)
	if err != nil {
		global.LOG.Error("查找或创建记录失败!", zap.Any("err", err))
		return response.FailWithMessage("查找或创建记录失败", c)
	}
	err, patch := utils.BreakPointContinue(cen, fileName, chunkNumber, fileMd5)
	if err != nil {
		global.LOG.Error("断点续传失败!", zap.Any("err", err))
		return response.FailWithMessage("断点续传失败", c)
	}

	if err = fileTransferService.CreateFileChunk(file.ID, patch, chunkNumber); err != nil {
		global.LOG.Error("创建文件记录失败!", zap.Any("err", err))
		return response.FailWithMessage("创建文件记录失败", c)
	}
	return response.OkWithMessage("切片创建成功", c)
}

// FindFile
//
// @Tags FileTransfer
// @Summary 查找文件
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "Find the file, 查找文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查找成功"}"
// @Router /fileTransfer/findFile [post]
func (u *FileTransferApi) FindFile(c *fiber.Ctx) error {
	fileMd5 := c.Query("fileMd5")
	fileName := c.Query("fileName")
	chunkTotal, _ := strconv.Atoi(c.Query("chunkTotal"))
	err, file := fileTransferService.FindOrCreateFile(fileMd5, fileName, chunkTotal)
	if err != nil {
		global.LOG.Error("查找失败!", zap.Any("err", err))
		return response.FailWithMessage("查找失败", c)
	} else {
		return response.OkWithDetailed(exampleRes.FileResponse{File: file}, "查找成功", c)
	}
}

// BreakpointContinueFinish
// @Tags FileTransfer
// @Summary 创建文件
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "上传文件完成"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"file uploaded, 文件创建成功"}"
// @Router /fileTransfer/findFile [post]
func (u *FileTransferApi) BreakpointContinueFinish(c *fiber.Ctx) error {
	fileMd5 := c.Query("fileMd5")
	fileName := c.Query("fileName")
	err, filePath := utils.MakeFile(fileName, fileMd5)
	if err != nil {
		global.LOG.Error("文件创建失败!", zap.Any("err", err))
		return response.FailWithDetailed(exampleRes.FilePathResponse{FilePath: filePath}, "文件创建失败", c)
	} else {
		return response.OkWithDetailed(exampleRes.FilePathResponse{FilePath: filePath}, "文件创建成功", c)
	}
}

// RemoveChunk
// @Tags FileTransfer
// @Summary 删除切片
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "删除缓存切片"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"缓存切片删除成功"}"
// @Router /fileTransfer/removeChunk [post]
func (u *FileTransferApi) RemoveChunk(c *fiber.Ctx) error {
	fileMd5 := c.Query("fileMd5")
	fileName := c.Query("fileName")
	filePath := c.Query("filePath")
	err := utils.RemoveChunk(fileMd5)
	err = fileTransferService.DeleteFileChunk(fileMd5, fileName, filePath)
	if err != nil {
		global.LOG.Error("缓存切片删除失败!", zap.Any("err", err))
		return response.FailWithDetailed(exampleRes.FilePathResponse{FilePath: filePath}, "缓存切片删除失败", c)
	} else {
		return response.OkWithDetailed(exampleRes.FilePathResponse{FilePath: filePath}, "缓存切片删除成功", c)
	}
}
