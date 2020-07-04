package mapper

import (
	clog "cwm.wiki/ad-CMS/common/log"
	"cwm.wiki/ad-CMS/model"
	"strconv"
)

// 添加订单
func InsertOrder(order model.Orders) error {
	err := model.DB.Create(&order).Error

	return err
}

//  查找全部订单
func SelectOrders() (*[]model.Orders, error) {
	var orders []model.Orders
	if err := model.DB.Find(&orders).Error; err != nil {
		clog.Error("DB record not found ", err)
		return nil, err
	}

	return &orders, nil
}

// 根据ID查找单个订单
func SelectOrderById(id string) (*model.Orders, error) {
	var order model.Orders
	if err := model.DB.Where("system_id = ?", id).First(&order).Error; err != nil {
		clog.Error("SelectOrderById", err)
		return nil, err
	}

	return &order, nil
}

// 修改订单
func UpdateOrder(update model.Orders) (*model.Orders, error) {
	var order *model.Orders
	order, err := SelectOrderById(strconv.Itoa(update.SystemID))
	if err != nil {
		clog.Error("UpdateUser", err)
		return nil, err
	}

	if order != nil {
		model.DB.Model(&order).Update(update)
	}

	return order, nil
}

// 删除订单
func DeleteOrder(id string) error {
	err := model.DB.Where("system_id = ?", id).Delete(model.Orders{}).Error

	return err
}
