package example

import (
	"FiberBoot/global"
	"FiberBoot/model/common/response"
	"FiberBoot/model/example"
	"FiberBoot/utils"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type ExcelApi struct {
}

// ExportExcel
//
// @Tags Excel
// @Summary 导出Excel
// @Security ApiKeyAuth
// @accept application/json
// @Produce  application/octet-stream
// @Param data body example.ExcelInfo true "导出Excel文件信息"
// @Success 200
// @Router /excel/exportExcel [post]
func (e *ExcelApi) ExportExcel(c *fiber.Ctx) error {
	var excelInfo example.ExcelInfo
	_ = c.BodyParser(&excelInfo)
	filePath := global.CONFIG.Excel.Dir + excelInfo.FileName
	err := excelService.ParseInfoList2Excel(excelInfo.InfoList, filePath)
	if err != nil {
		global.LOG.Error("转换Excel失败!", zap.Any("err", err))
		return response.FailWithMessage("转换Excel失败", c)
	}

	c.Response().Header.Add("success", "true")
	err = c.SendFile(filePath)
	if err != nil {
		return err
	}
	return nil
}

// ImportExcel
// @Tags Excel
// @Summary 导入Excel文件
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "导入Excel文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"导入成功"}"
// @Router /excel/importExcel [post]
func (e *ExcelApi) ImportExcel(c *fiber.Ctx) error {
	header, err := c.FormFile("file")
	if err != nil {
		global.LOG.Error("接收文件失败!", zap.Any("err", err))
		return response.FailWithMessage("接收文件失败", c)
	}
	_ = c.SaveFile(header, global.CONFIG.Excel.Dir+"ExcelImport.xlsx")
	return response.OkWithMessage("导入成功", c)
}

// LoadExcel
// @Tags Excel
// @Summary 加载Excel数据
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"加载数据成功"}"
// @Router /excel/loadExcel [get]
func (e *ExcelApi) LoadExcel(c *fiber.Ctx) error {
	menus, err := excelService.ParseExcel2InfoList()
	if err != nil {
		global.LOG.Error("加载数据失败!", zap.Any("err", err))
		return response.FailWithMessage("加载数据失败", c)
	}
	return response.OkWithDetailed(response.PageResult{
		List:     menus,
		Total:    int64(len(menus)),
		Page:     1,
		PageSize: 999,
	}, "加载数据成功", c)
}

// DownloadTemplate
// @Tags Excel
// @Summary 下载模板
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param fileName query string true "模板名称"
// @Success 200
// @Router /excel/downloadTemplate [get]
func (e *ExcelApi) DownloadTemplate(c *fiber.Ctx) error {
	fileName := c.Query("fileName")
	filePath := global.CONFIG.Excel.Dir + fileName
	ok, err := utils.PathExists(filePath)
	if !ok || err != nil {
		global.LOG.Error("文件不存在!", zap.Any("err", err))
		return response.FailWithMessage("文件不存在", c)
	}
	c.Response().Header.Add("success", "true")
	err = c.SendFile(filePath)
	if err != nil {
		return err
	}
	return nil
}
