package example

import (
	"FiberBoot/global"
	"FiberBoot/model/example"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"io/ioutil"
	"os"
	"strconv"
)

type SimpleUploadService struct {
}

//@author: Flame
//@function: SaveChunk
//@description: 保存文件切片路径
//@param: uploader model.SimpleUploader
//@return: err error

func (exa *SimpleUploadService) SaveChunk(uploader example.SimpleUploader) (err error) {
	return global.DB.Create(uploader).Error
}

//@author: Flame
//@function: CheckFileMd5
//@description: 检查文件是否已经上传过
//@param: md5 string
//@return: err error, uploads []model.SimpleUploader, isDone bool

func (exa *SimpleUploadService) CheckFileMd5(md5 string) (err error, uploads []example.SimpleUploader, isDone bool) {
	err = global.DB.Find(&uploads, "identifier = ? AND is_done = ?", md5, false).Error
	isDone = errors.Is(global.DB.First(&example.SimpleUploader{}, "identifier = ? AND is_done = ?", md5, true).Error, gorm.ErrRecordNotFound)
	return err, uploads, !isDone
}

//@author: Flame
//@function: MergeFileMd5
//@description: 合并文件
//@param: md5 string, fileName string
//@return: err error

func (exa *SimpleUploadService) MergeFileMd5(md5 string, fileName string) (err error) {
	finishDir := "./finish/"
	dir := "./chunk/" + md5
	// 如果文件上传成功 不做后续操作 通知成功即可
	if !errors.Is(global.DB.First(&example.SimpleUploader{}, "identifier = ? AND is_done = ?", md5, true).Error, gorm.ErrRecordNotFound) {
		return nil
	}

	// 打开切片文件夹
	rd, err := ioutil.ReadDir(dir)
	_ = os.MkdirAll(finishDir, os.ModePerm)
	// 创建目标文件
	fd, err := os.OpenFile(finishDir+fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return
	}
	// 关闭文件
	defer func(fd *os.File) {
		_ = fd.Close()
	}(fd)
	// 将切片文件按照顺序写入
	for k := range rd {
		content, _ := ioutil.ReadFile(dir + "/" + fileName + strconv.Itoa(k+1))
		_, err = fd.Write(content)
		if err != nil {
			_ = os.Remove(finishDir + fileName)
		}
	}

	if err != nil {
		return err
	}
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		// 删除切片信息
		if err = tx.Delete(&example.SimpleUploader{}, "identifier = ? AND is_done = ?", md5, false).Error; err != nil {
			fmt.Println(err)
			return err
		}
		data := example.SimpleUploader{
			Identifier: md5,
			IsDone:     true,
			FilePath:   finishDir + fileName,
			Filename:   fileName,
		}
		// 添加文件信息
		if err = tx.Create(&data).Error; err != nil {
			fmt.Println(err)
			return err
		}
		return nil
	})

	err = os.RemoveAll(dir) // 清除切片
	return err
}
