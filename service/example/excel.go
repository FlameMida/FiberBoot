package example

import (
	"FiberBoot/global"
	"FiberBoot/model/system"
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"
	"strconv"
)

type ExcelService struct {
}

func (exa *ExcelService) ParseInfoList2Excel(infoList []system.BaseMenu, filePath string) error {
	excel := excelize.NewFile()
	err := excel.SetSheetRow("Sheet1", "A1", &[]string{"ID", "路由Name", "路由Path", "是否隐藏", "父节点", "排序", "文件名称"})
	if err != nil {
		return err
	}
	for i, menu := range infoList {
		axis := fmt.Sprintf("A%d", i+2)
		err := excel.SetSheetRow("Sheet1", axis, &[]interface{}{
			menu.ID,
			menu.Name,
			menu.Path,
			menu.Hidden,
			menu.ParentId,
			menu.Sort,
			menu.Component,
		})
		if err != nil {
			return err
		}
	}
	err = excel.SaveAs(filePath)
	return err
}

func (exa *ExcelService) ParseExcel2InfoList() ([]system.BaseMenu, error) {
	skipHeader := true
	fixedHeader := []string{"ID", "路由Name", "路由Path", "是否隐藏", "父节点", "排序", "文件名称"}
	file, err := excelize.OpenFile(global.CONFIG.Excel.Dir + "ExcelImport.xlsx")
	if err != nil {
		return nil, err
	}
	menus := make([]system.BaseMenu, 0)
	rows, err := file.Rows("Sheet1")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		row, err := rows.Columns()
		if err != nil {
			return nil, err
		}
		if skipHeader {
			if exa.compareStrSlice(row, fixedHeader) {
				skipHeader = false
				continue
			} else {
				return nil, errors.New("Excel格式错误 ")
			}
		}
		if len(row) != len(fixedHeader) {
			continue
		}
		id, _ := strconv.Atoi(row[0])
		hidden, _ := strconv.ParseBool(row[3])
		sort, _ := strconv.Atoi(row[5])
		menu := system.BaseMenu{
			MODEL: global.MODEL{
				ID: uint(id),
			},
			Name:      row[1],
			Path:      row[2],
			Hidden:    hidden,
			ParentId:  row[4],
			Sort:      sort,
			Component: row[6],
		}
		menus = append(menus, menu)
	}
	return menus, nil
}

func (exa *ExcelService) compareStrSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	if (b == nil) != (a == nil) {
		return false
	}
	for key, value := range a {
		if value != b[key] {
			return false
		}
	}
	return true
}
