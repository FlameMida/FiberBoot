package example

import (
	"FiberBoot/global"
	"FiberBoot/model/common/request"
	"FiberBoot/model/example"
	"FiberBoot/model/system"
	systemService "FiberBoot/service/system"
)

type CustomerService struct {
}

//@author: Flame
//@function: CreateCustomer
//@description: 创建客户
//@param: e model.Customer
//@return: err error

func (exa *CustomerService) CreateCustomer(e example.Customer) (err error) {
	err = global.DB.Create(&e).Error
	return err
}

//@author: Flame
//@function: DeleteFileChunk
//@description: 删除客户
//@param: e model.Customer
//@return: err error

func (exa *CustomerService) DeleteCustomer(e example.Customer) (err error) {
	err = global.DB.Delete(&e).Error
	return err
}

//@author: Flame
//@function: UpdateCustomer
//@description: 更新客户
//@param: e *model.Customer
//@return: err error

func (exa *CustomerService) UpdateCustomer(e *example.Customer) (err error) {
	err = global.DB.Save(e).Error
	return err
}

//@author: Flame
//@function: GetCustomer
//@description: 获取客户信息
//@param: id uint
//@return: err error, customer model.Customer

func (exa *CustomerService) GetCustomer(id uint) (err error, customer example.Customer) {
	err = global.DB.Where("id = ?", id).First(&customer).Error
	return
}

//@author: Flame
//@function: GetCustomerInfoList
//@description: 分页获取客户列表
//@param: UserAuthorityID string, info request.PageInfo
//@return: err error, list interface{}, total int64

func (exa *CustomerService) GetCustomerInfoList(UserAuthorityID string, info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.DB.Model(&example.Customer{})
	var a system.Authority
	a.AuthorityId = UserAuthorityID
	err, auth := systemService.AuthorityServiceApp.GetAuthorityInfo(a)
	var dataId []string
	for _, v := range auth.DataAuthorityId {
		dataId = append(dataId, v.AuthorityId)
	}
	var CustomerList []example.Customer
	err = db.Where("user_authority_id in ?", dataId).Count(&total).Error
	if err != nil {
		return err, CustomerList, total
	} else {
		err = db.Limit(limit).Offset(offset).Preload("User").Where("user_authority_id in ?", dataId).Find(&CustomerList).Error
	}
	return err, CustomerList, total
}
