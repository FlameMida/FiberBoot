package example

import (
	"FiberBoot/global"
	"FiberBoot/model/example"
	"errors"
	"gorm.io/gorm"
)

type FileTransferService struct {
}

//@author: Flame
//@function: FindOrCreateFile
//@description: 上传文件时检测当前文件属性，如果没有文件则创建，有则返回文件的当前切片
//@param: fileMd5 string, fileName string, chunkTotal int
//@return: err error, file model.File

func (e *FileTransferService) FindOrCreateFile(fileMd5 string, fileName string, chunkTotal int) (err error, file example.File) {
	var cFile example.File
	cFile.FileMd5 = fileMd5
	cFile.FileName = fileName
	cFile.ChunkTotal = chunkTotal

	if errors.Is(global.DB.Where("file_md5 = ? AND is_finish = ?", fileMd5, true).First(&file).Error, gorm.ErrRecordNotFound) {
		err = global.DB.Where("file_md5 = ? AND file_name = ?", fileMd5, fileName).Preload("FileChunk").FirstOrCreate(&file, cFile).Error
		return err, file
	}
	cFile.IsFinish = true
	cFile.FilePath = file.FilePath
	err = global.DB.Create(&cFile).Error
	return err, cFile
}

//@author: Flame
//@function: CreateFileChunk
//@description: 创建文件切片记录
//@param: id uint, fileChunkPath string, fileChunkNumber int
//@return: error

func (e *FileTransferService) CreateFileChunk(id uint, fileChunkPath string, fileChunkNumber int) error {
	var chunk example.FileChunk
	chunk.FileChunkPath = fileChunkPath
	chunk.FileID = id
	chunk.FileChunkNumber = fileChunkNumber
	err := global.DB.Create(&chunk).Error
	return err
}

//@author: Flame
//@function: DeleteFileChunk
//@description: 删除文件切片记录
//@param: fileMd5 string, fileName string, filePath string
//@return: error

func (e *FileTransferService) DeleteFileChunk(fileMd5 string, fileName string, filePath string) error {
	var chunks []example.FileChunk
	var file example.File
	err := global.DB.Where("file_md5 = ? AND file_name = ?", fileMd5, fileName).First(&file).Update("IsFinish", true).Update("file_path", filePath).Error
	err = global.DB.Where("exa_file_id = ?", file.ID).Delete(&chunks).Unscoped().Error
	return err
}
