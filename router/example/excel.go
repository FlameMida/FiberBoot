package example

import (
	"FiberBoot/api/v1"
	"github.com/gofiber/fiber/v2"
)

type ExcelRouter struct {
}

func (e *ExcelRouter) InitExcelRouter(Router fiber.Router) {
	excelRouter := Router.Group("excel")
	var exaExcelApi = v1.AppApi.ExampleApi.ExcelApi
	{
		excelRouter.Post("/importExcel", exaExcelApi.ImportExcel)          // 导入Excel
		excelRouter.Get("/loadExcel", exaExcelApi.LoadExcel)               // 加载Excel数据
		excelRouter.Post("/exportExcel", exaExcelApi.ExportExcel)          // 导出Excel
		excelRouter.Get("/downloadTemplate", exaExcelApi.DownloadTemplate) // 下载模板文件
	}
}
