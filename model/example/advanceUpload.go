package example

import (
	"FiberBoot/global"
)

// File struct, 文件结构体
type File struct {
	global.MODEL
	FileName     string
	FileMd5      string
	FilePath     string
	ExaFileChunk []FileChunk
	ChunkTotal   int
	IsFinish     bool
}

// FileChunk file chunk struct, 切片结构体
type FileChunk struct {
	global.MODEL
	FileID          uint
	FileChunkNumber int
	FileChunkPath   string
}
