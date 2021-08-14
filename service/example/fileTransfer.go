package example

import (
	"FiberBoot/global"
	"FiberBoot/model/common/request"
	"FiberBoot/model/example"
	"FiberBoot/utils/upload"
	"errors"
	"mime/multipart"
	"strings"
)

//@author: Flame
//@function: Upload
//@description: 创建文件上传记录
//@param: file model.FileTransfer
//@return: error

func (e *FileTransferService) Upload(file example.FileTransfer) error {
	return global.DB.Create(&file).Error
}

//@author: Flame
//@function: FindFile
//@description: 删除文件切片记录
//@param: id uint
//@return: error, model.FileTransfer

func (e *FileTransferService) FindFile(id uint) (error, example.FileTransfer) {
	var file example.FileTransfer
	err := global.DB.Where("id = ?", id).First(&file).Error
	return err, file
}

//@author: Flame
//@function: DeleteFile
//@description: 删除文件记录
//@param: file model.FileTransfer
//@return: err error

func (e *FileTransferService) DeleteFile(file example.FileTransfer) (err error) {
	var fileFromDb example.FileTransfer
	err, fileFromDb = e.FindFile(file.ID)
	oss := upload.NewOss()
	if err = oss.DeleteFile(fileFromDb.Key); err != nil {
		return errors.New("文件删除失败")
	}
	err = global.DB.Where("id = ?", file.ID).Unscoped().Delete(&file).Error
	return err
}

//@author: Flame
//@function: GetFileRecordInfoList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: err error, list interface{}, total int64

func (e *FileTransferService) GetFileRecordInfoList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.DB
	var fileLists []example.FileTransfer
	err = db.Find(&fileLists).Count(&total).Error
	err = db.Limit(limit).Offset(offset).Order("updated_at desc").Find(&fileLists).Error
	return err, fileLists, total
}

//@author: Flame
//@function: UploadFile
//@description: 根据配置文件判断是文件上传到本地或者七牛云
//@param: header *multipart.FileHeader, noSave string
//@return: err error, file model.FileTransfer

func (e *FileTransferService) UploadFile(header *multipart.FileHeader, noSave string) (err error, file example.FileTransfer) {
	oss := upload.NewOss()
	filePath, key, uploadErr := oss.UploadFile(header)
	if uploadErr != nil {
		panic(err)
	}
	if noSave == "0" {
		s := strings.Split(header.Filename, ".")
		f := example.FileTransfer{
			Url:  filePath,
			Name: header.Filename,
			Tag:  s[len(s)-1],
			Key:  key,
		}
		return e.Upload(f), f
	}
	return
}
