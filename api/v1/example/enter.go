package example

import "FiberBoot/service"

type ApiGroup struct {
	CustomerApi
	ExcelApi
	FileTransferApi
	SimpleUploadApi
}

var fileTransferService = service.AppService.ExampleService.FileTransferService
var customerService = service.AppService.ExampleService.CustomerService
var excelService = service.AppService.ExampleService.ExcelService
var simpleUploadService = service.AppService.ExampleService.SimpleUploadService
